func move(c byte) (int, int) {
    switch c {
    	case 'U':
    		return 0, 1
    	case 'D':
    		return 0, -1
    	case 'L':
    		return -1, 0
    	case 'R':
    		return 1, 0
    	}
    	return 0, 0
    }

func distinctPoints(s string, k int) int {
    n := len(s)
	px := make([]int, n+1)
	py := make([]int, n+1)

	for i := 0; i < n; i++ {
		dx, dy := move(s[i])
		px[i+1] = px[i] + dx
		py[i+1] = py[i] + dy
	}

	X, Y := px[n], py[n]
	seen := make(map[[2]int]struct{})

	for l := 0; l <= n-k; l++ {
		dx := px[l+k] - px[l]
		dy := py[l+k] - py[l]
		final := [2]int{X - dx, Y - dy}
		seen[final] = struct{}{}
	}

	return len(seen)
}