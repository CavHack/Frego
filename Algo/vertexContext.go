package Algo

import (

      "github.com/cavhack/frego/agg"

)

type VertexContext  struct {

          id                             string
          Edges                         []*EdgeContext
          Value                         interface{}
          MutableValue                  interface{}
          Superstep                     int
          PrevSuperstepAggregators      *aggregator.ImmutableAggregatorSet
          op                            ContextOperations
}

func NewVertexContext(id string, superstep int, value interface{}, mutableValue interface{},

            operations ContextOperations, PrevSuperstepAggregators *aggregator.ImmutableAggregatorSet) *VertexContext {

            return &VertexContext{

                    id:                             id,
                    Superstep:                      superstep,
                    Value:                          value,
                    mutableValue:                   mutableValue,
                    PrevSuperstepAggregators:       prevSuperstepAggregators
                    op: operations,

            }

          }


func (c *VertexContext)  ID() string {

        return c.id

}


func (c *VertexContext) AddVertex(id string, value interface{}) {

    c.op.AddVertex(id, value)

}

func (c *VertexContext) RemoveVertex(id string) {

        c.op.RemoveVertex(id)

}

func (c *VertexContext) Remove() {

      c.op.RemoveVertex(c.id)

}


func (c *VertexContext) SetVertexValue(id string, value interface{}) {


  c.op.SetVertexValue(id, value)


}

func (c *VertexContext) SetValue (value interface{}) {

  c.op.SetVertexValue(c.id, value)
}

func (c *VertexContext) SendMessageTo(to string, message interface{}) {

  c.op.SendVertexMessage(to, message)
}
