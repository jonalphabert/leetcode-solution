# Indonesia
## Problem

Diberikan `n` bilangan bulat tidak negatif yang merepresentasikan peta elevasi di mana lebar setiap bar adalah 1, hitung berapa banyak air yang dapat ditampung setelah hujan.

**Contoh:**
```
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Penjelasan: Diagram di bawah menunjukkan 6 unit air yang dapat ditampung.
```

## Solution

Terdapat dua pendekatan yang dapat digunakan untuk menyelesaikan masalah ini:

### 1. Prefix Suffix Sum

- Buat dua array: `arrMaxLeft` dan `arrMaxRight` untuk menyimpan maksimum di kiri dan kanan setiap indeks.
- Untuk setiap indeks, air yang dapat ditampung adalah minimum dari maksimum kiri dan kanan dikurangi tinggi pada indeks tersebut.
- Kompleksitas Waktu: O(n)
- Kompleksitas Memori: O(n)

**Contoh Implementasi (Go):**
```go
func trap(height []int) int {
    length := len(height)
    arrMaxRight := make([]int, len(height))
    arrMaxLeft := make([]int, len(height))

    arrMaxLeft[0] = height[0]
    arrMaxRight[length-1] = height[length-1]

    for i := 1; i < length; i++ {
        arrMaxLeft[i] = max(arrMaxLeft[i-1], height[i])
    }

    for i := length-2; i >= 0; i-- {
        arrMaxRight[i] = max(arrMaxRight[i+1], height[i])
    }

    result := 0
    for i := 0; i < length; i++ {
        result += min(arrMaxRight[i], arrMaxLeft[i]) - height[i]
    }

    return result
}
```

### 2. Monotonic Decreasing Stack

- Gunakan stack untuk menyimpan indeks bar yang belum menemukan batas kanan.
- Saat menemukan bar yang lebih tinggi, hitung air yang dapat ditampung di antara batas kiri dan kanan.
- Kompleksitas Waktu: O(n)
- Kompleksitas Memori: O(n)

**Contoh Implementasi (Go):**
```go
func trap(height []int) int {
    stack := []int{}
    result := 0

    for i := 0; i < len(height); i++ {
        for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
            top := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]

            if len(stack) == 0 {
                break
            }

            left := stack[len(stack) - 1]
            witdh := i - left - 1
            result += (min(height[left], height[i]) - height[top]) * witdh
        }

        stack = append(stack, i)
    }

    return result
}
```

---

# English
## Problem

Given `n` non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

**Example:**
```
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The diagram below shows 6 units of trapped water.
```

## Solution

There are two approaches to solve this problem:

### 1. Prefix Suffix Sum

- Create two arrays: `arrMaxLeft` and `arrMaxRight` to store the maximum to the left and right of each index.
- For each index, the trapped water is the minimum of the left and right maximums minus the height at that index.
- Time Complexity: O(n)
- Space Complexity: O(n)

**Sample Implementation (Go):**
```go
func trap(height []int) int {
    length := len(height)
    arrMaxRight := make([]int, len(height))
    arrMaxLeft := make([]int, len(height))

    arrMaxLeft[0] = height[0]
    arrMaxRight[length-1] = height[length-1]

    for i := 1; i < length; i++ {
        arrMaxLeft[i] = max(arrMaxLeft[i-1], height[i])
    }

    for i := length-2; i >= 0; i-- {
        arrMaxRight[i] = max(arrMaxRight[i+1], height[i])
    }

    result := 0
    for i := 0; i < length; i++ {
        result += min(arrMaxRight[i], arrMaxLeft[i]) - height[i]
    }

    return result
}
```

### 2. Monotonic Decreasing Stack

- Use a stack to store indices of bars that have not found a right boundary.
- When a higher bar is found, calculate the trapped water between the left and right boundaries.
- Time Complexity: O(n)
- Space Complexity: O(n)

**Sample Implementation (Go):**
```go
func trap(height []int) int {
    stack := []int{}
    result := 0

    for i := 0; i < len(height); i++ {
        for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
            top := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]

            if len(stack) == 0 {
                break
            }

            left := stack[len(stack) - 1]
            witdh := i - left - 1
            result += (min(height[left], height[i]) - height[top]) * witdh
        }

        stack = append(stack, i)
    }

    return result
}
```