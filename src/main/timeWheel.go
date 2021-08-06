package main

import (
	"container/list"
	"errors"
	"fmt"
	"github.com/go-playground/log/v7"
	"net"
	"strconv"
	"sync"
	"time"
)

var TW *TimeWheel

// TimeWheel的核心结构体
type TimeWheel struct {
	// 时间轮盘的精度
	interval time.Duration
	// 时间轮盘每个位置存储的Task列表
	slots  []*list.List
	ticker *time.Ticker
	// 时间轮盘当前的位置
	currentPos int
	// 时间轮盘的齿轮数 interval*slotNums就是时间轮盘转一圈走过的时间
	slotNums          int
	addTaskChannel    chan *Task
	removeTaskChannel chan *Task
	stopChannel       chan bool
	// Map结构来存储Task对象，key是Task.key，value是Task在双向链表中的存储对象，本文的结构是list.Element
	taskRecords *sync.Map
	// 需要执行的任务，如果时间轮盘上的Task执行同一个Job，可以直接实例化到TimeWheel结构体中。
	// 此处的优先级低于Task中的Job参数
	//job       Job
	isRunning     bool
	finishedTasks *sync.Map //update according to the logs by RPC
	finishedTasksOrdinaryMap map[int]int
	mutex         sync.Mutex
	lis           net.Listener
}

// 需要执行的Job的函数结构体
type Job func(interface{})

// 时间轮盘上需要执行的任务
type Task struct {
	// 用来标识task对象，是唯一的
	key interface{}
	// 任务周期
	interval time.Duration
	// 任务的创建时间
	createdTime time.Time
	// 任务在轮盘的位置
	pos int
	// 任务需要在轮盘走多少圈才能执行
	circle int
	// 任务需要执行的Job，优先级高于TimeWheel中的Job
	job Job
	// 任务需要执行的次数，如果需要一直执行，设置成-1
	//times int
	stopTime int64
}

var once sync.Once

// CreateTimeWheel 用来实现TimeWheel的单例模式
func CreateTimeWheel(interval time.Duration, slotNums int) *TimeWheel {
	once.Do(func() {
		TW = New(interval, slotNums)
	})
	return TW
}

// New 初始化一个TimeWheel对象
func New(interval time.Duration, slotNums int) *TimeWheel {
	if interval <= 0 || slotNums <= 0 {
		return nil
	}
	tw := &TimeWheel{
		interval:          interval,
		slots:             make([]*list.List, slotNums),
		currentPos:        0,
		slotNums:          slotNums,
		addTaskChannel:    make(chan *Task),
		removeTaskChannel: make(chan *Task),
		stopChannel:       make(chan bool),
		taskRecords:       &sync.Map{},
		//job:               job,
		isRunning:     false,
		finishedTasks: &sync.Map{},
		finishedTasksOrdinaryMap: make(map[int]int),
	}

	tw.initSlots()
	return tw
}

// Start 启动时间轮盘
func (tw *TimeWheel) startTW() {
	tw.ticker = time.NewTicker(tw.interval)
	go tw.start(time.Now())
	tw.isRunning = true
}

// Stop 关闭时间轮盘
//func (tw *TimeWheel) Stop() {
//	tw.stopChannel <- true
//	tw.isRunning = false
//}

// IsRunning 检查一下时间轮盘的是否在正常运行
//func (tw *TimeWheel) IsRunning() bool {
//	return tw.isRunning
//}

//type foo struct {
//	bar bool
//}
//
//func (tw *TimeWheel) Finished(args interface{}, reply foo) error {
//	tw.taskRecords.Range(func(k, v interface{}) bool {
//		if k == nil && v == nil {
//			reply.bar = true
//			return true
//		}
//		reply.bar = false
//		return false
//	})
//	return nil
//}

// AddTask 向时间轮盘添加任务的开放函数
// @param interval    任务的周期
// @param key         任务的key，必须是唯一的，否则添加任务的时候会失败
// @param createTime  任务的创建时间
// func (tw *TimeWheel) AddTask(interval time.Duration, key interface{}, createdTime time.Time, job Job) error {
func (tw *TimeWheel) AddTask(args *AddTaskArgs, reply *AddTaskReply) error {
	interval := args.interval
	key := args.taskJob
	//createdTime := args.execTime
	uuid := args.uuid
	if interval <= 0 || key == nil {
		return errors.New("Invalid task params")
	}

	// 检查Task.Key是否已经存在
	_, exist := tw.taskRecords.Load(uuid)
	if exist {
		return errors.New("Duplicate task key")
	}
	tw.addTaskChannel <- &Task{
		key:         uuid,
		interval:    interval,
		createdTime: time.Now(),
		//job:         job,
		//times:       times,
	}
	fmt.Println("successfully add tasks")
	return nil
}

// 初始化时间轮盘，每个轮盘上的卡槽用一个双向队列表示，便于插入和删除
func (tw *TimeWheel) initSlots() {
	for i := 0; i < tw.slotNums; i++ {
		tw.slots[i] = list.New()
	}
}

// 启动时间轮盘的内部函数
func (tw *TimeWheel) start(startTime time.Time) {
	for {
		select {
		case <-tw.ticker.C:
			fmt.Println("=========== case 1: ticker ===========, and time is: ", time.Since(startTime))
			tw.checkAndRunTask()
		case task := <-tw.addTaskChannel:
			// 此处利用Task.createTime来定位任务在时间轮盘的位置和执行圈数
			// 如果直接用任务的周期来定位位置，那么在服务重启的时候，任务周器相同的点会被定位到相同的卡槽，
			// 会造成任务过度集中
			fmt.Println("=========== case 2: addTask ===========")
			tw.addTask(task)
			//case task := <-tw.removeTaskChannel:
			//	fmt.Println("case 3 ===== ")
			//	tw.taskExe(task)
			//case <-tw.stopChannel:
			//	fmt.Println("case 4 ===== ")
			//	tw.ticker.Stop()
			//	return
		}
	}
}

// 检查该轮盘点位上的Task，看哪个需要执行
func (tw *TimeWheel) checkAndRunTask() {
	// 获取该轮盘位置的双向链表
	currentList := tw.slots[tw.currentPos]

	if currentList != nil {
		for item := currentList.Front(); item != nil; {
			task := item.Value.(*Task)
			//fmt.Println("created task: ", task.key, "task time: ", task.interval.Seconds(), ", created time: ", task.createdTime.Format(Format))
			//fmt.Println("stop now: ", time.Now().Format(Format))
			// 如果圈数>0，表示还没到执行时间，更新圈数
			_, existed := tw.finishedTasks.Load(task.key)
			
			if existed {
				continue
			}
			if task.circle > 0 {
				task.circle--
				item = item.Next()
				continue
			}

			next := item.Next()
			item = next
			// 检查该Task是否存在
			_, ok := tw.taskRecords.Load(task.key)
			if !ok {
				log.Info(fmt.Sprintf("Task key %d doesn't existed in task list, please check your input", task.key))
			} else { //todo how to make it works?
				//tw.removeTaskChannel <- task
				task.stopTime = time.Now().Unix()
				tw.taskExe(task)
			}

		}
	}
	// 轮盘前进一步
	if tw.currentPos == tw.slotNums-1 {
		tw.currentPos = 0
	} else {
		tw.currentPos++
	}
}

// 添加任务的内部函数
// @param task       Task  Task对象
// Task.createTime生成
func (tw *TimeWheel) addTask(task *Task) {
	pos, circle := tw.getPosAndCircleByCreatedTime(task.createdTime, task.interval, task.key)

	task.circle = circle
	task.pos = pos

	element := tw.slots[pos].PushBack(task)
	tw.taskRecords.Store(task.key, element)
}

func WriteToMap(key interface{}) {
	TW.finishedTasks.Store(key, 1)
	TW.finishedTasksOrdinaryMap[key.(int)] = 250
}

//机器宕机之后，读log恢复map
func TraverseMap() {
	result, err := ReadFile(Filepath + logFilename)
	if err != nil {
		fmt.Println("err in read file")
	}
	// for each line in csv data structure:
	for _, items := range result {
		fmt.Println(items)
		uuid, err := strconv.Atoi(items[0])

		if err != nil {
			log.Error(err)
		}
		WriteToMap(uuid)
		fmt.Println("uuid: ", uuid)
	}
}

// 删除任务的内部函数
func (tw *TimeWheel) taskExe(task *Task) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	// 从map结构中删除
	val, _ := tw.taskRecords.Load(task.key)
	tw.taskRecords.Delete(task.key)

	// 通过TimeWheel.slots获取任务的
	currentList := tw.slots[task.pos]
	currentList.Remove(val.(*list.Element))

	//write to the local cache
	WriteToMap(task.key)

	data := &WriteDataByLine{
		StopTime:  task.stopTime,
		TaskId:    task.key,
		Duration:  task.interval,
		StartTime: task.createdTime.Unix(),
	}

	//write to local log
	writeCsvByLine(Filepath+logFilename, data)

	//write to other servers' log, mark as completed by paxos
	log.Info("origin data is: ", data)
	value := p.Propose(data)
	log.Info("propose value is: ", value)

	p.round = proposerID
	p.number = 0
}

// 该函数用任务的创建时间来计算下次执行的位置和圈数
func (tw *TimeWheel) getPosAndCircleByCreatedTime(createdTime time.Time, d time.Duration, key interface{}) (int, int) {
	delaySeconds := int(d.Seconds())
	intervalSeconds := int(tw.interval.Seconds())
	circle := delaySeconds / intervalSeconds / tw.slotNums
	pos := (tw.currentPos + delaySeconds/intervalSeconds) % tw.slotNums

	// 特殊case，当计算的位置和当前位置重叠时，因为当前位置已经走过了，所以circle需要减一
	if pos == tw.currentPos && circle != 0 {
		circle--
	}
	return pos-1, circle
}
