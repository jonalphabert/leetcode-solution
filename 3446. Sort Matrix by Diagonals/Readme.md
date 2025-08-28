# Indonesia
## Problem

Diberikan sebuah matriks persegi `grid` berukuran n x n yang berisi bilangan bulat.  
Kembalikan matriks sehingga:

- Diagonal-diagonal pada segitiga kiri bawah (termasuk diagonal tengah) diurutkan dalam urutan tidak meningkat (menurun).
- Diagonal-diagonal pada segitiga kanan atas diurutkan dalam urutan tidak menurun (menaik).

**Contoh:**
```
Input: grid = [
  [3, 1, 2],
  [2, 4, 5],
  [1, 2, 3]
]
Output: [
  [3, 2, 2],
  [4, 4, 1],
  [5, 3, 1]
]
```

## Solution

Solusi dilakukan dengan langkah-langkah berikut:

1. Kelompokkan elemen berdasarkan diagonalnya menggunakan indeks `i - j`.
2. Untuk setiap diagonal:
   - Jika diagonal berada di segitiga bawah (key >= 0), urutkan secara menaik (ascending).
   - Jika diagonal berada di segitiga atas (key < 0), urutkan secara menurun (descending).
3. Rekonstruksi matriks dengan mengambil elemen dari belakang array diagonal agar urutan sesuai dengan ketentuan.

**Kompleksitas Waktu:** O(n² log n)  
**Kompleksitas Memori:** O(n²)


---

# English
## Problem

You are given an n x n square matrix of integers `grid`.  
Return the matrix such that:

- The diagonals in the bottom-left triangle (including the middle diagonal) are sorted in non-increasing order.
- The diagonals in the top-right triangle are sorted in non-decreasing order.

**Example:**
```
Input: grid = [
  [3, 1, 2],
  [2, 4, 5],
  [1, 2, 3]
]
Output: [
  [3, 2, 2],
  [4, 4, 1],
  [5, 3, 1]
]
```

## Solution

The solution involves these steps:

1. Group elements by their diagonal using the index `i - j`.
2. For each diagonal:
   - If the diagonal is in the bottom triangle (`key >= 0`), sort in ascending order.
   - If the diagonal is in the top triangle (`key < 0`), sort in descending order.
3. Reconstruct the matrix by taking elements from the end of each diagonal array to ensure the required order.

**Time Complexity:** O(n² log n)  
**Space Complexity:** O(n²)