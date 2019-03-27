package main

/*
#cgo linux CFLAGS: -I./lpsolve
#cgo linux LDFLAGS: -L./lpsolve -llpsolve55 -Wl,-rpath=./lpsolve
#cgo darwin CFLAGS: -I/opt/local/include/lpsolve
#cgo darwin LDFLAGS: -L/opt/local/lib -llpsolve55
#include "lp_lib.h"
...
*/
import "C"

func GoSetPixel(Image gocv.Mat, Row int, Col int, Blue int, Green int, Red int) {
    C.SetPixel(Image, Row, Col, Blue, Green, Red)
}
func main() {
    // Image = gocv.IMRead("Mars.bmp", gocv.IMReadUnchanged) //read the base image as as RGB
    // for i := 0; i < 100; i++ {
    //     for j := 0; j < 100; j++ {
    //         GoSetPixel(Image, i, j, 0, 0, 0)
    //     }
    // }

    // window := gocv.NewWindow("did it worked?") //basic window
    // window.IMShow(Image)                       //show the image
    // window.WaitKey(0)
    C.cout << "oi"
}
