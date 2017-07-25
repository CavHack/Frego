package agg

import (

			"fmt"
			"reflect"

			"github.com/cavhack/frego/enc"
			"github.com/cavhack/frego/protos"

)


type int64Agg struct {

		value int64
}

func (agg *int64Agg) GetBytes() ([]byte, error) {

		return enc.Int64ValueEncoder().Marshal(&protos.Int64Value{Value:agg.value})

}


func (agg *int64Agg) SetBytes(bytes []byte) error {

	value, err := enc.Int64ValueEncoder().Unmarshal(bytes)
	if err != nil {

		return err
	}

	agg.value = value.(*protos.Int64Value).Value
	return nil
}






