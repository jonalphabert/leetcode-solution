## Problem

**Indonesia:**  
Kita bermain Guess Game. Permainannya sebagai berikut:

Saya memilih sebuah angka dari 1 sampai n. Kamu harus menebak angka yang saya pilih (angka yang dipilih tetap selama permainan).

Setiap kali tebakanmu salah, saya akan memberi tahu apakah angka yang saya pilih lebih tinggi atau lebih rendah dari tebakanmu.

Kamu dapat memanggil API yang sudah disediakan: `int guess(int num)`, yang mengembalikan:
- `-1`: Tebakanmu lebih tinggi dari angka yang saya pilih (`num > pick`)
- `1`: Tebakanmu lebih rendah dari angka yang saya pilih (`num < pick`)
- `0`: Tebakanmu sama dengan angka yang saya pilih (`num == pick`)

Kembalikan angka yang saya pilih.

**English:**  
We are playing the Guess Game. The game is as follows:

I pick a number from 1 to n. You have to guess which number I picked (the number I picked stays the same throughout the game).

Every time you guess wrong, I will tell you whether the number I picked is higher or lower than your guess.

You call a pre-defined API `int guess(int num)`, which returns three possible results:
- `-1`: Your guess is higher than the number I picked (`num > pick`)
- `1`: Your guess is lower than the number I picked (`num < pick`)
- `0`: Your guess is equal to the number I picked (`num == pick`)

Return the number that I picked.

---

## Solution

### Bahasa Indonesia

Solusi menggunakan **binary search**:
- Inisialisasi batas kiri (`left = 1`) dan kanan (`right = n`).
- Selama belum menemukan angka yang benar, lakukan:
  - Tebak angka tengah (`mid`).
  - Jika hasil API `guess(mid)` adalah -1, berarti angka yang dipilih lebih kecil, geser batas kanan ke `mid`.
  - Jika hasilnya 1, berarti angka yang dipilih lebih besar, geser batas kiri ke `mid + 1`.
  - Jika hasilnya 0, berarti angka ditemukan.
- Kembalikan angka yang ditemukan.

### English

The solution uses **binary search**:
- Initialize left boundary (`left = 1`) and right boundary (`right = n`).
- While the correct number is not found:
  - Guess the middle number (`mid`).
  - If the API `guess(mid)` returns -1, the picked number is smaller, move the right boundary to `mid`.
  - If it returns 1, the picked number is larger, move the left boundary to `mid + 1`.
  - If it returns 0, the number is found.
- Return the found number.

**Time Complexity:** O(log n)  
**Space Complexity:** O(1)