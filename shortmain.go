package main

//#include "MatPixel.hpp"
// void ReturnPixel(name char*, int Row, int Col,int channel) {
//     return Image.at<Vec3b>(Row,Col)[channel];
// }
// void SetPixel(name char*,int size, int Row, int Col,int Blue, int Green, int Red) {

//     Image = imread(name, CV_LOAD_IMAGE_COLOR); //reed the image
//     Image.at<Vec3b>(Row,Col)[0] = Blue;
//     Image.at<Vec3b>(Row,Col)[1] = Green;
//     Image.at<Vec3b>(Row,Col)[2] = Red;
//     imwrite(name,Image);
// }
// void ReturnGrayPixel(int[] IntName, int Row, int Col) {
//     return Image.at<Vec3b>(Row,Col);
// }
// void SetGrayPixel(int[] IntName, int Row, int Col,int value) {
//     Image.at<uint8_t>(Row,Col) = value;
// }
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

func SetPixelGo(name string, Row int, Col int, Blue int, Green int, Red int) bool {
    C.SetPixel(name, Row, Col, Blue, Green, Red)
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
    //var Mars = []int{77, 97, 114, 115, 46, 98, 109, 112}
    fmt.Printf("%d\n", C.AuxFunc())
    for i := 0; i < 100; i++ {
        for j := 0; j < 100; j++ {
            SetPixelGo("Mars.bmp", i, j, 0, 0, 0)
        }
    }
    imageprocessing.ShowImage("Oh mann, did it worked ?", Mars, 0)

}
