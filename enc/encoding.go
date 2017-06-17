package enc

type Encoder struct {


    Marshal Marshaler
    Unmarshal Unmarshaler

}

type Unmarshaler func([]byte) (interface{}, error)
type Marshaler func(interface{}) ([]byte, error)
