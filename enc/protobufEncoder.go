package enc


import (


            //proto
            //error

)


func NewProtobufEncoder(msgFactory func() proto.Message) Encoder {

      unmarshaler := func(bytes []byte) (interface{}, error) {

          return protobufUnmarshaler(bytes, msgFactory)

      }

      return Encoder { Marshaler(protobufMarshaler), Unmarshaler(unmarshaler)}

}

func protobufMarshaler(message interface{}) ([]byte, error) {





}
