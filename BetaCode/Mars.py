import sys
import numpy as np 
import cv2
import matplotlib.pyplot as plt


# Convert from RGB to gray

def convertion_gray(img, Y):
	b,g,r = cv2.split(img)
	height = img.shape[0]
	width = img.shape[1]
	# For each pixel from the image, 
	# the intesity is obtained 
	# by the rgb values of that pixel
	for i in range(0, height):
		for j in range(0,width): 
			Y[i][j] = 0.299*r[i][j] + 0.587*g[i][j] + 0.114*b[i][j]


# Histogram equalization of gray

def equalization(img):
	height = img.shape[0]
	width = img.shape[1]
	Hist = np.zeros(400)
	CDF = np.zeros(256)
	Total_px = height*width

	# Gray scale level values histogram
	for i in range(0, height):
		for j in range(0,width): 
			Hist[img[i][j]] += 1

	# Gray scale histogram's CDF
	for i in range(0,255):	
		if i != 0:
			CDF[i] = Hist[i]/Total_px + CDF[i-1]

		elif i == 0:
			CDF[i] = Hist[0]/Total_px

	# Image equalization (^o^)
	for i in range(0, height):
		for j in range(0,width): 
			img[i][j] = CDF[img[i][j]]*255 

#Path algorithm

def path(img_col, img_gray):
	# Each arrays is a dot. 0 is y value, and 1 is x value.
	dot_type = [('y', int), ('x', int), ('gray_level', int), ('distance', float), ('px', int)]
	start = np.array((260, 415, img_gray[260][415], 0 , 0), dtype = dot_type)
	end = np.array((815, 1000, img_gray[815][1000], 0 , 0), dtype = dot_type)

	# Until reach the destination
	while (start['y'] != end['y'] or start['x'] != end['x']):
#		print start

		# Euclidean distance analysis 

		# Neighbour's list
		neighbour = np.zeros(8, dtype = dot_type)

		# Px = 0 start at top-left corner(px express the neighbour position)
		pixel = 0
		# Moves through x
		for i in range(0, 3):
			# Moves through y
			for j in range(0,3):
				# When i = j = 1 the pixel is the start point
				if (i != 1 or j != 1):
					# Neighbour info
	
					y = start['y']-1+j
					x = start['x']-1+i
					gray = img_gray[y][x]
					# Euclidean distance calculation 
					dist = np.zeros(2)
					dist = [y - end['y'], x - end['x']]
					dist = np.linalg.norm(dist)

					# Pixel start point neighbour

					neighbour[pixel] = (y, x, gray, dist, pixel)
					pixel += 1
#					print neighbour
		neighbour = np.sort(neighbour, order = 'distance')
		next = neighbour[0:3]

		# Gray level analysis
		next = np.sort(next, order = 'gray_level')

		# Mark path's dot
		img_col[start['y']][start['x']] = (0,0,0)
		start = next[0]



# Image loading and gray convertion
def main():
	Mars_col = cv2.imread('Mars.bmp', 1)
	height = Mars_col.shape[0]
	width = Mars_col.shape[1]

	Mars_gray = np.zeros((height,width), np.uint8)	

	convertion_gray(Mars_col, Mars_gray)

	equalization(Mars_gray)

	path(Mars_col, Mars_gray)

	cv2.imwrite('image.bmp', Mars_col)
	cv2.waitKey(0)
	cv2.destroyAllWindows()
	return 0

r = main()
 



