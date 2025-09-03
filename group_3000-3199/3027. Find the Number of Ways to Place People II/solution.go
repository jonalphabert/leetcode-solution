/*
	The solution is same as 3025. Find the Number of Ways to Place People I
	
	But because of the constraints, we need to compress the points first.
	The contraint is 10^9 <= points[i][0], points[i][1] <= 10^9
	
	If we use the same approach as 3025, we will need a grid of size 10^9 x 10^9 which is not feasible.
	
	So we can compress the points to a smaller range and then use the same approach as 3025.
	The compression can be done by sorting the points and then mapping them to a smaller range.
	For example, if we have points (1, 2), (3, 4), (5, 6), we can map them to (0, 0), (1, 1), (2, 2).

	After compression, we can use the same approach as 3025 to count the number of valid pairs.
*/

import (
    "sort"
)

func unique(arr []int) []int{
     if len(arr) == 0 {
        return arr
    }
    res := []int{arr[0]}
    for i := 1; i < len(arr); i++ {
        if arr[i] != arr[i-1] {
            res = append(res, arr[i])
        }
    }
    return res
}

func compressPoint(points [][]int) (map[int]int, map[int]int) {
    xSlice := make([]int, len(points))
    ySlice := make([]int, len(points))

    for i, point := range points{
        xSlice[i] = point[0]
        ySlice[i] = point[1]
    }

    sort.Ints(xSlice)
    xSlice = unique(xSlice)

    sort.Ints(ySlice)
    ySlice = unique(ySlice)

    setX := make(map[int]int)
    setY := make(map[int]int)

    for i, xPoint := range xSlice {
        setX[xPoint] = i
    }


    for i, yPoint := range ySlice {
        setY[yPoint] = i
    }

    return setX, setY
}

func query(ps [][]int, x1, y1, x2, y2 int) int {
    res := ps[x2][y2]
    if x1 > 0 {
        res -= ps[x1-1][y2]
    }
    if y1 > 0 {
        res -= ps[x2][y1-1]
    }
    if y1 > 0 && x1 > 0 {
        res += ps[x1-1][y1-1]
    }
    return res
}

func numberOfPairs(points [][]int) int {
    xMap, yMap := compressPoint(points)
    n := len(points)

    for i := 0; i < n; i++ {
        points[i][0], points[i][1] = xMap[points[i][0]], yMap[points[i][1]]
    }

    lenX := len(xMap)
    lenY := len(yMap)

    grid := make([][]int, lenX)
    ps := make([][]int, lenX)
    count := 0

    for i:= 0; i<lenX; i++ {
        grid[i] = make([]int, lenY)
        ps[i] = make([]int, lenY)
    }

    for _, point := range points {
        x := point[0]
        y := point[1]

        grid[x][y] = 1
    }

    for i:= 0; i<lenX; i++ {
        for j:= 0; j<lenY; j++ {
            ps[i][j] = grid[i][j]
            if i > 0 {
                ps[i][j] += ps[i-1][j]
            }
            if j > 0 {
                ps[i][j] += ps[i][j-1]
            }
            if i > 0 && j > 0 {
                ps[i][j] -= ps[i-1][j-1]
            }
        }
    }

    for i:= 0; i < n; i++ {
        for j := 0; j < n ; j++ {
            if i == j { continue }
			Ax, Ay := points[i][0], points[i][1]
			Bx, By := points[j][0], points[j][1]

			if !(Ax <= Bx && Ay >= By) {
				continue
			}

			x1, x2 := Ax, Bx
			y1, y2 := By, Ay

			if query(ps, x1, y1, x2, y2) == 2 {
				count++
			}
        }
    }

    return count
}