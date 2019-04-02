package main

import (
    "./graph"
    "strconv"
)

/**
 * @brief      Reads lines.
 *
 * @param      path  The path
 *
 * @return     { All the file in the string and an error safe mensage}
 */
func ReadLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

/**
 * @brief      { Solve the travel problem by brute force }
 *
 * @param      Size  The size
 * @param      Base   The base node
 *
 * @return     { Return the not sorted path, which is a Cheackpoint struct}
 */
func SolveTravel(Size int, Base int, this *(mygraph.Graph)) [](mygraph.Checkpoint) {
    Path := make([]mygraph.Checkpoint, Size)
    Path[0].Label[0] = Base
    Path[(Size - 1)].Label[(Size - 1)] = Base
    for i := 1; i < Size; i++ {

        for B := 1; B < Size; B++ {
            if B == i {
                B++
            } //if
            for C := 1; C < Size; C++ {
                if C == B {
                    C++
                } //if
                for D := 1; D < Size; D++ {
                    if D == C {
                        D++
                    } //if
                    for E := 1; E < Size; E++ {
                        if E == D {
                            E++
                        } //if
                        for F := 1; F < Size; F++ {
                            if F == E {
                                F++
                            } //if
                            for G := 1; G < Size; G++ {
                                if G == F {
                                    G++
                                } //if
                                for H := 1; H < Size; H++ {
                                    if H == G {
                                        H++
                                    } //if
                                    Path[i*B*C*D*E*F*G].Label[H] = this.Nodes[i].Edges[H].Number
                                    Path[i*B*C*D*E*F*G].Weight[H] = this.Nodes[i].Edges[H].EdgeWeight
                                    Path[i*B*C*D*E*F*G].TotalWeight = (Path[i*B*C*D*E*F*G].TotalWeight + this.Nodes[i].Edges[H].EdgeWeight)
                                } //for
                            } //for
                        } //for
                    } //for
                } //for
            } //for
        } //for
    } //for
    return Path
} //funcmain

func main() {
    var GraphData []string
    var ThisGraph mygraph.Graph
    var size = ThisGraph.InitGraph(8)
    lines, err := ReadLines("dataset8.txt")
    for i := 0; i < len(lines); i++ {
        comma := strings.IndexByte(lines[i], '*')
        var begin = i * comma
        var buildstring = string(lines[begin])
        for index := 0; index < comma; index++ {
            buildstring = buildstring + string(lines[comma+index])

        }
        GraphData[i] = buildstring
    }
    if err != nil {
        log.Fatalf("readLines: %s", err)
    }
    SolveTravel(Size, 0, this)
}
