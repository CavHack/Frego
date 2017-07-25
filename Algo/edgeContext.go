package Algo

type EdgeContext struct {

	from 			*VertexContext
	To				string
	Value			interface{}
	MutableValue	interface{}


}


func NewEdgeContext(from *VertextContext, to string, value interface{}, mutableValue interface{}) * EdgeContext {


		return &EdgeContext{from: from, To: to, Value: value, MutableValue: mutableValues}

}

func (c *edgeContext ) setValue(value interface{}) {

	c.from.op.SetEdgeValue()
}
