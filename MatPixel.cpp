#include "MatPixel.hpp"

void ReturnPixel(cv::(&Mat) Image, int Row, int Col,int channel) {
    return Image.at<cv::Vec3b>(Row,Col)[channel];
}
void SetPixel(cv::(&Mat) Image, int Row, int Col,int Blue, int Green, int Red) {
    Image.at<cv::Vec3b>(Row,Col)[0] = Blue;
    Image.at<cv::Vec3b>(Row,Col)[1] = Green;
    Image.at<cv::Vec3b>(Row,Col)[2] = Red;
}
void ReturnGrayPixel(cv::(&Mat) Image, int Row, int Col) {
    return Image.at<Vec3b>(Row,Col);
}
void SetGrayPixel(cv::(&Mat) Image, int Row, int Col,int value) {
    Image.at<uint8_t>(Row,Col) = value;
}
