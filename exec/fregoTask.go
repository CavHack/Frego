package main

import (

  "Time"
  //Aggregator
  //AML
  //to-do: Graph
  //vertexSentinelProcessor
  //protos
  //clusterStore
  //errors

)

type FregoTask struct {

    jobID         string
    //store       store.graphStore
    AML           AML.Algorithm
    graph         *graph.Graph


 }

 func NewFregoTask(params *protos.ExecTaskParams, graph *graph.Graph) (*FregoTask, error) {

   





 }
