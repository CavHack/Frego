package Parser

import (

      "fmt"
      "io"

      "github/cavhack/frego"


)

type VertexParser  interface{

        Next() (*frego.Vertex, error)


}

type EdgeParser    interface {

        Next() (*frego.Edge, error)


}

func NewVertexParser(name string, reader io.Reader) (VertexParser, error) {

      var parser   VertexParser
      var error    error

      switch name {

      case "dimacs" :
        parser = NewDimacsVertexParser(reader)

      default:
          err = fmt.Errorf("Invalid parser '%s'", name)

      }

      return parser, err

}


func NewEdgeParser(name string, reader io.Reader) (EdgeParser, error) {

      var parser EdgeParser
      var err   error


      switch name {

      case "dimacs" :
        parser = NewDimacsEdgeParser(reader)

      default:
        err = fmt.Errorf("Invalid parser '%s'", name)


      }

      return parser, err



}
