package mr
//
//import (
//	"log"
//	"net"
//	"net/http"
//	"net/rpc"
//	"os"
//)
//
//
//type Coordinator struct {
//	// Your definitions here.
//	Workers map[string] string
//	MapInput []string
//	NReduce int
//	ReduceInput map[string] []string
//}
//
//// Your code here -- RPC handlers for the worker to call.
//
////
//// an example RPC handler.
////
//// the RPC argument and reply types are defined in rpc.go.
////
//func (c *Coordinator) Example(args *ExampleArgs, reply *ExampleReply) error {
//	reply.Y = args.X + 1
//	return nil
//}
//
//func (c *Coordinator)  AnswerRPC(args *MrRequest, reply *MrReply) error{
//	switch args.RequestType {
//	case 0:
//		c.AssignTask(args.MapOutput,reply)
//	default:
//		break
//	}
//	return nil
//}
//
//func (c *Coordinator) AssignTask(mapOutput map[string] []string, reply *MrReply)  {
//	if mapOutput!=nil && len(mapOutput)!=0{
//		for key :=range mapOutput{
//			c.ReduceInput[key]=append(c.ReduceInput[key],mapOutput[key]...)
//		}
//	}
//	if len(c.MapInput)>0 {
//		println(c.MapInput)
//		println(len(c.MapInput))
//		reply.TaskType = 0
//		reply.NReduce=c.NReduce
//		reply.MapInput = c.MapInput[len(c.MapInput)-1]
//		c.MapInput = c.MapInput[:len(c.MapInput)-1]
//		return
//	} else if c.NReduce>0{
//		println(c.NReduce)
//		reply.TaskType=1
//		reply.ReduceInput=c.ReduceInput[string(rune(c.NReduce))]
//		c.NReduce--
//		return
//	}else {
//		reply.TaskType=10
//		return
//	}
//}
//
//
////
//// start a thread that listens for RPCs from worker.go
////
//func (c *Coordinator) server() {
//	rpc.Register(c)
//	rpc.HandleHTTP()
//	//l, e := net.Listen("tcp", ":1234")
//	sockname := coordinatorSock()
//	os.Remove(sockname)
//	l, e := net.Listen("unix", sockname)
//	if e != nil {
//		log.Fatal("listen error:", e)
//	}
//	go http.Serve(l, nil)
//}
//
////
//// main/mrcoordinator.go calls Done() periodically to find out
//// if the entire job has finished.
////
//func (c *Coordinator) Done() bool {
//	ret := false
//
//	// Your code here.
//
//
//	return ret
//}
//
////
//// create a Coordinator.
//// main/mrcoordinator.go calls this function.
//// nReduce is the number of reduce tasks to use.
////
//func MakeCoordinator(files []string, nReduce int) *Coordinator {
//	c := Coordinator{MapInput: files,NReduce: nReduce}
//
//	// Your code here.
//
//
//	c.server()
//	return &c
//}
