package Algo


import (

      "github.com/cavhack/frego/agg"
      "github.com/cavhack/frego/enc"
)

type Algorithm interface {

            Compute(context *VertexContext, message interface{}) error
            GetResult(aggregators *aggregator.ImmutableAggregatorSet) interface{}
            GetResultDisplayValue(result interface{}) string

            VertexMessageCombiner()   VertexMessageCombiner

            VertexMessageEncoder()  encoding.Encoder
            VertexValueEncoder() encoding.Encoder
            EdgeValueEncoder()   encoding.Encoder
            VertexMutableValueEncoder() encoding.Encoder



}
