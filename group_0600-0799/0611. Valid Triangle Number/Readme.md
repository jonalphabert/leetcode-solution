# 📖 README

## 🇮🇩 Bahasa Indonesia - [Valid Triangle Number]

### Deskripsi

* **Konteks masalah:** Soal ini berasal dari LeetCode dengan judul *Valid Triangle Number*.
* **Aturan:** Sebuah triplet $(a, b, c)$ dapat membentuk segitiga jika memenuhi aturan segitiga: jumlah dua sisi terkecil harus lebih besar dari sisi terbesar.
* **Input:** Array integer `nums`.
* **Output:** Jumlah triplet valid yang dapat membentuk segitiga.
* **Kondisi khusus:** Panjang array `1 <= nums.length <= 1000`, dan nilai `0 <= nums[i] <= 1000`.

---

### Intuisi

* Sort array untuk mempermudah validasi syarat segitiga.
* Dengan asumsi $a \leq b \leq c$, cukup cek apakah $a + b > c$.
* Gunakan pendekatan *two pointers* agar lebih efisien dibanding binary search.

---

### Pendekatan

1. Urutkan array `nums`.
2. Fix sisi terbesar `c` dengan loop dari kanan ke kiri.
3. Gunakan dua pointer `l` dan `r` untuk mencari jumlah pasangan valid di kiri `c`.

   * Jika `nums[l] + nums[r] > nums[i]` → semua kombinasi dari `l` sampai `r-1` valid → tambahkan `(r - l)`.
   * Jika tidak, geser `l` ke kanan.
4. Ulangi hingga semua triplet dicek.

---

### Kompleksitas

* **Waktu:** $O(n^2)$, karena untuk setiap anchor kita menggunakan dua pointer linear.
* **Memori:** $O(1)$, hanya menggunakan variabel tambahan.

---

### Contoh

**Input:**

```
nums = [2, 2, 3, 4]
```

**Output:**

```
3
```

**Penjelasan:** Triplet valid adalah `[2,3,4]` (dua kali, karena ada dua angka 2), dan `[2,2,3]`.

---

---

## 🇬🇧 English - [Valid Triangle Number]

### Description

* **Context:** This problem comes from LeetCode under the title *Valid Triangle Number*.
* **Rule:** A triplet $(a, b, c)$ can form a triangle if the sum of the two smaller sides is greater than the largest side.
* **Input:** An integer array `nums`.
* **Output:** The number of valid triplets that can form a triangle.
* **Special condition:** Array length `1 <= nums.length <= 1000`, and values `0 <= nums[i] <= 1000`.

---

### Intuition

* Sort the array to simplify the triangle inequality check.
* Given $a \leq b \leq c$, only need to check $a + b > c$.
* Use a *two pointers* approach to improve efficiency compared to binary search.

---

### Approach

1. Sort the array `nums`.
2. Fix the largest side `c` by looping from right to left.
3. Use two pointers `l` and `r` to count valid pairs on the left of `c`.

   * If `nums[l] + nums[r] > nums[i]` → all pairs from `l` to `r-1` are valid → add `(r - l)`.
   * Otherwise, move `l` rightward.
4. Repeat until all triplets are checked.

---

### Complexity

* **Time:** $O(n^2)$, since for each anchor we scan linearly with two pointers.
* **Space:** $O(1)$, only a few extra variables are used.

---

### Example

**Input:**

```
nums = [2, 2, 3, 4]
```

**Output:**

```
3
```

**Explanation:** The valid triplets are `[2,3,4]` (twice, since there are two 2’s), and `[2,2,3]`.
