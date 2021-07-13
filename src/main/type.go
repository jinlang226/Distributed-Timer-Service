package main

import "modu/src/timeWheel"

const format string = "2006/1/2 15:04:05"
var tw *timeWheel.TimeWheel
var filename = "idk"
var filepath = "wholePathName"

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

