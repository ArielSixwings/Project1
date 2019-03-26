package mygraph

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
    var i int = 0
    var auxedge int
    for i < (Size - 1) {
        this.Nodes[i].Edges = make([]Edge, (Size - (1 + i)))
        auxedge = 0
        for auxedge < (Size - (1 + i)) {
            this.Nodes[i].Edges[auxedge].Number = (i * auxedge) + 1
            this.Nodes[i].Edges[auxedge].EdgeWeight = 20.1
            auxedge++
        } //for
        i++
    } //for
    return this.size
}
