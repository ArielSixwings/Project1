package mygraph

import (
    "fmt"
    //"math"
)

type Checkpoint struct {
    Weight float32
    label  int
}

func BestPath(this Graph, Base int) {
    Path := [5040]Checkpoint{}
    var count int = 0
    Path[0].label = Base
    Path[(this.size - 1)].label = Base
    for i := 1; i < this.size; i++ {
        for j := 0; j < this.size; j++ {
            Path[j].label = this.Nodes[i].Edges[j].Number
            Path[j].Weight = this.Nodes[i].Edges[j].EdgeWeight
            count++
        }
    }
    fmt.Printf(" I found an amount of %d paths \n ", count)
}
