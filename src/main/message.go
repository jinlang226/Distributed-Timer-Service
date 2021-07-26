package main

import (
	"time"
)

const Format string = "2006/1/2 15:04:05"

var readFilename = "data.csv"

//var Filepath = "/root/ft_local/Distributed-Timer-Service/src/test"
var Filepath = "./src/test/"
var logFilename = "log1.csv"

const (
	serverID0 = 0
	serverID1 = 1
	serverID2 = 2
	serverID3 = 3
	serverID4 = 4
)

var AcceptorIds = []int{serverID1, serverID2, serverID3, serverID4}
var LearnerIds = []int{serverID4}

const (
	Socketname0 string = "9.134.131.104" //mVUFd@2873tB
	Socketname2 string = "9.134.72.227"  //PMqpN@5628eJ
	Socketname3 string = "9.135.113.126" //wpgqs*9728Jn
	Socketname4 string = "9.134.167.39"  //dpksZ*5439Cp
	Socketname1 string = "9.134.131.104"
)

var SocketNames = []string{Socketname0, Socketname1, Socketname2, Socketname3, Socketname4}

//local server name
var LocalName = Socketname0

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
