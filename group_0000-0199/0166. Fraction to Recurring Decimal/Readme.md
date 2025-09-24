# 📖 Fraction to Recurring Decimal

## 🇮🇩 Bahasa Indonesia

### Deskripsi
Program ini mengimplementasikan fungsi `fractionToDecimal` yang menerima dua bilangan bulat `numerator` (pembilang) dan `denominator` (penyebut), lalu mengembalikan hasil bagi dalam bentuk string.

Aturannya:
1. Jika hasil bagi berupa bilangan bulat, kembalikan langsung sebagai string.
2. Jika hasil bagi memiliki bagian desimal:
   - Jika desimal **berhenti (finite)**, tampilkan apa adanya.
   - Jika desimal **berulang (repeating)**, tandai bagian yang berulang dengan kurung `()`.



### Cara Kerja
1. Tangani tanda negatif bila pembilang dan penyebut berbeda tanda.
2. Hitung bagian **integer** dengan `numerator / denominator`.
3. Simpan **sisa (remainder)** dengan `numerator % denominator`.
4. Jika `remainder == 0`, hasilnya bilangan bulat → selesai.
5. Jika tidak, proses bagian desimal:
   - Simpan posisi setiap `remainder` dalam sebuah map/dictionary.
   - Kalikan `remainder` dengan 10, ambil digit berikutnya dengan `remainder / denominator`, lalu update `remainder %= denominator`.
   - Jika `remainder` pernah muncul sebelumnya, berarti pola desimal mulai berulang → tambahkan kurung pada posisi tersebut.

## 🇬🇧 English

### Description

This program implements the function `fractionToDecimal` which takes two integers `numerator` and `denominator`, and returns the fraction result as a string.

Rules:

1. If the result is an integer, return it directly as a string.
2. If the result has a decimal part:

   * If the decimal **terminates (finite)**, display as is.
   * If the decimal **repeats**, enclose the repeating part in parentheses `()`.


### How It Works

1. Handle the sign if numerator and denominator have different signs.
2. Compute the **integer part** with `numerator / denominator`.
3. Store the **remainder** using `numerator % denominator`.
4. If `remainder == 0`, the result is an integer → done.
5. Otherwise, process the decimal part:

   * Store each remainder’s position in a map/dictionary.
   * Multiply remainder by 10, take the next digit using `remainder / denominator`, then update `remainder %= denominator`.
   * If a remainder repeats, the decimal cycle starts there → wrap the repeating part with parentheses.

**Time Complexity** : O(n)

**Space Complexity**: O(n)