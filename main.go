package main

import (
    "gocv.io/x/gocv"
    "image"
    "image/color"
)

func main() {

    Marshist := gocv.NewMat()//mat for histogram equalization
    MarsEdge := gocv.NewMat()//mat for Canny edge
    Marsgray := gocv.NewMat()//mat for read
    Mars :=  gocv.NewMat()//mat for paint

    Marsgray = ReadImage(Marsgray,"./Images/Mars.bmp",true,false,false)
    Mars = ReadImage(Mars,"./Images/Mars.bmp",true,false,true)

    Marshist = ProcessImage(Marsgray, Marshist,"./Images/MarsHistogram.bmp",true,true)
    MarsEdge = ProceedCanny(Marshist,"./Images/MarsEdges.bmp",true,true)

    var Center image.Point

    var Base image.Point
    var Suplies image.Point


    Center.X = 400
    Center.Y = 400

    Base.X = 283
    Base.Y = 320

    Suplies.X = 800
    Suplies.Y = 920

    PaintCircle(Mars,Center,10,color.RGBA{0x00, 0x00, 0x00, 0x00},5,true,false)
    PaintPath(Mars,Base,Suplies)

    ShowImage("The Canne was done!",MarsEdge,0)

}
