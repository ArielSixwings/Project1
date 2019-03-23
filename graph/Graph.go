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
}

func (ThisGraph *Graph) MakeGraph(Size int) {
    ThisGraph.Nodes = make([]Node, Size)
    var i int = 0
    var auxedge int
    for i < (Size - 1) {
        ThisGraph.Nodes[i].Edges = make([]Edge, (Size - (1 + i)))
        auxedge = 0
        for auxedge < (Size - (1 + i)) {
            ThisGraph.Nodes[i].Edges[auxedge].Number = (i * auxedge) + 1
            ThisGraph.Nodes[i].Edges[auxedge].EdgeWeight = 20.1
            fmt.Printf("%d\n", ThisGraph.Nodes[i].Edges[auxedge].Number)
            fmt.Printf("%f\n", ThisGraph.Nodes[i].Edges[auxedge].EdgeWeight)
            auxedge++
        }
        i++
    }
}
