package soliton-graph

type Graph  struct {

  key                  string
  superstep            int
  vertices             map[string]*vertex

}


type Vertex struct {

  value           interface{}
  mutableValue    interface{}
  edges           map[string]*edge


}

type edge struct {

  value         interface{}
  mutableValue  interface{}

}

func NewGraph(capacity int) *Graph {

      return &Graph{vertices: make(map[string]*vertex, capacity)}
}

func (g *Graph) setVertexMutableValue(id string, value interface{}){

        v:= g.ensureVertex(id)
        v.value = value
}

func (g *Graph) setVertexValue(id string, value interface{}) {

    v := g.ensureVertex(id)
    v.value = value
}

func (g *Graph) setEdgeValue(from string, to string, value interface{}) {

      e := g.ensureEdge(from, to)
      e.value = value

}

func (g *Graph) setEdgeMutableValue(from string, to string, value interface{}) {

      e:= g.ensureEdge(from, to)
      e.mutableValue = value

}

func (g *Graph) removeVertex(id string) {

        _, ok := g.vertices[id]
        if !ok {

              return
        }

        delete(g.vertices, id)

        for _, v:= range g.vertices {
              delete(v.edges,id)
        }

  }

func (g *Graph) removeEdge(from string, to string) {

  v, ok := g.vertices[from]
  if !ok {

      return
  }

  delete(v.edges, to)

}


func (g *Graph) VertexValue(id string)(interface{}, bool) {

    v, ok := g.vertices[id]
    if !ok {
        return nil, false
    }

    return v.value, true

}

func (g *Graph) VertexMutableValue(id string)(interface{}, bool) {

    v, ok := g.vertices[id]
    if !ok {

      return nil, false
    }

    return v.mutableValue, true

}

func (g *Graph) HasVertex(id string) bool {

      _, ok := g.VertexValue(id)
      return ok

  }

func (g *Graph) EdgeValue(from string, to string) (interface{}, bool) {

  v, ok := g.vertices[from]
    if !ok {
        return nil, false
    }

  e, ok := v.edges[to]
  if !ok {
      return nil, false
  }

  return e.value, true
}

func (g *Graph) EdgeMutableValue(from string, to string) (interface{}, bool) {

      v, ok := g.vertices[from]
      if !ok {
          return nil, false
      }

      e, ok := v.edges[to]
      if !ok {

        return nil, false
      }

      return e.mutableValue, true

}

func (g *Graph) HasEdge(from string, to string) bool {

    _, ok := g.EdgeValue(from, to)
    return ok

}

func (g *Graph) Vertices() []string {

    result := make([]string, len(g.vertices))
    i := 0
    for k := range g.vertices {
        result[i] = k
        i++

    }

    result

}



func (g *Graph) EdgesFrom(id string) []string {

      v, ok := g.vertices[id]
      if !ok {
                return make([]string, 0)
      }

    result := make([]string, len(v.edges))
    i := 0
    for k := range v.edges {
        result[i] = k
        i++
    }

    return result

}

func (g *Graph) ensureVertex(id string) *vertex {
  v, ok := g.vertices[id]
  if !ok  {

      v = &vertex{edges: make(map[string]*edge)}
      g.vertices[id] = v
  }

  return v
}

func (g *Graph) ensureEdge(from string, to string) *edge {

    v := g.ensureVertex(from)

    e, ok := v.edges[to]
    if !ok {
          e = &edge{}
          v.edges[to] = e

    }

    return e


}
