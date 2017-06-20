package frego

import "time"

type JobNode struct {

     ID		                string
     Label		            string
     DateCreated          Time.time
     Status		            JobGraphStatus
     Store		            string
     StoreParams	        []byte
     AML		              string
     AMLParams		        []byte
     CPU		              float64
     MEM		              float64
     Vertices		           int
    


}

type JobNodeStatus int

     const (
		_JobNodeStatus = iota
		JobNodeCreated
		JobNodeRunning
		JobNodeCompleted
		JobNodeCancelled
		JobNodeFailed


)


func (j JobNode) CanCancel() bool {
     	return j.Status == JobNodeCreated || j.Status == JobNodeRunning

}
