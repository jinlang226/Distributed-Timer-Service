package main

import (
	"time"
)

const Format string = "2006/1/2 15:04:05"

//var Filename = "log1.csv"
//var Filepath = "/root/ft_local"
var Filepath = "/Users/wjl/Desktop/Distributed-Timer-Service/src/test/log1.csv"
var Filename =""

const (
	serverID1 = 1
	serverID2 = 2
	serverID3 = 3
	serverID4 = 4
	serverID5 = 5
)

var AcceptorIds = []int{serverID2, serverID3, serverID4, serverID5}
var LearnerIds = []int{serverID5}

const (
	Socketname1 string = "9.134.131.104" //mVUFd@2873tB
	Socketname2 string = "9.134.72.227"  //PMqpN@5628eJ
	Socketname3 string = "9.135.113.126" //wpgqs*9728Jn
	Socketname4 string = "9.134.167.39"  //dpksZ*5439Cp
	Socketname5 string = "IP5"
)

var SocketNames = []string{Socketname1, Socketname2, Socketname3, Socketname4, Socketname5}

//local server name
var LocalName = Socketname1

// WriteDataByLine 写入数据数据结构
type WriteDataByLine struct {
	TaskId    interface{}
	Duration  time.Duration
	StartTime int64
	StopTime  int64
}

// 要写入的数据
type writeData struct {
	writeData []WriteDataByLine
}
