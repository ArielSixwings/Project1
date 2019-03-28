package main

import (
    "./graph"
)

/**
 * @brief      { Solve the travel problem by brute force }
 *
 * @param      Size  The size
 * @param      Base   The base node
 *
 * @return     { Return the not sorted path, which is a Cheackpoint struct}
 */
func (this *(mygraph.Graph)) SolveTravel(Size int, Base int) mygraph.Checkpoint {
    Path := [Size]mygraph.Checkpoint{}
    Path[0].Label = Base
    Path[(Size - 1)].Label = Base
    Path.TotalWeight = 0
for i := 1; i < Size     ; i++ {
    
            for B := 1; B < Size     ; B++ {
            if B == i {
                B++
            }  //if        
                for C := 1; C < Size     ; C++ {
                if C == B {
                    C++
                }  //if            
                    for D := 1; D < Size     ; D++ {
                    if D == C {
                        D++
                    }  //if                
                        for E := 1; E < Size     ; E++ {
                        if E == D {
                            E++
                        }  //if                    
                            for F := 1; F < Size     ; F++ {
                            if F == E {
                                F++
                            }  //if                        
                                for G := 1; G < Size     ; G++ {
                                if G == F {
                                    G++
                                }  //if                            
                                    for H := 1; H < Size     ; H++ {
                                    if H == G {
                                        H++
                                    }  //if
                                    Path[i*B*C*D*E*F*G].Label[H] = this.Nodes[i].Edges[auxedge].Number
                                    Path[i*B*C*D*E*F*G].Weight[H] = this.Nodes[i].Edges[auxedge].EdgeWeight
                                    Path[i*B*C*D*E*F*G].TotalWeight = (Path[i*B*C*D*E*F*G].TotalWeight + this.Nodes[i].Edges[auxedge].EdgeWeight)
                                }  //for
                            }  //for
                        }  //for
                    }  //for
                }  //for
            }  //for
        }  //for
    }  //for
    return path
}  //funcmain