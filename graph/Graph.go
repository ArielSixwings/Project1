package mygraph

import "fmt"

/**
 * @brief      Class for edge.
 */
type Edge struct {
    Number     int
    EdgeWeight float32
}

/**
 * @brief      Class for node.
 */
type Node struct {
    Edges []Edge
}

/**
 * @brief      Class for graph. Main
 */
type Graph struct {
    Nodes []Node
    size  int
}

/**
 * @brief      Makes a graph.
 *
 * @param      Size  The size
 *
 * @return     { The Graph size }
 */
func (this *Graph) MakeGraph(Size int) int {
    this.Nodes = make([]Node, Size)
    this.size = Size
    // var Weights []float32{}
    for i := 0; i < Size; i++ {
        this.Nodes[i].Edges = make([]Edge, (Size - i))
        for edge := 1; edge < Size-i; edge++ {
            this.Nodes[i].Edges[edge].Number = edge + i
            fmt.Printf("%d node %d    ", (edge + i), i)
            this.Nodes[i].Edges[edge].EdgeWeight = float32(i + edge)
        } //for
        fmt.Println("\n")
    } //for
    return this.size
}

/**
 * @brief      { Inits a empty graph }
 *
 * @param      Size  The size
 *
 * @return     { The Graph size }
 */
func (this *Graph) InitGraph(Size int) int {
    this.Nodes = make([]Node, Size)
    this.size = Size

    return this.size
}
