package enc

import (

          "github.com/cavhack/frego/protos"
          "github.com/gogo/protobuf/proto"
)

var (

        boolValueEncoder                  = NewProtobufEncoder(func() proto.Message {  return new(protos.BoolValue) })
        int32ValueEncoder                 = NewProtobufEncoder(func() proto.Message {  return new(protos.Int32) })
        Int64ValueEncoder                 = NewProtobufEncoder(func() proto.Message {  return new(protos.Int64Value) })
        stringValueEncoder                = NewProtobufEncoder(func()

)
