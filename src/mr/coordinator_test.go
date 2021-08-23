package mr

//import (
//	"testing"
//)
//
//func TestAssignTask(t *testing.T){
//	mapOutput:=map[string] []string{"1":[]string{"mr-test.json"}}
//	reply := MrReply{TaskType: -1}
//	coordinator:= Coordinator{MapInput: []string{"test.txt"},NReduce: 10}
//	coordinator.AssignTask(mapOutput,&reply)
//	if reply.TaskType!=0 || reply.NReduce!=coordinator.NReduce || reply.MapInput!= "test.txt"{
//		t.Errorf("Do not run as expected. TaskType:%d, Nreduce: %d, MapInput: %s", reply.TaskType,reply.NReduce,reply.MapInput)
//	}
//	for key :=range mapOutput{
//		if coordinator.ReduceInput[key]!=mapOutput[key]{
//
//		}
//	}
//}