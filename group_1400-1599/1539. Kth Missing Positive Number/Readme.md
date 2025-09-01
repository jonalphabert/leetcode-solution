## Problem

**Indonesia:**  
Diberikan sebuah array `arr` berisi bilangan bulat positif yang terurut secara menaik dan sebuah bilangan bulat `k`.  
Kembalikan bilangan bulat positif ke-k yang hilang dari array tersebut.

**English:**  
Given an array `arr` of positive integers sorted in strictly increasing order, and an integer `k`.  
Return the kth positive integer that is missing from this array.

---

## Solution

### Bahasa Indonesia

Ide solusi:
- Lakukan binary search dari indeks 0 hingga panjang array.
- Untuk setiap indeks `mid`, hitung berapa banyak angka yang hilang sampai indeks tersebut: `arr[mid] - (mid + 1)`.
- Jika jumlah angka yang hilang lebih kecil dari `k`, pindahkan `left` ke `mid + 1`.
- Jika lebih besar atau sama, pindahkan `right` ke `mid - 1`.
- Jika hasil akhirnya `right == -1`, maka bilangan yang hilang adalah `k` itu sendiri.
- Jika tidak, maka bilangan yang hilang adalah `k + right + 1`.

### English

Solution idea:
- Perform binary search from index 0 to the size of the array.
- For each index `mid`, count how many numbers are missing up to that index: `arr[mid] - (mid + 1)`.
- If the missing count is less than `k`, move `left` to `mid + 1`.
- If greater or equal, move `right` to `mid - 1`.
- If the result is `right == -1`, the missing number is `k` itself.
- Otherwise, the missing number is `k + right + 1`.

**Time Complexity:** O(log n)
**Space Complexity:** O(1)