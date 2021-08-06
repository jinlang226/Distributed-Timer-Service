package main

import (
	"encoding/csv"
	"fmt"
	"github.com/go-playground/log/v7"
	"os"
	"strconv"
	"syscall"
	"time"
)

//文件锁
type FileLock struct {
	dir string
	f   *os.File
}
func NewFileLock(dir string) *FileLock {
	return &FileLock{
		dir: dir,
	}
}
//加锁
func (l *FileLock) Lock() error {
	f, err := os.Open(l.dir)
	if err != nil {
		return err
	}
	l.f = f
	err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		return fmt.Errorf("cannot flock directory %s - %s", l.dir, err)
	}
	return nil
}
//释放锁
func (l *FileLock) Unlock() error {
	defer l.f.Close()
	return syscall.Flock(int(l.f.Fd()), syscall.LOCK_UN)
}


// ReadFile reads csv file
func ReadFile(filename string) ([][]string, error) {
	csvFile, err := os.Open(filename)
	if err != nil {
		log.Error("read file error: %v", err)
		return nil, err
	}
	defer csvFile.Close()
	//创建csv读取接口实例
	ReadCsv := csv.NewReader(csvFile)

	stringValue, _ := ReadCsv.ReadAll()
	return stringValue, nil
}

//  写入一行数据
func writeCsvByLine(path string, dataStruct *WriteDataByLine) error {
	//todo: bugs might remain, need mutex
	err := flock.Lock()
	defer flock.Unlock()
	if err != nil {
		log.Error("file lock err: ", err.Error())
	}

	//OpenFile 读取文件，不存在时则创建，使用追加模式
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Error(err)
	}
	defer file.Close()

	// 创建写入接口
	WriterCsv := csv.NewWriter(file)

	startTime := strconv.Itoa(int(dataStruct.StartTime))
	stopTime := strconv.Itoa(int(dataStruct.StopTime))
	taskId:=fmt.Sprintf("%v", dataStruct.TaskId)
	duration:= strconv.Itoa(int(dataStruct.Duration)/int(time.Second))
	dataLine := []string{taskId, duration, startTime, stopTime}

	// 写数据
	if err := WriterCsv.Write(dataLine); err != nil {
		log.Error(err)
	}

	WriterCsv.Flush() // 刷新，不刷新是无法写入的
	return nil
}
