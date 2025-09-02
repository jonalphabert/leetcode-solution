## Problem

**Indonesia:**  
Diberikan sebuah string `s` dan sebuah integer `numRows`.  
Tuliskan string tersebut dalam pola zigzag sebanyak `numRows` baris, lalu baca hasilnya baris per baris.

Contoh:
```
Input: s = "PAYPALISHIRING", numRows = 3

P   A   H   N
A P L S I I G
Y   I   R

Output: "PAHNAPLSIIGYIR"
```

**English:**  
Given a string `s` and an integer `numRows`, write the string in a zigzag pattern on `numRows` rows, then read the result line by line.

Example:
```
Input: s = "PAYPALISHIRING", numRows = 3

P   A   H   N
A P L S I I G
Y   I   R

Output: "PAHNAPLSIIGYIR"
```

---

## Solution

### 1. Pattern Observation

**Indonesia:**  
Solusi ini mengamati pola lompatan karakter pada setiap baris.  
- Untuk setiap baris, lakukan iterasi dan tambahkan karakter sesuai pola lompatan zigzag.
- Baris pertama dan terakhir memiliki pola lompatan tetap, sedangkan baris tengah memiliki dua pola lompatan yang bergantian.

**English:**  
This solution observes the jump pattern for each row.  
- For each row, iterate and add characters according to the zigzag jump pattern.
- The first and last rows have a fixed jump pattern, while the middle rows alternate between two jump patterns.

---

### 2. Zigzag Simulation

**Indonesia:**  
Solusi ini mensimulasikan penulisan karakter ke baris zigzag:
- Buat array string sebanyak `numRows`.
- Iterasi setiap karakter, tambahkan ke baris saat ini.
- Ubah arah (naik/turun) jika mencapai baris pertama atau terakhir.
- Gabungkan semua baris untuk mendapatkan hasil akhir.

**English:**  
This solution simulates writing characters to zigzag rows:
- Create an array of strings for each row.
- Iterate through each character, adding it to the current row.
- Change direction (down/up) when reaching the first or last row.
- Concatenate all rows to get the final result.

---

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Memori:** O(n)