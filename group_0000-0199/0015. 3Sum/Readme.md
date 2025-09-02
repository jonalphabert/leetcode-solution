## Problem

**Indonesia:**  
Diberikan sebuah array integer `nums`.  
Kembalikan semua triplet `[nums[i], nums[j], nums[k]]` sehingga:
- i, j, k adalah indeks yang berbeda (i ≠ j, i ≠ k, j ≠ k)
- nums[i] + nums[j] + nums[k] == 0

Perhatikan bahwa solusi tidak boleh mengandung triplet yang duplikat.

**English:**  
Given an integer array `nums`, return all the triplets `[nums[i], nums[j], nums[k]]` such that:
- i, j, k are distinct indices (i ≠ j, i ≠ k, j ≠ k)
- nums[i] + nums[j] + nums[k] == 0

Notice that the solution set must not contain duplicate triplets.

---

## Solution

### Bahasa Indonesia

- Urutkan array `nums` terlebih dahulu.
- Untuk setiap indeks `j`, gunakan dua pointer (`left` di `j+1` dan `right` di akhir array).
- Hitung jumlah tiga angka (`nums[j] + nums[left] + nums[right]`).
- Jika hasilnya 0, simpan triplet ke hasil jika belum ada sebelumnya (gunakan set untuk menghindari duplikasi).
- Setelah menemukan triplet, geser pointer `left` dan `right` untuk melewati angka yang sama agar tidak duplikat.
- Jika jumlah lebih besar dari 0, geser `right` ke kiri. Jika lebih kecil, geser `left` ke kanan.

### English

- Sort the `nums` array first.
- For each index `j`, use two pointers (`left` at `j+1` and `right` at the end of the array).
- Calculate the sum of three numbers (`nums[j] + nums[left] + nums[right]`).
- If the result is 0, add the triplet to the result if it hasn't been added before (use a set to avoid duplicates).
- After finding a triplet, move `left` and `right` pointers to skip duplicate numbers.
- If the sum is greater than 0, move `right` left. If less, move `left` right.

**Time Complexity:** O(n²)
**Space Complexity:** O(n) for the result set