## Problem

**Indonesia:**  
Diberikan dua bilangan bulat `num1` dan `num2`.

Dalam satu operasi, kamu dapat memilih integer `i` pada rentang [0, 60] dan mengurangi `2^i + num2` dari `num1`.

Kembalikan jumlah minimum operasi yang diperlukan agar `num1` menjadi 0.

Jika tidak mungkin membuat `num1` menjadi 0, kembalikan -1.

---

## Cara Berpikir dan Solusi

### 1. Bentuk Operasi

Setiap operasi mengurangi `num1` dengan nilai `num2` ditambah sebuah pangkat dua:  
`num1 -= (2^i + num2)`, dengan `i ∈ [0, 60]`.

### 2. Reformulasi Problem

Jika kita melakukan `k` operasi dengan memilih indeks `i1, i2, ..., ik`, maka total yang dikurangi dari `num1` adalah:  
`k * num2 + sum_{j=1}^{k} 2^{ij}`

Kita ingin `num1` menjadi 0, jadi:  
`num1 = k * num2 + sum_{j=1}^{k} 2^{ij}`

### 3. Observasi Penting

Bagian `sum_{j=1}^{k} 2^{ij}` adalah jumlah dari beberapa pangkat dua, sehingga selalu bilangan bulat ≥ 0.

Artinya, kita butuh:  
`num1 - k * num2 ≥ 0`  
dan sisi kanan harus bisa direpresentasikan sebagai jumlah pangkat dua.

Jumlah pangkat dua itu bisa direpresentasikan jika jumlahnya ≥ jumlah bit 1 pada representasi binernya (popcount).

Jadi, untuk sebuah angka `X = num1 - k * num2`,  
- Minimal jumlah pangkat dua untuk membentuk `X` adalah `popcount(X)`.
- Maksimal jumlah pangkat dua adalah `X` (jika semuanya pakai 2^0).

Kondisi yang harus dipenuhi:  
`popcount(num1 - k * num2) ≤ k ≤ (num1 - k * num2)`

### 4. Strategi Solusi

- Loop `k` dari 1 sampai batas tertentu.
- Hitung `rem = num1 - k * num2`.
- Jika `rem < 0`, stop (karena semakin besar k, nilai semakin kecil).
- Cek apakah `popcount(rem) ≤ k ≤ rem`.
  - Jika ya, return k.
  - Jika tidak, lanjutkan.
- Jika tidak ada yang cocok, return -1.

### 5. Kompleksitas

- Loop `k` maksimal sekitar `10^9 / |num2|`, tapi biasanya kecil karena constraint.
- Praktis cukup cepat karena setiap iterasi hanya cek popcount.

---

## English

**Problem:**  
You are given two integers `num1` and `num2`.

In one operation, you can choose integer `i` in the range [0, 60] and subtract `2^i + num2` from `num1`.

Return the minimum number of operations needed to make `num1` equal to 0.

If it is impossible to make `num1` equal to 0, return -1.

---

**Solution Approach:**

1. **Operation Formulation:**  
   Each operation subtracts `num2` plus a power of two from `num1`:  
   `num1 -= (2^i + num2)`, with `i ∈ [0, 60]`.

2. **Problem Reformulation:**  
   If we perform `k` operations with indices `i1, i2, ..., ik`, the total subtracted from `num1` is:  
   `k * num2 + sum_{j=1}^{k} 2^{ij}`  
   We want `num1 = k * num2 + sum_{j=1}^{k} 2^{ij}`.

3. **Key Observation:**  
   The sum of powers of two is always an integer ≥ 0.  
   So, we need:  
   `num1 - k * num2 ≥ 0`  
   and the remainder must be representable as a sum of powers of two.

   For a number `X = num1 - k * num2`:
   - The minimum number of powers of two needed to represent `X` is `popcount(X)`.
   - The maximum is `X` (if all are 2^0).

   So, the condition is:  
   `popcount(num1 - k * num2) ≤ k ≤ (num1 - k * num2)`

4. **Solution Strategy:**  
   - Loop `k` from 1 up to a reasonable bound.
   - Compute `rem = num1 - k * num2`.
   - If `rem < 0`, stop.
   - Check if `popcount(rem) ≤ k ≤ rem`.
     - If yes, return k.
     - If not, continue.
   - If no valid k is found, return -1.

5. **Complexity:**  
   - The loop for k is at most about `10^9 / |num2|`, but usually much smaller due to constraints.
   - Each iteration is fast (just popcount check).