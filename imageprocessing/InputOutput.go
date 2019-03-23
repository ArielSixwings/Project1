package imageprocessing

import (
    "gocv.io/x/gocv"
    "path/filepath"
)

func ReadImage(Image gocv.Mat, path string, show bool, save bool, colorfull bool) gocv.Mat {

    ImagePath := filepath.Join(path) //set path to the base image

    if colorfull {
        Image = gocv.IMRead(ImagePath, gocv.IMReadUnchanged) //read the base image as as RGB
    } else {
        Image = gocv.IMRead(ImagePath, gocv.IMReadGrayScale) //read the base image in grayscale
    }

    if show {
        ShowImage("And this is yout image", Image, 0)
    }

    return Image
}

func SaveImage(Name string, Image gocv.Mat) {
    gocv.IMWrite(Name, Image) //save the image
}
func ShowImage(Menssage string, Image gocv.Mat, time int) {
    window := gocv.NewWindow(Menssage) //basic window
    window.IMShow(Image)               //show the image
    window.WaitKey(time)
}
