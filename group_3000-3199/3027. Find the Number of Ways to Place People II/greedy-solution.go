/*
	This solution uses a greedy approach to solve the problem.

	The idea is to sort the points by x-coordinate (and by y-coordinate in case of a tie).
	Then, for each point, we try to find valid pairs by maintaining boundaries for x and y coordinates.

	For each point A, we set initial boundaries:
	- xMin = A.x - 1
	- xMax = +infinity
	- yMin = -infinity
	- yMax = A.y + 1

	Then, for each point B, we check if B.x is between xMin and xMax, and B.y is between yMin and yMax.
	If yes, we update xMin, yMin, and increment the answer.
*/

import (
    "sort"
)

func numberOfPairs(points [][]int) int {
    ans := 0
    n := len(points)

    sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] { 
			return points[i][1] > points[j][1] 
		}
		return points[i][0] < points[j][0] 
	})

    for i, point := range points {
        xMin := point[0]-1
        xMax := math.MaxInt
        yMin := math.MinInt
        yMax := point[1]+1

        for j := i+1; j < n; j++ {
            xB, yB := points[j][0], points[j][1]
            if xB > xMin && xB < xMax && yB > yMin && yB < yMax {
                xMin = xB
                yMin = yB
                ans++
            }
        }
    }
    
    return ans
}