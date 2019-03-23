package imageprocessing

import (
    "gocv.io/x/gocv"
    "image"
    "image/color"
)

func PaintCircle(Image gocv.Mat, center image.Point, radius int, color color.RGBA, thickness int, show bool, save bool) {
    gocv.Circle(&Image, center, radius, color, thickness)

    //ALL available colors
    {

    }

    //func Circle(img *Mat, center image.Point, radius int, c color.RGBA := color.RGBA, thickness int)

    if show {
        ShowImage("The Circle was done!", Image, 0)
    }

    if save {
        SaveImage("CircleMade", Image)
    }
}

// func PaintPath(Image gocv.Mat, Initial image.Point, Destiny image.Point) {
//     for Initial.X != Destiny.X {
//         Initial.X = Initial.X + 1
//         for Initial.Y != Destiny.Y {
//             Initial.Y = Initial.Y + 1
//             Image.at < gocv.Vec3b > (Initial) = HotPink
//             ShowImage("Steps", Image, 5)
//         }
//     }
// }
