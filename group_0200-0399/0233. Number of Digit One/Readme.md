# 0233. Number of Digit One (Greedy Math Solution and Digit DP Solution)

## 🇮🇩 Bahasa Indonesia

### 📌 Deskripsi Masalah

Diberikan sebuah bilangan bulat non-negatif `n`.
Hitung jumlah kemunculan digit **1** di semua bilangan dari `0` hingga `n` (inklusif).

Contoh:

* Input: `n = 13`
* Output: `6`
  Penjelasan: digit `1` muncul pada `1, 10, 11, 12, 13` total **6 kali**.

---

### 🔍 Abstraksi & Ide

Ada dua pendekatan populer:

#### 1. Matematika (Greedy / Analitik)

Untuk setiap posisi digit (satuan, puluhan, ratusan, dst.), kita hitung berapa banyak `1` muncul pada posisi tersebut.
Misalnya untuk angka `n = 1234`:

* Posisi satuan: hitung berapa banyak blok `10` yang menyumbang digit `1`.
* Posisi puluhan: hitung berapa banyak blok `100` yang menyumbang digit `1` di puluhan.
* Dst.

Formula umumnya:

```
count += (n / (base*10)) * base + min(max(n % (base*10) - base + 1, 0), base)
```

di mana `base = 1, 10, 100, ...`.

---

#### 2. Dynamic Programming (Digit DP)

Pendekatan DP digunakan dengan **digit by digit traversal**:

* State: `(pos, tight, countOne)`

  * `pos`: indeks digit yang sedang diproses
  * `tight`: apakah prefix masih sama dengan `n` (agar tidak melebihi)
  * `countOne`: jumlah `1` yang muncul sejauh ini
* Transition: pilih digit `d` dari `0` hingga batas (`9` atau `digits[pos]` jika tight).
* Jika `d == 1`, maka tambahkan `countOne`.
* Rekurens berlanjut sampai digit terakhir.

---

### 💡 Penyelesaian

#### 1. Matematika

Iterasi setiap posisi digit dengan `base = 1, 10, 100, ...` dan gunakan formula untuk menjumlahkan kemunculan `1`.

#### 2. Digit DP

Gunakan DFS + memoization untuk menelusuri semua kemungkinan angka hingga `n`.
State `(pos, tight, countOne)` di-memo agar tidak dihitung ulang.

---

### ⏱️ Kompleksitas

#### Matematika

* **Waktu:** `O(log n)`
* **Ruang:** `O(1)`

#### Digit DP

* **Waktu:** `O(d * 10)` dengan `d = jumlah digit dari n`
* **Ruang:** `O(d * d * 2)` untuk memo (kecil karena d ≤ 10)

---

## 🇬🇧 English

### 📌 Problem Description

You are given a non-negative integer `n`.
Count the total number of times the digit **1** appears in all numbers from `0` to `n` (inclusive).

Example:

* Input: `n = 13`
* Output: `6`
  Explanation: digit `1` appears in `1, 10, 11, 12, 13` → total **6 times**.

---

### 🔍 Abstraction & Idea

Two common approaches:

#### 1. Mathematics (Greedy / Analytical)

For each digit position (ones, tens, hundreds, …), calculate how many times digit `1` appears at that position.
Formula:

```
count += (n / (base*10)) * base + min(max(n % (base*10) - base + 1, 0), base)
```

where `base = 1, 10, 100, …`.

---

#### 2. Dynamic Programming (Digit DP)

Use **digit-by-digit DP traversal**:

* State: `(pos, tight, countOne)`

  * `pos`: current digit index
  * `tight`: whether current prefix matches `n` (not allowed to exceed)
  * `countOne`: how many `1`s counted so far
* Transition: pick digit `d` in `[0, limit]`.
* If `d == 1`, increase `countOne`.
* Recurse until the last digit.

---

### 💡 Solution

#### 1. Mathematics

Iterate through each digit position using the formula and accumulate counts.

#### 2. Digit DP

Use DFS + memoization to explore all valid digit combinations up to `n`, counting how many times digit `1` occurs.

---

### ⏱️ Complexity

#### Mathematics

* **Time:** `O(log n)`
* **Space:** `O(1)`

#### Digit DP

* **Time:** `O(d * 10)`, where `d = number of digits in n`
* **Space:** `O(d * d * 2)` due to memoization (small since `d ≤ 10`)

---