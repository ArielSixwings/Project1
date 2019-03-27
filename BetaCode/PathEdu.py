def path(img_col, img_gray):
    # Each arrays is a dot. 0 is y value, and 1 is x value.
    dot_type = [('y', int), ('x', int), ('gray_level', int), ('distance', float), ('px', int)]
    start = np.array((260, 415, img_gray[260][415], 0, 0), dtype=dot_type)
    end = np.array((815, 1000, img_gray[815][1000], 0, 0), dtype=dot_type)

    # Until reach the destination
    while (start['y'] != end['y'] or start['x'] != end['x']):
        #       print start

        # Euclidean distance analysis

        # Neighbour's list
        neighbour = np.zeros(8, dtype=dot_type)

        # Px = 0 start at top-left corner(px express the neighbour position)
        pixel = 0
        # Moves through x
        for i in range(0, 3):
            # Moves through y
            for j in range(0, 3):
                # When i = j = 1 the pixel is the start point
                if (i != 1 or j != 1):
                    # Neighbour info

                    y = start['y'] - 1 + j
                    x = start['x'] - 1 + i
                    gray = img_gray[y][x]
                    # Euclidean distance calculation
                    dist = np.zeros(2)
                    dist = [y - end['y'], x - end['x']]
                    dist = np.linalg.norm(dist)

                    # Pixel start point neighbour

                    neighbour[pixel] = (y, x, gray, dist, pixel)
                    pixel += 1
#                   print neighbour
        neighbour = np.sort(neighbour, order='distance')
        next = neighbour[0:3]

        # Gray level analysis
        next = np.sort(next, order='gray_level')

        # Mark path's dot
        img_col[start['y']][start['x']] = (0, 0, 0)
        start = next[0]
