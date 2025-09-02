/*
	This solution is the brute force approach to solve the problem.
	It iterates through all pairs of points and checks if any other point lies within the rectangle formed by the two points.
	If no other point lies within the rectangle, it counts the pair as valid.
*/

import (
    "sort"
)

func numberOfPairs(points [][]int) int {
    count := 0
    n := len(points)

    sort.Slice(points, func(i, j int) bool {
		if points[i][1] == points[j][1] { 
			return points[i][0] < points[j][0] 
		}
		return points[i][1] > points[j][1] 
	})

    for i := 0 ; i < n ; i++ {
        for j := i+1 ; j < n ; j++ {
            if points[i][0] > points[j][0] && points[i][1] > points[j][1] {
                continue
            }
            isValid := true

            for k := 0 ; k < n; k++ {
                if k == i || k == j {
                    continue
                }

                if points[k][0] >= points[i][0] && points[k][0] <= points[j][0] && points[k][1] >= points[j][1] && points[k][1] <= points[i][1] {
                    isValid = false
                    break
                }
            }

            if isValid {
                count ++;
            }
        }
    }

    return count
}