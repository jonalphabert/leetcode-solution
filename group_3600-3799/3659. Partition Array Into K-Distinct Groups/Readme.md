## Problem

**Indonesia:**  
Diberikan sebuah array integer `nums` dan sebuah bilangan bulat `k`.

Tentukan apakah memungkinkan untuk membagi semua elemen `nums` ke dalam satu atau lebih grup sehingga:
- Setiap grup berisi tepat `k` elemen.
- Semua elemen dalam setiap grup berbeda (unik).
- Setiap elemen di `nums` harus masuk ke tepat satu grup.

Kembalikan `true` jika pembagian seperti itu memungkinkan, jika tidak kembalikan `false`.

**English:**  
You are given an integer array `nums` and an integer `k`.

Determine whether it is possible to partition all elements of `nums` into one or more groups such that:
- Each group contains exactly `k` elements.
- All elements in each group are distinct.
- Each element in `nums` must be assigned to exactly one group.

Return `true` if such a partition is possible, otherwise return `false`.

---

## Solution

### Bahasa Indonesia

- Jumlah total elemen harus habis dibagi `k`, jika tidak, langsung return `false`.
- Setiap nilai di `nums` tidak boleh muncul lebih dari jumlah grup yang akan dibentuk (`len(nums) / k`), karena setiap grup harus berisi elemen unik.
- Hitung frekuensi setiap elemen, jika ada yang melebihi jumlah grup, return `false`.
- Jika semua syarat terpenuhi, return `true`.

### English

- The total number of elements must be divisible by `k`; otherwise, return `false`.
- Each value in `nums` cannot appear more than the number of groups (`len(nums) / k`), since each group must have distinct elements.
- Count the frequency of each element; if any frequency exceeds the number of groups, return `false`.
- If all conditions are met, return `true`.

**Sample Implementation (Go):**
```go
func partitionArray(nums []int, k int) bool {
    if len(nums) % k != 0 {
        return false
    }

    group := len(nums) / k
    freq := make(map[int]int)

    for _, val := range nums {
        freq[val]++
        if freq[val] > group {
            return false
        }
    }
    
    return true
}
```