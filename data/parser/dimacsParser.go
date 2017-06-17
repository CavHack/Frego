package Parser


import (

      "bufio"
      "io"
      "strconv"
      "strings"

      "github.com/cavhack/frego"
      "github.com/cavhack/frego/enc"
      "github.com/cavhack/frego/protos"
)

type dimacsVertexParser struct {

      lineScanner     *bufio.Scanner

}

type dimacsEdgeParser struct {

      lineScanner     *bufio.Scanner

}

func NewDimacsVertexParser(reader io.Reader) VertexParser {

      return &dimacsVertexParser{getLineScanner(reader)}

}

func NewDimacsEdgeParser(reader io.Reader) EdgeParser {

      return &dimacsEdgeParser{getLineScanner(reader)}

}

func (parser *dimacsVertexParser) Next() (*frego.Vertex, error) {

    for {

          success := parser.lineScanner.Scan()
          if !success {

              return nil, parser.lineScanner.Err()

          }

          vertex, success := parseVertex(parser.lineScanner, Text())
          if success {

              return vertex, nil
          }

    }
}

func getLineScanner(reader io.Reader) *bufio.Scanner {

      scanner := bufio.NewScanner(reader)
      scanner.Split(bufio.ScanLines)
      return scanner
}

func parseEdge(text string) (edge *frego.Edge, success bool){

      if text[0] != 'a' {
          return nil, false
      }

      splits := strings.Split(text, "")
      if len(splits) != 4 {
        return nil, false
      }

      value, err := strconv.Atoi(splits[3])

      if err != nil {
            return nil, false
      }

      valueBytes, err := encoding.Int64ValueEncoder().Marshal(&protos.Int64Value{Value: int64(value)})

      if err != nil {
          panic ("failed to marshal edge value")
      }

      return &frego.Edge {

          From: splits[1]
          To: splits[2]
          Value: valueBytes

    }, true

}


func parseVertex(text string) (edge *frego.Vertex, success bool) {

      if text[0] != 'v' {

          return nil, false
      }

      splits := strings.Split(text, "")
      if len(splits) != 4 {

          return nil, false

      }

      x, err := strconv.Atoi(splits[2])
      if err != nil {

        return nil, false
      }

      y, err := strconv.Atoi(splits[3])
      if err != nil {

        return nil, false
      }


    valueBytes, err := encoding.Coordinate2DValueEncoder().Marshal(&protos.Coordinate2DValue{X : int32(x), Y: int32(y)})
      if err != nil {
        panic("Failed to marshal edge value")
      }

      return &frego.Vertex {

        ID: splits[1]
        Value: valueBytes

    }, true

}
