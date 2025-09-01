## Problem

**Indonesia:**  
Kamu memiliki `n` koin dan ingin membangun sebuah tangga menggunakan koin-koin tersebut. Tangga terdiri dari `k` baris, di mana baris ke-i memiliki tepat i koin. Baris terakhir dari tangga bisa saja tidak lengkap.

Diberikan integer `n`, kembalikan jumlah baris lengkap yang dapat dibangun dari tangga tersebut.

**English:**  
You have `n` coins and you want to build a staircase with these coins. The staircase consists of `k` rows where the ith row has exactly i coins. The last row of the staircase may be incomplete.

Given the integer `n`, return the number of complete rows of the staircase you will build.

---

## Solution

### Bahasa Indonesia

Solusi menggunakan **binary search**:
- Cari nilai maksimum `k` sehingga jumlah koin yang dibutuhkan untuk membangun k baris lengkap (`k*(k+1)/2`) tidak melebihi `n`.
- Lakukan pencarian biner pada rentang 1 hingga n.
- Jika jumlah koin yang dibutuhkan sama dengan n, kembalikan k.
- Jika jumlah koin lebih besar dari n, cari di bagian kiri.
- Jika jumlah koin lebih kecil dari n, cari di bagian kanan.
- Hasil akhir adalah nilai `right` setelah pencarian selesai.

### English

The solution uses **binary search**:
- Find the maximum `k` such that the total coins needed to build k complete rows (`k*(k+1)/2`) does not exceed `n`.
- Perform binary search in the range 1 to n.
- If the required coins equal n, return k.
- If the required coins are greater than n, search the left half.
- If the required coins are less than n, search the right half.
- The final answer is the value of `right` after the search ends.

**Time Complexity:** O(log n)  
**Space Complexity:** O(1)