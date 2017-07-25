package Algo

import (
			"fmt"

			"github.com/cavhack/frego/enc"

)

type AlgorithmFactory interface {

		Create(params interface{}) (Algo, error)
		ParamsEncoder() enc.Encoder


}