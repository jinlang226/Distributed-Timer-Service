package main

import "time"

const format string = "2006/1/2 15:04:05"

var tw *TimeWheel
var filename = "idk"
var filepath = "wholePathName"

var serverID1 = 1
var serverID2 = 2
var serverID3 = 3
var serverID4 = 4
var serverID5 = 5

var acceptorIds = []int{serverID2, serverID3, serverID4}
var learnerIds = []int{serverID5}

const (
	Socketname1 string = "IP1"
	Socketname2 string = "IP2"
	Socketname3 string = "IP3"
	Socketname4 string = "IP4"
	Socketname5 string = "IP5"
)

var socketNames = []string{Socketname1, Socketname2, Socketname3, Socketname4, Socketname5}

//local server name
var localName = Socketname1

// writeDataByLine 写入数据数据结构
type writeDataByLine struct {
	taskId    interface{}
	duration  time.Duration
	startTime int64
	stopTime  int64
}

// 要写入的数据
type writeData struct {
	writeData []writeDataByLine
}
