## Problem

**Indonesia:**  
Angka Romawi direpresentasikan oleh tujuh simbol: I, V, X, L, C, D, dan M.

| Simbol | Nilai  |
|--------|--------|
| I      | 1      |
| V      | 5      |
| X      | 10     |
| L      | 50     |
| C      | 100    |
| D      | 500    |
| M      | 1000   |

Contoh:
- 2 ditulis sebagai II (1 + 1)
- 12 ditulis sebagai XII (10 + 1 + 1)
- 27 ditulis sebagai XXVII (10 + 10 + 5 + 1 + 1)

Biasanya angka Romawi ditulis dari yang terbesar ke yang terkecil dari kiri ke kanan. Namun, untuk angka seperti 4 (IV) dan 9 (IX), digunakan aturan pengurangan: jika angka kecil di depan angka besar, maka dikurangkan.

Aturan pengurangan:
- I sebelum V (5) dan X (10) untuk 4 dan 9
- X sebelum L (50) dan C (100) untuk 40 dan 90
- C sebelum D (500) dan M (1000) untuk 400 dan 900

Diberikan sebuah string angka Romawi, konversikan menjadi bilangan bulat.

**English:**  
Roman numerals are represented by seven different symbols: I, V, X, L, C, D, and M.

| Symbol | Value  |
|--------|--------|
| I      | 1      |
| V      | 5      |
| X      | 10     |
| L      | 50     |
| C      | 100    |
| D      | 500    |
| M      | 1000   |

Examples:
- 2 is written as II (1 + 1)
- 12 is written as XII (10 + 1 + 1)
- 27 is written as XXVII (10 + 10 + 5 + 1 + 1)

Roman numerals are usually written largest to smallest from left to right. However, for numbers like 4 (IV) and 9 (IX), subtraction is used: if a smaller number appears before a larger one, it is subtracted.

Subtraction rules:
- I before V (5) and X (10) for 4 and 9
- X before L (50) and C (100) for 40 and 90
- C before D (500) and M (1000) for 400 and 900

Given a Roman numeral string, convert it to an integer.

---

## Solution

### Bahasa Indonesia

- Buat map dari simbol Romawi ke nilai integer.
- Iterasi setiap karakter dari kiri ke kanan.
- Jika nilai karakter sebelumnya lebih kecil dari karakter sekarang, lakukan pengurangan (aturan pengurangan).
- Jika tidak, tambahkan nilai karakter ke hasil.

### English

- Create a map from Roman symbols to integer values.
- Iterate through each character from left to right.
- If the previous character's value is less than the current character's value, subtract (subtraction rule).
- Otherwise, add the value to the result.

**Time Complexity:** O(n)