package mygraph

import "fmt"

type Edge struct {
    Number     int
    EdgeWeight float32
}

type Node struct {
    Edges []Edge
}

type Graph struct {
    Nodes []Node
    size  int
}

func (this *Graph) MakeGraph(Size int) {
    this.Nodes = make([]Node, Size)
    var i int = 0
    var auxedge int
    for i < (Size - 1) {
        this.Nodes[i].Edges = make([]Edge, (Size - (1 + i)))
        auxedge = 0
        for auxedge < (Size - (1 + i)) {
            this.Nodes[i].Edges[auxedge].Number = (i * auxedge) + 1
            this.Nodes[i].Edges[auxedge].EdgeWeight = 20.1
            fmt.Printf("%d\n", this.Nodes[i].Edges[auxedge].Number)
            fmt.Printf("%f\n", this.Nodes[i].Edges[auxedge].EdgeWeight)
            auxedge++
        }
        i++
    }
}
