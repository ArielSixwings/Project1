package main

import (
    "./imageprocessing"

    "fmt"
    "gocv.io/x/gocv"
    "image"
    "io"
    "os"
)

/**
 *
 * @brief      { function_description }
 * @return     { description_of_the_return_value }
 */
func main() {

    Marshist := gocv.NewMat() //mat for histogram equalization
    //MarsEdge := gocv.NewMat() //mat for Canny edge
    Marsgray := gocv.NewMat() //mat for read
    Mars := gocv.NewMat()     //mat for paint

    Marsgray = imageprocessing.ReadImage(Marsgray, "./imageprocessing/Images/Mars.bmp", true, false, false)
    Mars = imageprocessing.ReadImage(Mars, "./imageprocessing/Images/Mars.bmp", true, false, false)

    Marshist = imageprocessing.ProcessImage(Marsgray, Marshist, "./imageprocessing/Images/MarsHistogram.bmp", false, true)
    //MarsEdge = ProceedCanny(Marshist, "./Images/MarsEdges.bmp", true, true)

    var Base image.Point
    var Suplies = make([]image.Point, 8)
    for i := 0; i < 8; i++ {
        _, err := fmt.Scanf("%d", &Suplies[i].X)
        _, err = fmt.Scanf("%d", &Suplies[i].Y)
        if err != nil {
            fmt.Println(err)

        }
    }
    filename := "spots.txt"
    fmt.Println("the file name", filename)
    file, err := os.Create(filename)
    WhatIWrite, err := io.WriteString(file, fmt.Sprintf("%d\n", Base.X))
    WhatIWrite, err = io.WriteString(file, fmt.Sprintf("%d\n", Base.Y))

    for i := 0; i < 8; i++ {
        WhatIWrite, err = io.WriteString(file, fmt.Sprintf("%d\n", Suplies[i].X))
        WhatIWrite, err = io.WriteString(file, fmt.Sprintf("%d\n", Suplies[i].Y))
    }
    if err != nil {
        fmt.Println(err)
        fmt.Println(WhatIWrite)
    }

}
