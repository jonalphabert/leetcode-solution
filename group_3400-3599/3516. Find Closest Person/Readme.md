## Problem

**Indonesia:**  
Diberikan tiga bilangan bulat `x`, `y`, dan `z` yang merepresentasikan posisi tiga orang pada sebuah garis bilangan:
- `x` adalah posisi Orang 1.
- `y` adalah posisi Orang 2.
- `z` adalah posisi Orang 3 (Orang 3 tidak bergerak).

Orang 1 dan Orang 2 bergerak menuju Orang 3 dengan kecepatan yang sama.

Tentukan siapa yang lebih dulu sampai ke posisi Orang 3:
- Kembalikan `1` jika Orang 1 sampai lebih dulu.
- Kembalikan `2` jika Orang 2 sampai lebih dulu.
- Kembalikan `0` jika keduanya sampai bersamaan.

**English:**  
You are given three integers `x`, `y`, and `z`, representing the positions of three people on a number line:
- `x` is the position of Person 1.
- `y` is the position of Person 2.
- `z` is the position of Person 3 (Person 3 does not move).

Both Person 1 and Person 2 move toward Person 3 at the same speed.

Determine which person reaches Person 3 first:
- Return `1` if Person 1 arrives first.
- Return `2` if Person 2 arrives first.
- Return `0` if both arrive at the same time.

---

## Solution

### Bahasa Indonesia

- Hitung jarak absolut antara Orang 1 ke Orang 3 (`|z - x|`) dan Orang 2 ke Orang 3 (`|z - y|`).
- Jika jarak keduanya sama, kembalikan `0`.
- Jika jarak Orang 1 lebih kecil, kembalikan `1`.
- Jika jarak Orang 2 lebih kecil, kembalikan `2`.

### English

- Calculate the absolute distance from Person 1 to Person 3 (`|z - x|`) and from Person 2 to Person 3 (`|z - y|`).
- If both distances are equal, return `0`.
- If Person 1's distance is smaller, return `1`.
- If Person 2's distance is smaller, return `2`.

**Time Complexity:** O(1)
**Space Complexity:** O(1)