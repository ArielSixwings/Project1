package main

import (
    "gocv.io/x/gocv"
    "image"
)

func main() {
    Marsgray := gocv.NewMat() //mat for read
    Marsgray = gocv.IMRead("Mars.bmp", gocv.IMReadUnchanged)
    for i := 0; i < 100; i++ {
        for j := 0; j < 100; j++ {
            image.Pt(i, j) = 0

        }
    }
    window := gocv.NewWindow("Did it worked ?") //basic window
    window.IMShow(Marsgray)                     //show the image
    window.WaitKey(0)
}
