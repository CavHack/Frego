package syntacticProcessor 

import(   "sync"

		  "github.com/cavhack/frego"
		  "github.com/cavhack/frego/agg"
		  "github.com/cavhack/frego/Algo"
		  "github.com/pkg/errors"
)

type contextOperations struct {

	algorithm					algorithm.algorithm
	errorChan					chan error
	addedVertices				map[string]interface{}
	addedVerticesMutex			sync.Mutex
	removedVertices				map[string] bool
	removedVerticesMutex		sync.Mutex
	changedVertexValue			map[string] interface{}
	changedVertexValueMutex     



}

type edge struct { 

	from string
	to string

}

func newContextOperations(Algo algo.Algo) {

}














