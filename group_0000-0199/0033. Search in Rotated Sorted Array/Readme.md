## Problem

**Indonesia:**  
Diberikan sebuah array integer `nums` yang terurut menaik (tanpa duplikasi), namun mungkin telah dirotasi ke kiri pada indeks yang tidak diketahui.  
Contoh: `[0,1,2,4,5,6,7]` bisa menjadi `[4,5,6,7,0,1,2]` jika dirotasi sebanyak 3 indeks.

Diberikan juga sebuah integer `target`.  
Kembalikan indeks dari `target` jika ditemukan di `nums`, atau `-1` jika tidak ditemukan.

Algoritma harus berjalan dalam waktu O(log n).

**English:**  
You are given an integer array `nums` sorted in ascending order (with distinct values), but possibly left rotated at an unknown index.  
For example, `[0,1,2,4,5,6,7]` might become `[4,5,6,7,0,1,2]` after a rotation by 3 indices.

Given the array `nums` after possible rotation and an integer `target`, return the index of `target` if it is in `nums`, or `-1` if it is not.

You must write an algorithm with O(log n) runtime complexity.

---

## Solution

### Bahasa Indonesia

Ide penyelesaian:
1. Lakukan binary search untuk mencari nilai terkecil dari array (pivot rotasi).
2. Index terkecil ini akan menjadi penentu partisi:
   - Jika `left` tidak berubah, berarti seluruh array sudah terurut, tidak ada rotasi.
   - Jika `target` lebih besar dari `nums[0]`, maka target pasti ada di antara indeks `0` sampai `left-1` (partisi besar).
   - Jika `target` lebih kecil dari `nums[0]`, maka target pasti ada di antara indeks `left` sampai `size-1` (partisi kecil).
3. Lakukan binary search pada partisi yang sudah ditentukan untuk mencari target.

### English

Solution idea:
1. Use binary search to find the smallest value in the array (rotation pivot).
2. This smallest index determines the partition:
   - If `left` does not change, the array is fully sorted, no rotation.
   - If `target` is greater than `nums[0]`, the target must be between indices `0` and `left-1` (large partition).
   - If `target` is less than `nums[0]`, the target must be between indices `left` and `size-1` (small partition).
3. Perform binary search on the determined partition to find the target.

**Time Complexity:** O(log n)  
**Auxiliary Space Complexity:** O(1)