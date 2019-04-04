package main

//#include "MatPixel.hpp"
// void SetPixel(Mat Image, int Row, int Col,int Blue, int Green, int Red) {
//     Image.at<Vec3b>(Row,Col)[0] = Blue;
//     Image.at<Vec3b>(Row,Col)[1] = Green;
//     Image.at<Vec3b>(Row,Col)[2] = Red;

//int AuxFunc(){
//    return 1;
//}
import "C"
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

func SetPixelGo(Image gocv.Mat, Row int, Col int, Blue int, Green int, Red int) bool {
    C.SetPixel
}

/**
 *
 * @brief      { function_description }
 * @return     { description_of_the_return_value }
 */
func main() {

    Mars := gocv.NewMat() //mat for paint

    Mars = imageprocessing.ReadImage(Mars, "./imageprocessing/Images/Mars.bmp", true, false, true)

    for i := 0; i < 100; i++ {
        for j := 0; j < 100; j++ {
            C.SetPixel(&Mars, i, j, 0, 0, 0)
        }
    }
    fmt.Printf("%d", C.AuxFunc())
    imageprocessing.ShowImage("Oh mann, did it worked ?", Mars, 0)

}
