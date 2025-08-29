## Problem

**Indonesia:**  
Diberikan sebuah array integer `heights` yang merepresentasikan tinggi bar pada histogram, di mana lebar setiap bar adalah 1.  
Kembalikan luas maksimum dari sebuah persegi panjang yang dapat terbentuk di dalam histogram tersebut.

**English:**  
Given an array of integers `heights` representing the histogram's bar height where the width of each bar is 1, return the area of the largest rectangle in the histogram.

**Contoh:**
```
Input: heights = [2,1,5,6,2,3]
Output: 10
Penjelasan: Persegi panjang terbesar memiliki area 10 (bar ke-3 dan ke-4, tinggi 5 dan 6).
```

---

## Solution

### Bahasa Indonesia

Solusi menggunakan **Monotonic Stack**:

- Stack menyimpan indeks bar yang tinggi-nya belum menemukan batas kanan yang lebih rendah.
- Saat menemukan bar yang lebih pendek, keluarkan indeks dari stack dan hitung area dengan bar tersebut sebagai tinggi minimum.
- Lebar dihitung dari indeks saat ini ke indeks sebelumnya di stack.
- Tambahkan iterasi ekstra dengan tinggi 0 di akhir untuk memastikan semua bar diproses.

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Memori:** O(n)

**Implementasi (Go):**
```go
func largestRectangleArea(heights []int) int {
    stack := []int{}
    maxArea := 0
    n := len(heights)

    // Iterasi sampai n supaya ada iterasi tambahan untuk memproses sisa elemen di stack
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

The solution uses a **Monotonic Stack**:

- The stack stores indices of bars whose heights have not yet found a right boundary that is lower.
- When a shorter bar is found, pop indices from the stack and calculate the area using the popped bar as the minimum height.
- The width is calculated from the current index to the previous index in the stack.
- An extra iteration with height 0 at the end ensures all bars are processed.

**Time Complexity:** O(n)  
**Space Complexity:** O(n)

**Implementation (Go):**
```go
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