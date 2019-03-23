package main

import (
    "./colorgalery"
    "./graph"
    "./imageprocessing"

    "gocv.io/x/gocv"
    "image"
    "image/color"
)

func main() {

    Marshist := gocv.NewMat() //mat for histogram equalization
    //MarsEdge := gocv.NewMat() //mat for Canny edge
    Marsgray := gocv.NewMat() //mat for read
    Mars := gocv.NewMat()     //mat for paint

    Marsgray = imageprocessing.ReadImage(Marsgray, "./imageprocessing/Images/Mars.bmp", true, false, false)
    Mars = imageprocessing.ReadImage(Mars, "./imageprocessing/Images/Mars.bmp", true, false, true)

    Marshist = imageprocessing.ProcessImage(Marsgray, Marshist, "./imageprocessing/Images/MarsHistogram.bmp", true, true)
    //MarsEdge = ProceedCanny(Marshist, "./Images/MarsEdges.bmp", true, true)

    var Center image.Point

    var Base image.Point
    var Suplies image.Point

    Center.X = 400
    Center.Y = 400

    Base.X = 283
    Base.Y = 320

    Suplies.X = 800
    Suplies.Y = 920
    var HotPink = [4]uint8{153, 51, 102, 0} //HotPink
    imageprocessing.PaintCircle(Mars, Center, 10, color.RGBA{
        SetColor(HotPink, 0, 0.0),
        SetColor(HotPink, 1, 0.0),
        SetColor(HotPink, 2, 0.0),
        SetColor(HotPink, 3, 1.1)}, 5, true, false)
    //PaintPath(Mars, Base, Suplies)
    //
    var ThisGraph mygraph.Graph
    ThisGraph.MakeGraph(8)
    // ShowImage("The Canne was done!",MarsEdge,0)

}
