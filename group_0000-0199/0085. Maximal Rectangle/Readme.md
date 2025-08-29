## Problem

**Indonesia:**  
Diberikan sebuah matriks biner berukuran `rows x cols` yang berisi angka 0 dan 1.  
Temukan luas maksimum dari sebuah persegi panjang yang hanya berisi angka 1 dan kembalikan luasnya.

**English:**  
Given a `rows x cols` binary matrix filled with 0's and 1's, find the largest rectangle containing only 1's and return its area.

**Contoh:**
```
Input: matrix = [
  ['1','0','1','0','0'],
  ['1','0','1','1','1'],
  ['1','1','1','1','1'],
  ['1','0','0','1','0']
]
Output: 6
Penjelasan: Persegi panjang terbesar berisi 1 memiliki area 6.
```

---

## Solution

### Bahasa Indonesia

Solusi menggunakan pendekatan **histogram** dan **monotonic stack**:

1. Untuk setiap baris, buat array `height` yang merepresentasikan tinggi berturut-turut angka 1 di setiap kolom.
2. Setiap kali menemukan '1', tambahkan tinggi pada kolom tersebut; jika '0', reset ke 0.
3. Setelah membentuk histogram untuk baris tersebut, gunakan algoritma "Largest Rectangle in Histogram" (monotonic stack) untuk mencari area maksimum pada histogram tersebut.
4. Ulangi untuk setiap baris dan simpan area maksimum yang ditemukan.

**Kompleksitas Waktu:** O(rows * cols)  
**Kompleksitas Memori:** O(cols)

**Implementasi (Go):**
```go
func maximalRectangle(matrix [][]byte) int {
    maxArea := 0
    rows, cols := len(matrix), len(matrix[0])
    height := make([]int, cols)

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if matrix[i][j] == '0' {
                height[j] = 0
            } else {
                height[j]++
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

    for i := 0; i <= n; i++ {
        curHeight := 0
        if i < n {
            curHeight = heights[i]
        }

        for len(stack) > 0 && curHeight < heights[stack[len(stack)-1]] {
            h := heights[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]

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
```

---

### English

The solution uses the **histogram** and **monotonic stack** approach:

1. For each row, build a `height` array representing the consecutive number of 1's in each column.
2. For every '1', increment the height at that column; for '0', reset to 0.
3. After building the histogram for the row, use the "Largest Rectangle in Histogram" algorithm (monotonic stack) to find the maximal area in that histogram.
4. Repeat for each row and keep track of the maximum area found.

**Time Complexity:** O(rows * cols)  
**Space Complexity:** O(cols)

**Implementation (Go):**
```go
func maximalRectangle(matrix [][]byte) int {
    maxArea := 0
    rows, cols := len(matrix), len(matrix[0])
    height := make([]int, cols)

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if matrix[i][j] == '0' {
                height[j] = 0
            } else {
                height[j]++
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

    for i := 0; i <= n; i++ {
        curHeight := 0
        if i < n {
            curHeight = heights[i]
        }

        for len(stack) > 0 && curHeight < heights[stack[len(stack)-1]] {
            h := heights[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]

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
```