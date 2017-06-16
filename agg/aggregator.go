package aggregator

type Aggregator interface {

      Type()  string


      Get()   interface{}
      Set(value interface{}) error

      GetBytes() ([]byte, error)
      SetBytes(bytes []byte) error

      Clone() Aggregator


}
