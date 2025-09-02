/*
	This solution uses a prefix sum array to efficiently count the number of points within a given rectangle.
	It first constructs a grid to mark the presence of points, then builds a prefix sum array from this grid.
	For each pair of points, it checks if they can form a valid rectangle and uses the prefix sum array to count the number of points within that rectangle.
	If the count is exactly 2 (the two points themselves), it counts the pair as valid.
	But with the constraints, this approach is not efficient enough and brute force approach is better.
*/

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
    grid := make([][]int, 51)
    ps := make([][]int, 51)
    count := 0
    n := len(points)

    for i:= 0; i<51; i++ {
        grid[i] = make([]int, 51)
        ps[i] = make([]int, 51)
    }

    for _, point := range points {
        x := point[0]
        y := point[1]

        grid[x][y] = 1
    }

    for i:= 0; i<51; i++ {
        for j:= 0; j<51; j++ {
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