# 0279. Perfect Squares (DP Solution)

## 🇮🇩 Bahasa Indonesia

### 📌 Deskripsi Masalah

Diberikan sebuah bilangan bulat `n`.
Temukan jumlah paling sedikit dari bilangan **kuadrat sempurna** yang jika dijumlahkan akan sama dengan `n`.

Contoh:

* Input: `n = 12`
* Output: `3`
  Penjelasan: `12 = 4 + 4 + 4`.
* Input: `n = 13`
* Output: `2`
  Penjelasan: `13 = 4 + 9`.

---

### 🔍 Abstraksi & Ide

* Bilangan kuadrat sempurna: `1, 4, 9, 16, …`.
* Masalah ini mirip dengan **coin change problem (minimum coin)**.
* Untuk setiap `i`, kita bisa mencoba kurangi dengan `j*j` (kuadrat ≤ i), lalu ambil minimum hasilnya.

Definisi DP:

```
dp[i] = minimum jumlah kuadrat untuk membentuk i
```

Transisi:

```
dp[i] = min(dp[i - j*j] + 1)   untuk setiap j dengan j*j ≤ i
```

---

### 💡 Penyelesaian

1. Inisialisasi array `dp` berukuran `n+1` dengan nilai besar (infinity).
2. Set `dp[0] = 0`.
3. Iterasi `i` dari 1 sampai `n`:

   * Untuk setiap kuadrat `j*j ≤ i`, update `dp[i] = min(dp[i], dp[i - j*j] + 1)`.
4. Jawaban akhir ada di `dp[n]`.

---

### ⏱️ Kompleksitas

* **Waktu:** `O(n * sqrt(n))`
  (karena untuk setiap `i`, kita cek hingga `√i` kemungkinan kuadrat).
* **Ruang:** `O(n)`
  (untuk menyimpan array `dp`).

---

## 🇬🇧 English

### 📌 Problem Description

Given an integer `n`, return the least number of **perfect square numbers** that sum to `n`.

Example:

* Input: `n = 12`
* Output: `3`
  Explanation: `12 = 4 + 4 + 4`.
* Input: `n = 13`
* Output: `2`
  Explanation: `13 = 4 + 9`.

---

### 🔍 Abstraction & Idea

* Perfect squares are numbers like `1, 4, 9, 16, …`.
* This problem is similar to the **minimum coin change problem**.
* For each `i`, we try subtracting all possible `j*j` and take the minimum result.

DP definition:

```
dp[i] = minimum number of perfect squares to form i
```

Transition:

```
dp[i] = min(dp[i - j*j] + 1)   for all j such that j*j ≤ i
```

---

### 💡 Solution

1. Initialize `dp` array of size `n+1` with infinity.
2. Set `dp[0] = 0`.
3. Iterate `i` from 1 to `n`:

   * For each square `j*j ≤ i`, update `dp[i] = min(dp[i], dp[i - j*j] + 1)`.
4. The final answer is `dp[n]`.

---

### ⏱️ Complexity

* **Time:** `O(n * sqrt(n))`
  (since for each `i`, we check up to `√i` squares).
* **Space:** `O(n)`
  (for storing the DP array).