package agg

import (

		"fmt"
		"reflect"

		"github.com/cavhack/frego/enc"
		"github.com/cavhack/frego/protos"

)


type boolAgg struct {

	value bool

}

func (agg *boolAgg) GetBytes() ([]byte, error) {

	return enc.BoolValueEncoder().Marshal(&protos.BoolValue{Value: agg.value})

}

func (agg *boolAgg) SetBytes(bytes []byte) error {

	value, err := enc.BoolValueEncoder().Unmarshal(bytes)
	if err != nil {
		return err
	}

	agg.value = value.(*protos.BoolValue).Value
	return nil


}


func (agg *boolAgg) convertToBool(value interface{}) (bool, error ){

				b, ok := value.(bool)
				if ok {

					return b, nil
				}

				proto, ok := value.(*protos.BoolValue)
				if ok {

					return proto.Value, nil

				}

				return false, fmt.Errorf("invalid value type")
}

func (agg *boolAgg) clone() *boolAgg {


	return &boolAgg{agg.value}
}

type orAgg struct {


	*boolAgg

}

func newOrAgg() Aggregator {
	return &andAgg{boolAgg: &boolAgg}
}




type andAgg struct {

	*boolAgg
}




