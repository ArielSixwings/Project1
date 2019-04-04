#ifndef PrefixHeader_matpixel
#define PrefixHeader_matpixel

#ifdef __cplusplus

#include <opencv2/highgui.hpp>
#include <stdio>
#include "stdint"
#include <iostream>
#include <string>
using namespace std;
using namespace cv;

extern "C" {
/**
 * @brief      Returns a pixel.
 *
 * @param      Image  The image
 * @param[in]  Row    The row
 * @param[in]  Col    The col
 */
void ReturnPixel(int[] IntName, int Row, int Col,int channel)

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
void SetPixel(name char*,int size, int Row, int Col,int Blue, int Green, int Red)

/**
 * @brief      Returns a gray pixel.
 *
 * @param      Image  The image
 * @param[in]  Row    The row
 * @param[in]  Col    The col
 */
void ReturnGrayPixel(int[] IntName, int Row, int Col)

/**
 * @brief      Sets the gray pixel.
 *
 * @param      Image  The image
 * @param[in]  Row    The row
 * @param[in]  Col    The col
 * @param[in]  value  The value
 */
void SetGrayPixel(int[] IntName, int Row, int Col,int value)
#endif

#endif /* PrefixHeader_matpixel */
