## Problem

**Indonesia:**  
Diberikan sebuah array 2D `points` berukuran n x 2 yang merepresentasikan koordinat integer dari beberapa titik pada bidang 2D, di mana `points[i] = [xi, yi]`.

Hitung jumlah pasangan titik `(A, B)` yang memenuhi:
- Titik A berada di kiri atas titik B (A.x ≤ B.x dan A.y ≥ B.y).
- Tidak ada titik lain di dalam persegi panjang (atau garis) yang dibentuk oleh A dan B (termasuk pada sisi/border).

Kembalikan jumlah pasangan yang valid.

**English:**  
You are given a 2D array `points` of size n x 2 representing integer coordinates of some points on a 2D plane, where `points[i] = [xi, yi]`.

Count the number of pairs of points `(A, B)` where:
- A is on the upper left side of B (A.x ≤ B.x and A.y ≥ B.y).
- There are no other points in the rectangle (or line) they make (including the border).

Return the count.

---

## Solution

### Brute Force

**Indonesia:**  
Iterasi semua pasangan titik `(A, B)` yang memenuhi syarat posisi (A.x ≤ B.x dan A.y ≥ B.y).  
Untuk setiap pasangan, cek semua titik lain apakah ada yang berada di dalam persegi panjang yang dibentuk oleh A dan B (termasuk sisi/border).  
Jika tidak ada titik lain di dalam area tersebut, pasangan dianggap valid.

**English:**  
Iterate through all pairs of points `(A, B)` that satisfy the position condition (A.x ≤ B.x and A.y ≥ B.y).  
For each pair, check all other points to see if any are inside the rectangle formed by A and B (including the border).  
If no other point is inside, count the pair as valid.

**Kompleksitas Waktu / Time Complexity:** O(n³)

---

### Prefix Sum Solution

**Indonesia:**  
Buat grid untuk menandai keberadaan titik, lalu bangun prefix sum 2D dari grid tersebut.  
Untuk setiap pasangan `(A, B)` yang memenuhi syarat posisi, gunakan prefix sum untuk menghitung jumlah titik di dalam persegi panjang yang dibentuk oleh A dan B.  
Jika jumlah titik di area tersebut hanya 2 (A dan B sendiri), pasangan dianggap valid.

**English:**  
Create a grid to mark the presence of points, then build a 2D prefix sum from the grid.  
For each pair `(A, B)` that satisfies the position condition, use the prefix sum to count the number of points inside the rectangle formed by A and B.  
If the count is exactly 2 (just A and B), the pair is valid.

**Kompleksitas Waktu / Time Complexity:** O(n² + M²), dengan M adalah rentang koordinat.

---

**Catatan:**  
- Solusi brute force lebih sederhana namun lambat untuk n besar.
- Solusi prefix sum lebih efisien untuk koordinat kecil, namun tetap bisa lambat