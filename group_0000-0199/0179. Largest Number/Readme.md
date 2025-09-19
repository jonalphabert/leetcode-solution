# Largest Number

---

## 📝 Problem Statement (English)
Given a list of non-negative integers `nums`, arrange them such that they form the **largest number** and return it.  
Since the result may be very large, return it as a **string** instead of an integer.

---

## 📝 Deskripsi Soal (Bahasa Indonesia)
Diberikan sebuah array `nums` yang berisi bilangan bulat non-negatif. Susun elemen-elemen array tersebut sehingga membentuk **angka terbesar** yang mungkin.  
Karena hasilnya bisa sangat besar, kembalikan hasil dalam bentuk **string**, bukan integer.

---

## 🔒 Constraints
- `1 <= nums.length <= 100`
- `0 <= nums[i] <= 10^9`

---

## 📊 Example / Contoh

### Example 1
```

Input: nums = \[10, 2]
Output: "210"

```

### Example 2
```

Input: nums = \[3, 30, 34, 5, 9]
Output: "9534330"

```

### Example 3
```

Input: nums = \[0, 0]
Output: "0"

````

---

## 💡 Ideation / Pendekatan
1. Convert semua integer ke string agar bisa dibandingkan dengan cara concatenation.
2. Lakukan sorting dengan aturan khusus:
   - Untuk dua string `a` dan `b`, bandingkan `a+b` dengan `b+a`.
   - Jika `a+b > b+a`, maka `a` harus berada di depan `b`.
3. Setelah sorting, gabungkan semua string menjadi satu hasil.
4. Jika hasilnya dimulai dengan `"0"`, maka semua elemen adalah nol → return `"0"`.

---

## ⏱️ Time & Space Complexity
- **Time Complexity:**  
  - Sorting `n` elemen dengan comparator berbasis concatenation → `O(n log n * k)`  
    (`k` = rata-rata panjang string angka, maksimum 10).  
- **Space Complexity:**  
  - `O(n * k)` untuk menyimpan array string.  

---

## 🚀 Implementation (Go)
```go
package main

import (
    "fmt"
    "sort"
    "strconv"
    "strings"
)

func largestNumber(nums []int) string {
    strNums := make([]string, len(nums))
    for i, n := range nums {
        strNums[i] = strconv.Itoa(n)
    }

    sort.Slice(strNums, func(i, j int) bool {
        return strNums[i]+strNums[j] > strNums[j]+strNums[i]
    })

    result := strings.Join(strNums, "")
    if result[0] == '0' {
        return "0"
    }
    return result
}

func main() {
    fmt.Println(largestNumber([]int{10, 2}))        // "210"
    fmt.Println(largestNumber([]int{3, 30, 34, 5, 9})) // "9534330"
    fmt.Println(largestNumber([]int{0, 0}))         // "0"
}
````
