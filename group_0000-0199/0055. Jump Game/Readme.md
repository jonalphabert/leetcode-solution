## Problem

**Indonesia:**  
Diberikan sebuah array integer `nums`.  
Awalnya kamu berada di indeks pertama array.  
Setiap elemen `nums[i]` menunjukkan panjang maksimum lompatan dari posisi tersebut.

Kembalikan `true` jika kamu bisa mencapai indeks terakhir, atau `false` jika tidak bisa.

**English:**  
You are given an integer array `nums`.  
You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return `true` if you can reach the last index, or `false` otherwise.

---

## Solution

### Bahasa Indonesia

Solusi menggunakan pendekatan **greedy**:
- Lacak indeks terjauh yang bisa dicapai (`farthest`).
- Iterasi dari indeks pertama, pada setiap langkah perbarui `farthest` dengan nilai maksimum dari `i + nums[i]`.
- Jika indeks saat ini mencapai batas lompatan (`end`), perbarui batas ke `farthest`.
- Jika ditemukan batas lompatan (`end`) lebih kecil dari sebuah index `i`, maka returnnya `false` karena tidak ada langkah yang valid jika index `i` lebih besar dari batas lompatan `end`
- Jika `farthest` sudah mencapai atau melewati indeks terakhir, return `true`.
- Jika tidak, return `false`.

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Memori:** O(1)

### English

The solution uses a **greedy** approach:
- Track the farthest index that can be reached (`farthest`).
- Iterate from the first index, updating `farthest` to the maximum of `i + nums[i]` at each step.
- If the current index reaches the end of the current jump (`end`), update `end` to `farthest`.
- If at any point `end` is less than the current index `i`, return `false` since there are no valid moves if index `i` is greater than the jump limit `end`.
- If `farthest` reaches or exceeds the last index, return `true`.
- Otherwise, return `false`.

**Time Complexity:** O(n)  
**Space Complexity:**

### Reference
Problem ini mirip dengan [0045. Jump Game II](../0045.%20Jump%20Game%20II/Readme.md)