package main

//import (
//    "fmt"
//    "io/ioutil"
//      "mygraph"
//)

/**
 * @brief      { Solve the travel problem by brute force }
 *
 * @param      Size  The size
 * @param      Base   The base node
 *
 * @return     { Return the not sorted path, which is a Cheackpoint struct}
 */
func (this *Graph) SolveTravel(Size int, Base int) Checkpoint {
    Path := [Size]Checkpoint{}
    Path[0].Label = Base
    Path[(Size - 1)].Label = Base
    Path.TotalWeight = (2*this.Nodes[i].Edges[auxedge].EdgeWeight)
for i := 1; i < this.size; i++ {
    
            for B := 1; B < this.size; B++ {
            if B == i {
                B++
            }  //If        
                for C := 1; C < this.size; C++ {
                if C == B {
                    C++
                }  //If            
                    for D := 1; D < this.size; D++ {
                    if D == C {
                        D++
                    }  //If                
                        for E := 1; E < this.size; E++ {
                        if E == D {
                            E++
                        }  //If                    
                            for F := 1; F < this.size; F++ {
                            if F == E {
                                F++
                            }  //If                        
                                for G := 1; G < this.size; G++ {
                                if G == F {
                                    G++
                                }  //If                            
                                    for H := 1; H < this.size; H++ {
                                    if H == G {
                                        H++
                                    }  //If
                                    Path[A*B*C*D*E*F*G].Label[H] = this.Nodes[i].Edges[auxedge].Number
                                    Path[A*B*C*D*E*F*G].Weight[H] = this.Nodes[i].Edges[auxedge].EdgeWeight
                                    Path[A*B*C*D*E*F*G].TotalWeight = (Path[A*B*C*D*E*F*G].TotalWeight + this.Nodes[i].Edges[auxedge].EdgeWeight)
                                }  //For
                            }  //For
                        }  //For
                    }  //For
                }  //For
            }  //For
        }  //For
    }  //For
    return path
}  //funcmain