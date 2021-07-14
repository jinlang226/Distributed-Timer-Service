package main

import (
	"encoding/csv"
	"fmt"
	"github.com/go-playground/log"
	"os"
	"strconv"
	"time"
)

//const (
//	taskName = iota
//	duration
//)

// data 读入数据保存map
//var data map[string]string

//func main() {
//	if err := readFiles(); err != nil {
//		log.Error(err)
//	}
//	log.Info("read data: ", data)
//
//	if err := writeCsv("./testdata/result.csv", data); err != nil {
//		log.Error(err)
//	}
//}

// readFile reads csv file
func readFile(filename string) ([][]string, error) {
	csvFile, err := os.Open(filename)
	if err != nil {
		fmt.Printf("read file error: %v", err)
		return nil, err
	}
	defer csvFile.Close()
	//创建csv读取接口实例
	ReadCsv := csv.NewReader(csvFile)

	stringValue, _ := ReadCsv.ReadAll()
	return stringValue, nil
}

// readFiles 读文件夹
//func readFiles() error {
//	// 路径
//	wholePathName := "./testdata/read_data"
//	// 读文件夹
//	rd, err := ioutil.ReadDir(wholePathName)
//	if err != nil {
//		return err
//	}
//	var results [][][]string
//	// 遍历文件夹内文件，数据存入results
//	for _, fi := range rd {
//		result, err := readFile(wholePathName + "/" + fi.Name())
//		if err != nil {
//			return err
//		}
//		if result != nil {
//			results = append(results, result)
//		}
//	}
//	// 数据存入data内
//	data = make(map[string]string)
//	for _, tasks := range results {
//		for _, items := range tasks {
//			data[items[taskName]] = items[duration]
//		}
//	}
//	return nil
//}

// writeCsv 写入csv文件
//func writeCsv(path string, data map[string]string) error {
//	for taskId, duration := range data {
//		// 记录开始睡眠的时间点
//		startTime := time.Now().Unix()
//		// 睡眠duration
//		sleepDuration, err := strconv.Atoi(duration)
//		if err != nil {
//			log.Error(err)
//		}
//		// 睡眠
//		time.Sleep(time.Duration(sleepDuration) * time.Second)
//		// 记录睡眠结束时间点
//		stopTime := time.Now().Unix()
//		// 写入的csv行数据：对应map中的一个KV
//		writeDataLine := []string{taskId, duration, strconv.FormatInt(startTime, 10), strconv.FormatInt(stopTime, 10)}
//		log.Info("write data by line: ", writeDataLine)
//		// 写入一行数据
//		if err := writeCsvByLine(path, writeDataLine); err != nil {
//			log.Error(err)
//		}
//	}
//	return nil
//}

// writeCsvByLine 写入一行数据
func writeCsvByLine(path string, dataStruct *writeDataByLine) error {
	//todo: bugs might remain, need mutex
	//OpenFile 读取文件，不存在时则创建，使用追加模式
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Error(err)
	}
	defer file.Close()

	// 创建写入接口
	WriterCsv := csv.NewWriter(file)

	startTime := strconv.Itoa(int(dataStruct.startTime))
	stopTime := strconv.Itoa(int(dataStruct.stopTime))
	taskId:=fmt.Sprintf("%v", dataStruct.taskId)
	duration:= fmt.Sprintf("%v", (dataStruct.duration)*time.Second)
	dataLine := []string{taskId, duration, startTime, stopTime}

	// 写数据
	if err := WriterCsv.Write(dataLine); err != nil {
		log.Error(err)
	}

	WriterCsv.Flush() // 刷新，不刷新是无法写入的
	return nil
}
