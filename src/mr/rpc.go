//package mr
//
////
//// RPC definitions.
////
//// remember to capitalize all names.
////
//
//import "os"
//import "strconv"

package mr

//
// RPC definitions.
//
// remember to capitalize all names.
//

import (
"os"
"time"
)
import "strconv"

type worker struct {
	UUID   string
	Status string
	// tasks timeout
	TaskTimeout time.Time
	Task        *task
}
type task struct {
	Action string
	File   string
}

type Args struct {
	Worker *worker
}

type Reply struct {
	//Success bool
	NReduce int
	IsMapFinished bool
	IsAllFinished bool
	NextWorker *worker
}

// Add your RPC definitions here.


// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the master.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func masterSock() string {
	s := "/var/tmp/824-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}

//
// example to show how to declare the arguments
// and reply for an RPC.
//

//type ExampleArgs struct {
//	X int
//}
//
//type MrRequest struct {
//	//0: ask for new task
//	//1: heart beat
//	//
//	RequestType int
//	MapOutput map[string] []string
//}
//
//type MrReply struct {
//	TaskType int
//	MapInput string
//	ReduceInput []string
//	NReduce int
//}
//
//type ExampleReply struct {
//	Y int
//}

// Add your RPC definitions here.


// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the coordinator.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func coordinatorSock() string {
	s := "/var/tmp/824-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}
