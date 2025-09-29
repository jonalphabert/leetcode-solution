# 📖 Readme

## [Bahasa Indonesia] - [1039. Minimum Score Triangulation of Polygon](https://leetcode.com/problems/minimum-score-triangulation-of-polygon/description/?envType=daily-question&envId=2025-09-29)

### Deskripsi 

Diberikan convex poligon `n` sisi yang dimana semua vertex memiliki biaya yang disimpan di array `values[index]` dengan urutan sesuai dengan arah jarum jam.

Polygon Triangulation adalah proses dimana kamu dapat membagi sebuah polygon menjadi beberapa segitiga sehingga segitiga yang kamu buat akan membentuk polygon lagi. Hasil dari proses ini akan selalu menghasilkan `n-2` segitiga pembentuk polygon

Tugasmu adalah membuat proses polygon triangulation dengan biaya paling sedikit

--- 

### Intuisi 

- Ambil segitiga dari titik `i` dan `j`
- Lalu ambil titik `k` dimana `i < k < j` sehingga memposisikan `k` sebagai titik tengah
- Jika kita ambil titik `k` maka akan ada area yang belum tercakup segitiga di `[i...k]` dan `[k...j]`
- Maka kita perlu mencari solusi optimal di kedua area tersebut kembali
- Jika dilakukan secara brute force maka sangat memungkinkan ditemukan komputasi berulang. Sehingga diperlukan dynamic programming.
- Kenapa dynamic programming cocok dengan problem ini? Karena kita membutuhkan pencarian nilai optimal suatu problem dari nilai optimal sub-problem

--- 

### Pendekatan 

- Definisikan sebuah array `dp[n][n]`
- Iterasi `i` dari `n` ke `0`
- Iterasikan `j` dari `i +2` hingga `n`
- Lalu iterasikan `k` dari `i+1` ke `j-1`
    - Inisiasikan `dp[i][j]` dengan nilai maksimum
    - Hitung nilai optimum `dp[i][j]` dengan mencari nilai minimum antara `dp[i][j]` dengan `min(dp[i][k] + dp[k][j] + values[i] * values[j] * values[k])`
- Kembalikan nilai `dp[0][n-1]`

--- 

### Kompleksitas 
- **Waktu:** O(N^3) 
- **Memori:** O(N^2)
--- 

## [English] - [1039. Minimum Score Triangulation of Polygon](https://leetcode.com/problems/minimum-score-triangulation-of-polygon/description/?envType=daily-question&envId=2025-09-29)

### Description 

You have a convex n-sided polygon where each vertex has an integer value. You are given an integer array values where values[i] is the value of the ith vertex in clockwise order.

Polygon triangulation is a process where you divide a polygon into a set of triangles and the vertices of each triangle must also be vertices of the original polygon. Note that no other shapes other than triangles are allowed in the division. This process will result in n - 2 triangles.

You will triangulate the polygon. For each triangle, the weight of that triangle is the product of the values at its vertices. The total score of the triangulation is the sum of these weights over all n - 2 triangles.

Return the minimum possible score that you can achieve with some triangulation of the polygon.

--- 

### Intuition 

- Take two point `i` and `j`
- Then take point `k` where `i < k < j` to make `k` the center point
- If we take point `k` then there will be area that not yet covered by triangle in `[i...k]` and `[k...j]`
- So we need to find optimal solution in both area again
- If we do it in brute force way, there is a possibility of repeated computations. So we need dynamic programming.
- Why dynamic programming is suitable for this problem? Because we need to find optimal value of a problem from optimal value of sub-problem

--- 

### Approach 

- Define array `dp[i][j]`
- Iterate `i` from `n` to `0`
- Iterate `j` from `i + 2` to `n`
- Then iterate `k` from `i + 1` to `j - 1`
    - Initialize `dp[i][j]` with maximum value
    - Calculate optimum value of `dp[i][j]` by finding minimum value between `dp[i][j]` and `min(dp[i][k] + dp[k][j] + values[i] * values[j] * values[k])`
- Return `dp[0][n-1]`

--- 

### Complexity 
- **Time Complexity:** O(N^3) 
- **Auxiliary Space Complexity:** O(N^2)
--- 
