## 🇮🇩 Bahasa Indonesia

### 📌 Deskripsi Masalah

Diberikan sebuah array segitiga (triangle array), kembalikan jumlah jalur minimum dari atas ke bawah.

Untuk setiap langkah, Anda dapat berpindah ke angka yang adjacent (berdekatan) pada baris di bawahnya. Lebih formalnya, jika Anda berada pada indeks `i` di baris saat ini, Anda dapat berpindah ke indeks `i` atau indeks `i + 1` pada baris berikutnya.

### 🔍 Abstraksi

Masalah ini dapat diselesaikan dengan menggunakan pendekatan **dynamic programming** (pemrograman dinamis). Ide utamanya adalah menghitung jalur minimum untuk setiap posisi dalam segitiga dengan memanfaatkan hasil perhitungan dari baris sebelumnya. Untuk setiap elemen pada baris saat ini, kita menjumlahkannya dengan nilai minimum dari dua elemen yang adjacent di baris sebelumnya.

### 💡 Implementasi

Solusi ini menggunakan pendekatan **bottom-up** dengan memodifikasi array segitiga asli:

1. Iterasi melalui setiap baris segitiga mulai dari baris kedua
2. Untuk setiap elemen dalam baris, hitung nilai minimum dari dua elemen adjacent di baris sebelumnya
3. Tambahkan nilai minimum tersebut ke elemen saat ini
4. Setelah memproses semua baris, cari nilai minimum pada baris terakhir

```go
func minimumTotal(triangle [][]int) int {
    minValue := math.MaxInt

    for i := 1; i < len(triangle) ; i++ {
        for j := 0; j < len(triangle[i]); j++ {
            nums1 := math.MaxInt
            nums2 := math.MaxInt
            if j != 0 {
                nums1 = triangle[i-1][j-1]
            }
            if j != len(triangle[i])-1 {
                nums2 = triangle[i-1][j]
            }

            triangle[i][j] += min(nums1, nums2)
        }
    }

    for _, val := range triangle[len(triangle) -1] {
        minValue = min(minValue, val)
    }

    return minValue
}
```

### ⏱️ Kompleksitas

- **Time Complexity**: O(n²) dimana n adalah jumlah baris dalam segitiga, karena kita perlu mengiterasi melalui setiap elemen dalam segitiga
- **Space Complexity**: O(1) karena kita memodifikasi array input secara langsung tanpa menggunakan struktur data tambahan

---

## 🇬🇧 English

### 📌 Problem Description

Given a triangle array, return the minimum path sum from top to bottom.

For each step, you may move to an adjacent number of the row below. More formally, if you are on index i on the current row, you may move to either index i or index i + 1 on the next row.

### 🔍 Abstraction

This problem can be solved using a **dynamic programming** approach. The main idea is to calculate the minimum path for each position in the triangle by leveraging the results from the previous row. For each element in the current row, we add it to the minimum value of the two adjacent elements from the previous row.

### 💡 Implementation

This solution uses a **bottom-up** approach by modifying the original triangle array:

1. Iterate through each row of the triangle starting from the second row
2. For each element in the row, calculate the minimum value of the two adjacent elements from the previous row
3. Add this minimum value to the current element
4. After processing all rows, find the minimum value in the last row

```go
func minimumTotal(triangle [][]int) int {
    minValue := math.MaxInt

    for i := 1; i < len(triangle) ; i++ {
        for j := 0; j < len(triangle[i]); j++ {
            nums1 := math.MaxInt
            nums2 := math.MaxInt
            if j != 0 {
                nums1 = triangle[i-1][j-1]
            }
            if j != len(triangle[i])-1 {
                nums2 = triangle[i-1][j]
            }

            triangle[i][j] += min(nums1, nums2)
        }
    }

    for _, val := range triangle[len(triangle) -1] {
        minValue = min(minValue, val)
    }

    return minValue
}
```

### ⏱️ Complexity

- **Time Complexity**: O(n²) where n is the number of rows in the triangle, as we need to iterate through each element in the triangle
- **Space Complexity**: O(1) because we modify the input array directly without using additional data structures