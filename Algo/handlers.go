package Algo

import (

		"github.com/cavhack/frego"
)

type Handlers struct {

	OnDuplicateVertex        DuplicateVertexFunc
	OnMissingVertex			 MissingVertexFunc
	OnDuplicateEdge			 DuplicateEdgeFunc
	OnMissingEdge			 MissingEdgeFunc


}

type DuplicateVertexFunc func(id string, value1 interface{}, value2 interface{})(finalValue interface{}, err error)
type MissingVertexFunc func(id string) (initialValue interface{}, err error)
type DuplicateEdgeFunc func(from string, to string, value1 interface{}, value2 interface{})(finalValue interface{}, err error)
type MissingEdgeFunc func(from string, to string) (initialValue, interface{}, err error)

func DefaultHandlers() *Handlers {

				return &Handlers{onDuplicateVertex, onMissingVertex, onDuplicateEdge, onMissingEdge}

}

func onDuplicateVertex() {

}

func onMissingVertex(id string) (initialValue interface{}, err error) {
	glog.Infof("running default onDuplicateEdge handler; edge from %s to %s", from, to)
	return value1, nil
}

func onDuplicateEdge(from string, to string, value1 interface{}, value2 interface{})(interface{}, error) {
	glog.Infof("running default onDuplicateEdge handler; edge from %s to %s", from, to)
	return value1, nil
}

func onMissingEdge(from string, to string) (initialValue interface{}, err error) {

	glog.Info("running default onMissingEdge handler; edge from %s to %s", from, to)
	return nil, nil
}









