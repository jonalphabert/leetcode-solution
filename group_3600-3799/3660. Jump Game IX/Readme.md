## Problem

**Indonesia:**  
Diberikan sebuah array integer `nums`.

Dari setiap indeks `i`, kamu dapat melompat ke indeks lain `j` dengan aturan berikut:
- Melompat ke indeks `j > i` hanya diperbolehkan jika `nums[j] < nums[i]`.
- Melompat ke indeks `j < i` hanya diperbolehkan jika `nums[j] > nums[i]`.

Untuk setiap indeks `i`, cari nilai maksimum di `nums` yang dapat dicapai dengan mengikuti urutan lompatan yang valid mulai dari indeks `i`.

Kembalikan array `ans` di mana `ans[i]` adalah nilai maksimum yang dapat dicapai dari indeks `i`.

**English:**  
You are given an integer array `nums`.

From any index `i`, you can jump to another index `j` under the following rules:
- Jump to index `j > i` is allowed only if `nums[j] < nums[i]`.
- Jump to index `j < i` is allowed only if `nums[j] > nums[i]`.

For each index `i`, find the maximum value in `nums` that can be reached by following any sequence of valid jumps starting at `i`.

Return an array `ans` where `ans[i]` is the maximum value reachable starting from index `i`.

---

## Solution

### Bahasa Indonesia

Solusi menggunakan **prefix maksimum** dan **suffix minimum** untuk menemukan "cut" atau pemisah komponen:
- Hitung prefix maksimum (`prefMax`) dan suffix minimum (`suffMin`) untuk setiap indeks.
- Cari indeks cut di mana `prefMax[i] <= suffMin[i+1]`. Cut ini memastikan semua elemen di kiri cut lebih kecil atau sama dengan semua elemen di kanan cut.
- Untuk setiap blok yang dipisahkan oleh cut, semua elemen di blok tersebut dapat diubah menjadi nilai maksimum di blok tersebut tanpa melanggar aturan lompatan.
- Hasil akhir adalah setiap elemen di blok diisi dengan nilai maksimum blok tersebut.

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Memori:** O(n)

### English

The solution uses **prefix maximum** and **suffix minimum** to find "cuts" or component separators:
- Compute prefix maximum (`prefMax`) and suffix minimum (`suffMin`) for each index.
- Find cut indices where `prefMax[i] <= suffMin[i+1]`. This ensures all elements to the left of the cut are less than or equal to all elements to the right.
- For each block separated by cuts, all elements in the block can be changed to the maximum value in that block without violating the jump rules.
- The final result is each element in the block is set to the maximum value of that block.

**Time Complexity:** O(n)  
**Space Complexity:** O(n)

**Sample Implementation (Go):**
```go
func maxValue(nums []int) []int {
    n := len(nums)
    if n == 0 {
        return []int{}
    }

    // prefix maximum
    prefMax := make([]int, n)
    prefMax[0] = nums[0]
    for i := 1; i < n; i++ {
        if nums[i] > prefMax[i-1] {
            prefMax[i] = nums[i]
        } else {
            prefMax[i] = prefMax[i-1]
        }
    }

    // suffix minimum
    suffMin := make([]int, n)
    suffMin[n-1] = nums[n-1]
    for i := n - 2; i >= 0; i-- {
        if nums[i] < suffMin[i+1] {
            suffMin[i] = nums[i]
        } else {
            suffMin[i] = suffMin[i+1]
        }
    }

    // find cuts
    cuts := []int{}
    for i := 0; i < n-1; i++ {
        if prefMax[i] <= suffMin[i+1] {
            cuts = append(cuts, i)
        }
    }

    // assign result
    ans := make([]int, n)
    start := 0
    cuts = append(cuts, n-1)
    for _, c := range cuts {
        maxVal := nums[start]
        for j := start; j <= c; j++ {
            if nums[j] > maxVal {
                maxVal = nums[j]
            }
        }
        for j := start; j <= c; j++ {
            ans[j] = maxVal
        }
        start = c + 1
    }

    return ans
}
```