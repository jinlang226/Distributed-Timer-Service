package main

// writeDataByLine 写入数据数据结构
type writeDataByLine struct {
	taskId    string
	duration  string
	startTime string
	stopTime  string
}

// 要写入的数据
type writeData struct {
	writeData []writeDataByLine
}