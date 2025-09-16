# 🧮 Replace Non-Coprimes

## 🇮🇩 Bahasa Indonesia

### 📌 Deskripsi Masalah

Diberikan array integer `nums`. Kita perlu **menggabungkan dua bilangan bersebelahan yang non-coprime (GCD > 1)** menjadi **LCM-nya**, lalu ulangi proses ini sampai tidak ada lagi pasangan non-coprime yang bersebelahan.
Hasil akhirnya adalah array yang sudah dimodifikasi.

---

### 💡 Insight

* Dua bilangan **non-coprime** jika `gcd(a, b) > 1`.
* Jika ditemukan, maka diganti dengan `lcm(a, b) = a * b / gcd(a, b)`.
* Bisa terjadi **chaining**: setelah terbentuk angka baru, bisa jadi non-coprime lagi dengan elemen sebelumnya.
* Karena itu, solusi efektif adalah menggunakan struktur **stack**.

---

### 🛠️ Algoritma

1. Buat array `res` untuk menyimpan hasil sementara (anggap sebagai stack).
2. Masukkan elemen pertama `nums` ke `res`.
3. Iterasi elemen berikutnya di `nums`:

   * Jika gcd(elemen, res terakhir) == 1 → push ke `res`.
   * Jika gcd > 1 → ganti res terakhir dengan lcm(elemen, res terakhir), lalu cek mundur apakah perlu merge lagi.
4. Kembalikan `res` sebagai hasil akhir.

---

### ⏱️ Kompleksitas

* **Time:** O(n log M) (karena setiap elemen masuk-keluar stack sekali, plus gcd).
* **Space:** O(n) untuk stack.

---

### 📝 Contoh

```
nums = [6, 4, 3, 2, 7, 6, 2]

Step:
res = [6]
6 & 4 → gcd=2 → lcm=12 → res = [12]
12 & 3 → gcd=3 → lcm=12 → res = [12]
12 & 2 → gcd=2 → lcm=12 → res = [12]
12 & 7 → gcd=1 → res = [12, 7]
7 & 6 → gcd=1 → res = [12, 7, 6]
6 & 2 → gcd=2 → lcm=6 → res = [12, 7, 6]

Final: [12, 7, 6]
```

---

## 🇬🇧 English Version

### 📌 Problem Description

You are given an integer array `nums`.
We need to **merge two adjacent numbers that are non-coprime (GCD > 1)** into their **LCM**, and repeat the process until no such adjacent pairs remain.
Return the final modified array.

---

### 💡 Insight

* Two numbers are **non-coprime** if `gcd(a, b) > 1`.
* Replace them with `lcm(a, b) = a * b / gcd(a, b)`.
* This may cause **chaining merges**, so we need to keep checking backwards.
* A **stack-based approach** efficiently handles this.

---

### 🛠️ Algorithm

1. Initialize a result array `res` (used as a stack).
2. Push the first element of `nums` into `res`.
3. Iterate through the rest of `nums`:

   * If gcd(current, last of `res`) == 1 → push to `res`.
   * If gcd > 1 → replace last element of `res` with lcm, then keep merging backwards while possible.
4. Return `res` as the final result.

---

### ⏱️ Complexity

* **Time:** O(n log M), where `M` is the maximum element value.
* **Space:** O(n) for the stack.

---

### 📝 Example

```
nums = [6, 4, 3, 2, 7, 6, 2]

Step:
res = [6]
6 & 4 → gcd=2 → lcm=12 → res = [12]
12 & 3 → gcd=3 → lcm=12 → res = [12]
12 & 2 → gcd=2 → lcm=12 → res = [12]
12 & 7 → gcd=1 → res = [12, 7]
7 & 6 → gcd=1 → res = [12, 7, 6]
6 & 2 → gcd=2 → lcm=6 → res = [12, 7, 6]

Final: [12, 7, 6]
```