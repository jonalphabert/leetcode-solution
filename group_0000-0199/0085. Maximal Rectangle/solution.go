func maximalRectangle(matrix [][]byte) int {
    maxArea := 0
    rows, cols := len(matrix), len(matrix[0])
    height := make([]int, cols)

	// Iterasi setiap baris untuk mendapatkan tinggi histogram di setiap kolom
	// Lalu hitung luas maksimum dari histogram tersebut dengan fungsi yang sama di solution 84
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if matrix[i][j] == '0' {
                height[j] = 0;
            } else if matrix[i][j] == '1' {
                height[j]++;
            }
        }
        maxArea = max(maxArea, largestRectangleArea(height))
    }

    return maxArea
}

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