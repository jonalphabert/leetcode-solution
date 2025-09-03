## Problem

**Indonesia:**  
Implementasikan fungsi `myAtoi(string s)` yang mengkonversi sebuah string menjadi bilangan bulat 32-bit bertanda.

Algoritma untuk `myAtoi(string s)` adalah sebagai berikut:
- **Whitespace:** Abaikan semua whitespace di awal string.
- **Tanda:** Tentukan tanda dengan memeriksa apakah karakter berikutnya adalah '-' atau '+'. Jika tidak ada, asumsikan positif.
- **Konversi:** Baca digit angka, lewati nol di depan, hingga menemukan karakter non-digit atau akhir string. Jika tidak ada digit yang terbaca, hasilnya adalah 0.
- **Pembulatan:** Jika hasil keluar dari rentang integer 32-bit bertanda `[-2^31, 2^31 - 1]`, bulatkan ke batas terdekat.
- Kembalikan hasil konversi sebagai integer.

**English:**  
Implement the `myAtoi(string s)` function, which converts a string to a 32-bit signed integer.

The algorithm for `myAtoi(string s)` is as follows:
- **Whitespace:** Ignore any leading whitespace (" ").
- **Signedness:** Determine the sign by checking if the next character is '-' or '+', assuming positivity if neither present.
- **Conversion:** Read the integer by skipping leading zeros until a non-digit character is encountered or the end of the string is reached. If no digits were read, then the result is 0.
- **Rounding:** If the integer is out of the 32-bit signed integer range `[-2^31, 2^31 - 1]`, then round the integer to remain in the range.
- Return the integer as the final result.

---

## Solution

### Bahasa Indonesia

- Trim whitespace di awal string.
- Tentukan tanda (positif/negatif) dari karakter pertama setelah whitespace.
- Iterasi setiap karakter digit, konversi ke angka, dan bangun hasil.
- Cek overflow sebelum menambah digit baru.
- Jika hasil keluar dari rentang integer, kembalikan batas terdekat.
- Kembalikan hasil akhir dengan tanda yang sesuai.

### English

- Trim leading whitespace from the string.
- Determine the sign (positive/negative) from the first non-whitespace character.
- Iterate through each digit character, convert to number, and build the result.
- Check for overflow before adding a new digit.
- If the result is out of integer range, return the nearest bound.
- Return the final result with the correct sign.

**Time Complexity:** O(len(s))  
**Space Complexity:** O(1)