## Problem

**Indonesia:**  
Diberikan sebuah string palindrom `s` dan sebuah bilangan bulat `k`.

Kembalikan permutasi palindromik ke-k yang paling kecil secara leksikografis dari `s`.  
Jika jumlah permutasi palindromik yang berbeda kurang dari `k`, kembalikan string kosong.

Catatan:  
- Rearrangement yang menghasilkan string palindrom yang sama dianggap identik dan hanya dihitung satu kali.

**English:**  
You are given a palindromic string `s` and an integer `k`.

Return the k-th lexicographically smallest palindromic permutation of `s`.  
If there are fewer than `k` distinct palindromic permutations, return an empty string.

Note:  
- Different rearrangements that yield the same palindromic string are considered identical and are counted once.

---

## Solution

### Bahasa Indonesia

- Hitung frekuensi setiap karakter dalam string.
- Untuk membentuk palindrom, bagi setiap karakter menjadi dua bagian untuk sisi kiri dan kanan.
- Untuk membangun palindrom ke-k secara leksikografis:
  - Pada setiap posisi, coba pilih karakter terkecil yang masih tersedia.
  - Setelah memilih satu karakter, hitung berapa banyak permutasi palindromik yang bisa dibentuk dengan sisa karakter (menggunakan rumus permutasi kombinasi).
  - Jika jumlah permutasi cukup untuk mencapai k, pilih karakter tersebut.
  - Jika tidak cukup, kurangi k dengan jumlah permutasi dan lanjutkan ke karakter berikutnya.
- Jika panjang string ganjil, letakkan karakter tengah sesuai posisi tengah.
- Gabungkan sisi kiri, karakter tengah (jika ada), dan sisi kanan (reverse dari sisi kiri) untuk membentuk palindrom.

### English

- Count the frequency of each character in the string.
- To form a palindrome, divide each character into two parts for the left and right sides.
- To construct the k-th lexicographically smallest palindromic permutation:
  - At each position, try to pick the smallest available character.
  - After picking a character, calculate how many palindromic permutations can be formed with the remaining characters (using permutation formulas).
  - If the number of permutations is enough to reach k, pick that character.
  - If not, subtract the number of permutations from k and try the next character.
- If the string length is odd, place the middle character at the center.
- Concatenate the left side, middle character (if any), and the reversed left side to form the palindrome.

**Perbedaan dengan versi sebelumnya:**  
- Di sini, kita perlu menghitung jumlah permutasi palindromik yang bisa dibentuk setiap kali memilih karakter, agar bisa menentukan karakter mana yang harus dipilih untuk mencapai urutan ke-k.

**Time Complexity:** O(n * 26)  
**Space Complexity:** O(n)