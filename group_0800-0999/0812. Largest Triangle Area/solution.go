func triangleArea(a, b, c []int) float64 {
	return math.Abs(float64(a[0]*(b[1]-c[1]) +
		b[0]*(c[1]-a[1]) +
		c[0]*(a[1]-b[1]))) / 2.0
}


func largestTriangleArea(points [][]int) float64 {
    n := len(points)
    maxArea := float64(0)

    for i:=0; i<n; i++{
        for j:=i+1; j<n; j++ {
            for k:=j+1; k<n; k++ {
                area := triangleArea(points[i], points[j], points[k])
				if area > maxArea {
					maxArea = area
				}
            }
        }
    }

    return maxArea
}