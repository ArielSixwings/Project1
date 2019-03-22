package main


import (
    "gocv.io/x/gocv"
    "image"
    "image/color"
)

func PaintCircle(Image gocv.Mat, center image.Point,radius int, color color.RGBA, thickness int,show bool, save bool){
    gocv.Circle(&Image,center,radius,color, thickness)

    //ALL available colors
    {
        //var HotPink RGBA
        //HotPink = color.RGBA{0x99,0x33,0x66}//HotPink

        // var LightPink color.RGBA := color.RGBA{0xFF,0x99,0xCC}//LightPink

        // var Fuchsia color.RGBA := color.RGBA{0xFF,0x00,0xFF}//Fuchsia

        // var Purple color.RGBA := color.RGBA{0x80,0x00,0x80}//Purple

        // var LightPurple color.RGBA := color.RGBA{0xCC,0x99,0xFF}//LightPurple

        // var Navy color.RGBA := color.RGBA{0x00,0x00,0x80}//Navy

        // var DeeoBlue color.RGBA := color.RGBA{0x33,0x33,0x99}//DeeoBlue

        // var GrayBlue color.RGBA := color.RGBA{0x66,0x66,0x99}//GrayBlue

        // var MediumBlue color.RGBA := color.RGBA{0x33,0x66,0xFF}//MediumBlue

        // var DarkBlue color.RGBA := color.RGBA{0x00,0x33,0x66}//DarkBlue

        // var WhiteBlue color.RGBA := color.RGBA{0x99,0xCC,0xFF}//WhiteBlue

        // var GrayLightBlue color.RGBA := color.RGBA{0x00,0xCC,0xFF}//GrayLightBlue

        // var Aqua color.RGBA := color.RGBA{0x00,0xFF,0xFF}//Aqua

        // var Teal color.RGBA := color.RGBA{0x00,0x80,0x80}//Teal

        // var LightBlue color.RGBA := color.RGBA{0x33,0xCC,0xCC}//LightBlue

        // var WhiteGreen color.RGBA := color.RGBA{0xCC,0xFF,0xFF}//WhiteGreen

        // var BlueGreen color.RGBA := color.RGBA{0x33,0x99,0x66}//BlueGreen

        // var Lime color.RGBA := color.RGBA{0x00,0xFF,0x00}//Lime

        // var Green color.RGBA := color.RGBA{0x00,0x80,0x00}//Green

        // var DarckGreen color.RGBA := color.RGBA{0x00,0x33,0x00}//DarckGreen

        // var BabyBlue color.RGBA := color.RGBA{0xCC,0xFF,0xCC}//BabyBlue

        // var LightGreen color.RGBA := color.RGBA{0x99,0xCC,0x00}//LightGreen

        // var Yellow color.RGBA := color.RGBA{0xFF,0xFF,0x00}//Yellow

        // var Olive color.RGBA := color.RGBA{0x80,0x80,0x00}//Olive

        // var MossGreen color.RGBA := color.RGBA{0x33,0x33,0x00}//MossGreen

        // var PaleYellow color.RGBA := color.RGBA{0xFF,0xFF,0x99}//PaleYellow

        // var StrongYellow color.RGBA := color.RGBA{0xFF,0xCC,0x00}//StrongYellow

        // var BurntYellow color.RGBA := color.RGBA{1xFF,0x99,0x00}//BurntYellow

        // var PaleOrange color.RGBA := color.RGBA{1xFF,0xCC,0x99}//PaleOrange

        // var Orange color.RGBA := color.RGBA{0xFF,0x66,0x00}//Orange

        // var Brown color.RGBA := color.RGBA{0x99,0x33,0x00}//Brown

        // var Red color.RGBA := color.RGBA{0xFF,0x00,0x00}//Red

        // var Maroon color.RGBA := color.RGBA{0x80,0x00,0x00}//Maroon

        // var White color.RGBA := color.RGBA{0xFF,0xFF,0xFF}//White

        // var Silver color.RGBA := color.RGBA{0xC0,0xC0,0xC0}//Silver

        // var LightGray color.RGBA := color.RGBA{0x96,0x96,0x96}//LightGray

        // var Gray color.RGBA := color.RGBA{0x80,0x80,0x80}//Gray

        // var DarkGray color.RGBA := color.RGBA{0x33,0x33,0x33}//DarkGray

        // var Black color.RGBA := color.RGBA{0x00,0x00,0x00}//Black
    }

//func Circle(img *Mat, center image.Point, radius int, c color.RGBA := color.RGBA, thickness int)

    if show{
        ShowImage("The Circle was done!",Image,0)
    }

    if save{
        SaveImage("CircleMade",Image)
    }
}

func PaintPath(Image gocv.Mat,Initial image.Point,Destiny image.Point){
    for Initial.X != Destiny.X{
        Initial.X = Initial.X + 1
        for Initial.Y != Destiny.Y{
            Initial.Y = Initial.Y + 1
            Image.at<Vec3b>(Initial) = color.RGBA{0x00, 0x00, 0x00, 0x00};
            ShowImage("Steps",Image,5)
        }
    }
}
