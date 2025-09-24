## 📖 \[Indonesia] - N-Queens Problem

### Deskripsi

* **Konteks masalah:** N-Queens adalah problem klasik dalam ilmu komputer dan algoritma backtracking. Tantangannya adalah menempatkan `n` ratu pada papan catur berukuran `n x n` sehingga tidak ada dua ratu yang saling menyerang.
* **Aturan:** Ratu dapat menyerang dalam satu baris, kolom, atau diagonal.
* **Input:** Sebuah integer `n` yang merepresentasikan ukuran papan.
* **Output:** Semua kemungkinan konfigurasi papan `n x n` dalam bentuk array string, di mana:

  * `'Q'` merepresentasikan posisi ratu.
  * `'.'` merepresentasikan sel kosong.
* **Kondisi khusus:**

  * Jika `n = 2` atau `n = 3`, tidak ada solusi yang valid.
  * Hasil boleh dikembalikan dalam urutan apa pun.

---

### Intuisi

* Kunci utama adalah **backtracking**: mencoba menaruh ratu satu per satu pada setiap baris, lalu memeriksa apakah penempatan valid.
* Jika ada konflik (kolom/diagonal sudah ditempati), posisi tersebut dilewati.
* Setelah menaruh semua ratu, konfigurasi valid disimpan sebagai salah satu solusi.

---

### Pendekatan

1. Gunakan **rekursi + backtracking**:

   * Letakkan ratu di baris `i`.
   * Cek apakah posisi tersebut valid (tidak ada ratu di kolom/diagonal).
   * Jika valid, lanjut ke baris berikutnya.
   * Jika tidak valid, backtrack dan coba posisi lain.
2. Representasikan papan sebagai `[][]byte`, lalu convert ke `[]string` untuk hasil akhir.
3. Simpan semua solusi dalam array `[][]string`.

---

### Kompleksitas

* **Waktu:** `O(N!)` → karena pada dasarnya mencoba berbagai kombinasi posisi ratu.
* **Memori:** `O(N^2)` untuk menyimpan papan sementara + `O(N)` untuk rekursi stack.
  Optimasi dapat dilakukan dengan array tambahan untuk tracking kolom dan diagonal.

---

### Contoh

**Input:**

```
n = 4
```

**Output:**

```
[
 [".Q..",
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",
  "Q...",
  "...Q",
  ".Q.."]
]
```

---

---

## 📖 \[English] - N-Queens Problem

### Description

* **Context:** The N-Queens puzzle is a classic problem in computer science and backtracking. The challenge is to place `n` queens on an `n x n` chessboard so that no two queens attack each other.
* **Rules:** A queen can attack in the same row, column, or diagonal.
* **Input:** An integer `n` representing the size of the chessboard.
* **Output:** All possible board configurations as arrays of strings, where:

  * `'Q'` represents a queen.
  * `'.'` represents an empty cell.
* **Special Cases:**

  * If `n = 2` or `n = 3`, there are no valid solutions.
  * The result can be returned in any order.

---

### Intuition

* The key is **backtracking**: try to place queens one by one in each row while ensuring that the placement is valid.
* If a conflict occurs (another queen in the same column/diagonal), skip that position.
* Once all queens are placed, store the configuration as a solution.

---

### Approach

1. Use **recursion + backtracking**:

   * Place a queen in row `i`.
   * Check if the position is valid (no queens in column or diagonals).
   * If valid, move to the next row.
   * If invalid, backtrack and try another column.
2. Represent the board as `[][]byte`, then convert it to `[]string` for the final output.
3. Collect all solutions in `[][]string`.

---

### Complexity

* **Time:** `O(N!)` → since many permutations of queen placements must be tried.
* **Space:** `O(N^2)` for storing the board + `O(N)` for recursion stack.
  Optimizations can use arrays to track column and diagonal usage.

---

### Example

**Input:**

```
n = 4
```

**Output:**

```
[
 [".Q..",
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",
  "Q...",
  "...Q",
  ".Q.."]
]
```