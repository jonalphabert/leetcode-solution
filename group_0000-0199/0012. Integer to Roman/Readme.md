## Problem

**Indonesia:**  
Terdapat tujuh simbol yang merepresentasikan angka Romawi dengan nilai berikut:

| Simbol | Nilai  |
|--------|--------|
| I      | 1      |
| V      | 5      |
| X      | 10     |
| L      | 50     |
| C      | 100    |
| D      | 500    |
| M      | 1000   |

Angka Romawi dibentuk dengan menambahkan konversi nilai desimal dari yang terbesar ke yang terkecil.  
Aturan konversi:
- Jika nilai tidak diawali dengan 4 atau 9, pilih simbol dengan nilai maksimal yang bisa dikurangkan dari input, tambahkan simbol ke hasil, kurangi nilainya, dan konversikan sisa nilai.
- Jika nilai diawali dengan 4 atau 9, gunakan bentuk pengurangan (subtractive form), misal 4 adalah IV, 9 adalah IX. Bentuk pengurangan yang digunakan: 4 (IV), 9 (IX), 40 (XL), 90 (XC), 400 (CD), dan 900 (CM).
- Hanya simbol pangkat 10 (I, X, C, M) yang boleh diulang maksimal 3 kali untuk merepresentasikan kelipatan 10. Simbol 5 (V), 50 (L), atau 500 (D) tidak boleh diulang. Jika perlu 4 kali, gunakan bentuk pengurangan.

Diberikan sebuah bilangan bulat, konversikan menjadi angka Romawi.

**English:**  
Seven different symbols represent Roman numerals with the following values:

| Symbol | Value  |
|--------|--------|
| I      | 1      |
| V      | 5      |
| X      | 10     |
| L      | 50     |
| C      | 100    |
| D      | 500    |
| M      | 1000   |

Roman numerals are formed by appending the conversions of decimal place values from highest to lowest.  
Conversion rules:
- If the value does not start with 4 or 9, select the symbol of the maximal value that can be subtracted from the input, append that symbol to the result, subtract its value, and convert the remainder.
- If the value starts with 4 or 9, use the subtractive form (one symbol subtracted from the next), e.g., 4 is IV, 9 is IX. Only the following subtractive forms are used: 4 (IV), 9 (IX), 40 (XL), 90 (XC), 400 (CD), and 900 (CM).
- Only powers of 10 (I, X, C, M) can be appended consecutively at most 3 times to represent multiples of 10. You cannot append 5 (V), 50 (L), or 500 (D) multiple times. If you need to append a symbol 4 times, use the subtractive form.

Given an integer, convert it to a Roman numeral.

---

## Solution

### Bahasa Indonesia

- Siapkan dua array: satu berisi nilai-nilai Romawi dari terbesar ke terkecil, satu lagi berisi simbol Romawi yang sesuai.
- Iterasi dari nilai terbesar ke terkecil, selama nilai masih bisa dikurangkan dari input, tambahkan simbol ke hasil dan kurangi nilainya.
- Ulangi hingga nilai menjadi 0.

### English

- Prepare two arrays: one with Roman numeral values from largest to smallest, and one with the corresponding Roman symbols.
- Iterate from the largest value to the smallest, while the value can be subtracted from the input, append the symbol to the result and subtract its value.
- Repeat until the value becomes 0.

**Time Complexity:** O(1)