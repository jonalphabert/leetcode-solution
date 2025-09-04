## Problem

**Indonesia:**  
Diberikan sebuah string palindrom `s`.

Kembalikan permutasi palindromik dari `s` yang paling kecil secara leksikografis.

**English:**  
You are given a palindromic string `s`.

Return the lexicographically smallest palindromic permutation of `s`.

---

## Solution

### Bahasa Indonesia

- Hitung frekuensi setiap karakter dalam string.
- Untuk membentuk palindrom, bagi setiap karakter menjadi dua bagian yang sama besar untuk sisi kiri dan kanan.
- Susun sisi kiri dari karakter terkecil ke terbesar.
- Jika panjang string ganjil, letakkan karakter tengah sesuai posisi tengah.
- Susun sisi kanan dari karakter terbesar ke terkecil.
- Gabungkan sisi kiri, karakter tengah (jika ada), dan sisi kanan untuk mendapatkan palindrom paling kecil secara leksikografis.

### English

- Count the frequency of each character in the string.
- To form a palindrome, divide each character into two equal parts for the left and right sides.
- Arrange the left side from the smallest to the largest character.
- If the string length is odd, place the middle character at the center position.
- Arrange the right side from the largest to the smallest character.
- Concatenate the left side, middle character (if any), and right side to get the lexicographically smallest palindromic permutation.

**Time Complexity:** O(n + k), dengan k adalah jumlah karakter
**Space Complexity:** O(n)