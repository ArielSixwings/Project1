
#include <opencv2/highgui.hpp>
#include <stdio>
extern "C" {
/**
 * @brief      Returns a pixel.
 *
 * @param      Image  The image
 * @param[in]  Row    The row
 * @param[in]  Col    The col
 */
void ReturnPixel(cv::Mat& Image, int Row, int Col)

/**
 * @brief      Sets the pixel.
 *
 * @param      Image  The image
 * @param[in]  Row    The row
 * @param[in]  Col    The col
 * @param[in]  Blue   The blue
 * @param[in]  Green  The green
 * @param[in]  Red    The red
 */
void SetPixel(cv::Mat& Image, int Row, int Col,int Blue, int Green, int Red)

/**
 * @brief      Returns a gray pixel.
 *
 * @param      Image  The image
 * @param[in]  Row    The row
 * @param[in]  Col    The col
 */
void ReturnGrayPixel(cv::Mat& Image, int Row, int Col)

/**
 * @brief      Sets the gray pixel.
 *
 * @param      Image  The image
 * @param[in]  Row    The row
 * @param[in]  Col    The col
 * @param[in]  value  The value
 */
void SetGrayPixel(cv::Mat& Image, int Row, int Col,int value)
}
