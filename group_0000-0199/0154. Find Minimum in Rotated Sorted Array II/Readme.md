## Problem

**Indonesia:**  
Diberikan sebuah array `nums` yang terurut menaik dan telah dirotasi sebanyak 1 hingga n kali.  
Array dapat berisi elemen duplikat.  
Contoh:  
- `[0,1,4,4,5,6,7]` bisa menjadi `[4,5,6,7,0,1,4]` jika dirotasi 4 kali.
- `[0,1,4,4,5,6,7]` jika dirotasi 7 kali tetap sama.

Rotasi satu kali berarti elemen terakhir pindah ke depan.

Tugas: Kembalikan elemen minimum dari array tersebut.

Algoritma harus mengurangi jumlah operasi sebanyak mungkin.

**English:**  
Suppose an array of length n sorted in ascending order is rotated between 1 and n times.  
The array `nums` may contain duplicates.  
For example:  
- `[0,1,4,4,5,6,7]` might become `[4,5,6,7,0,1,4]` if rotated 4 times.
- `[0,1,4,4,5,6,7]` if rotated 7 times remains the same.

Rotating once means the last element moves to the front.

Task: Return the minimum element of this array.

You must decrease the overall operation steps as much as possible.

---

## Solution

### Bahasa Indonesia

Ide solusi:
1. Lakukan binary search seperti biasa.
2. Jika menemukan kasus `nums[left] == nums[right]`, perkecil ruang pencarian dengan `left++` dan `right--`. Jangan lupa simpan nilai minimum dari `m` dan `nums[left]`.
3. Jika `nums[left] <= nums[mid]`, simpan nilai minimum dari `m` dan `nums[left]`, lalu geser `left` ke `mid + 1`.
4. Jika tidak, simpan nilai minimum dari `m` dan `nums[mid]`, lalu geser `right` ke `mid`.

Pendekatan ini memastikan pencarian minimum tetap efisien meskipun ada duplikat.

### English

Solution idea:
1. Use standard binary search.
2. If `nums[left] == nums[right]`, shrink the search space by `left++` and `right--`. Also, update the minimum value between `m` and `nums[left]`.
3. If `nums[left] <= nums[mid]`, update the minimum value between `m` and `nums[left]`, then move `left` to `mid + 1`.
4. Otherwise, update the minimum value between `m` and `nums[mid]`, then move `right` to `mid`.

This approach keeps the search for the minimum efficient even with duplicates.

**Time Complexity:** O(n) in worst case (when many duplicates), O(log n) in average  
**Auxiliary Space Complexity** O(1)