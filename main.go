package main

import (
    "./colorgalery"
    "./graph"
    "./imageprocessing"

    "fmt"
    "gocv.io/x/gocv"
    "image"
    "image/color"
    "os/exec"
    //"math"
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

    var ThisGraph mygraph.Graph
    var auxsize = ThisGraph.InitGraph(8)

    mygraph.BuildPaths(auxsize)

    out, err := exec.Command("./run.sh").Output()
    if err != nil {
        fmt.Printf("I found some error to run the program: %s", err)
    } else {
        fmt.Println("\nCommand Successfully Executed")
    }

    output := string(out[:])
    fmt.Println(output)

}
