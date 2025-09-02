## Problem

**Indonesia:**  
Diberikan sebuah array integer `nums` dengan panjang n dan sebuah integer `target`.  
Temukan tiga bilangan dalam `nums` sehingga jumlahnya paling mendekati `target`.

Kembalikan jumlah dari tiga bilangan tersebut.

Diasumsikan setiap input memiliki tepat satu solusi.

**English:**  
Given an integer array `nums` of length n and an integer `target`, find three integers in `nums` such that the sum is closest to `target`.

Return the sum of the three integers.

You may assume that each input would have exactly one solution.

---

## Solution

### Bahasa Indonesia

- Urutkan array `nums` terlebih dahulu.
- Untuk setiap indeks `i`, gunakan dua pointer (`left` di `i+1` dan `right` di akhir array).
- Hitung jumlah tiga angka (`nums[i] + nums[left] + nums[right]`).
- Jika selisih dengan `target` lebih kecil dari minimum sebelumnya, simpan hasil dan selisih baru.
- Jika jumlah lebih besar dari `target`, geser `right` ke kiri. Jika lebih kecil atau sama, geser `left` ke kanan.
- Ulangi hingga semua kemungkinan diperiksa.

### English

- Sort the `nums` array first.
- For each index `i`, use two pointers (`left` at `i+1` and `right` at the end of the array).
- Calculate the sum of three numbers (`nums[i] + nums[left] + nums[right]`).
- If the difference with `target` is smaller than the previous minimum, update the result and minimum difference.
- If the sum is greater than `target`, move `right` left. If less than or equal, move `left` right.
- Repeat until all possibilities are checked.

**Time Complexity:** O(n²)  
**Space Complexity:** O(1)