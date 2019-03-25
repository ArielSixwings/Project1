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
    imageprocessing.PaintCircle(Mars, Center, 10, color.RGBA{
        rgbacolor.SetColor(rgbacolor.Black, 0),
        rgbacolor.SetColor(rgbacolor.Black, 1),
        rgbacolor.SetColor(rgbacolor.Black, 2),
        0}, 5, true, false)
    //PaintPath(Mars, Base, Suplies)
    //
    var ThisGraph mygraph.Graph
    ThisGraph.MakeGraph(8)
    // ShowImage("The Canne was done!",MarsEdge,0)

    mygraph.BestPath(ThisGraph, 1)

}
