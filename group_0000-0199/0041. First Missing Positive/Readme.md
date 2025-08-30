## Problem

**Indonesia:**  
Diberikan sebuah array integer `nums` yang tidak terurut.  
Kembalikan bilangan bulat positif terkecil yang **tidak ada** di dalam `nums`.

Algoritma harus berjalan dalam waktu O(n) dan menggunakan ruang tambahan O(1).

**English:**  
Given an unsorted integer array `nums`.  
Return the smallest positive integer that is **not present** in `nums`.

You must implement an algorithm that runs in O(n) time and uses O(1) auxiliary space.

---

## Solution

### Bahasa Indonesia

- Tempatkan setiap angka positif `x` (1 ≤ x ≤ n) ke indeks `x-1` dengan cara swap.
- Setelah penempatan, iterasi array dari kiri ke kanan.  
  Bilangan pertama di indeks `i` yang bukan `i+1` adalah jawaban.
- Jika semua posisi sudah benar, maka jawabannya adalah `n+1`.

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Memori:** O(1)

### English

- Place each positive number `x` (1 ≤ x ≤ n) at index `x-1` by swapping.
- After placement, iterate the array from left to right.  
  The first index `i` where `nums[i] != i+1` is the answer.
- If all positions are correct, the answer is `n+1`.

**Time Complexity:** O(n)  
**Space Complexity:** O(1)