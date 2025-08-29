func largestRectangleArea(heights []int) int {
    stack := []int{}
    maxArea := 0
    n := len(heights)

	// Iterasi sampai n supaya ada iterasi tambahan untuk memproses sisa elemen di stack
    for i := 0; i <= n; i++ {
		//  Jika iterasi sudah mencapai n, set curHeight ke 0 untuk memproses sisa elemen di stack
        curHeight := 0
        if i < n {
            curHeight = heights[i]
        }

        for len(stack) > 0 && curHeight < heights[stack[len(stack)-1]] {
            h := heights[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]

			// Width by default adalah i ketika stack kosong
			// Karena artinya semua elemen akan memiliki tinggi minimal ketika stack kosong
            width := i
            if len(stack) > 0 {
                width = i - stack[len(stack)-1] - 1
            }

            area := h * width
            if area > maxArea {
                maxArea = area
            }
        }
        stack = append(stack, i)
    }

    return maxArea
}