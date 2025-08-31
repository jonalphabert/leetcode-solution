## Problem

**Indonesia:**  
Diberikan sebuah bilangan bulat `n`, temukan digit yang paling jarang muncul dalam representasi desimalnya.  
Jika ada beberapa digit dengan frekuensi yang sama, pilih digit yang paling kecil.

Kembalikan digit yang dipilih sebagai integer.

Frekuensi digit x adalah jumlah kemunculan x dalam representasi desimal dari n.

**English:**  
Given an integer `n`, find the digit that occurs least frequently in its decimal representation.  
If multiple digits have the same frequency, choose the smallest digit.

Return the chosen digit as an integer.

The frequency of a digit x is the number of times it appears in the decimal representation of n.

---

## Solution

### Bahasa Indonesia

- Hitung frekuensi setiap digit (0-9) dalam bilangan n.
- Cari digit yang memiliki frekuensi paling kecil (lebih dari 0).
- Jika ada beberapa digit dengan frekuensi sama, pilih digit terkecil.

### English

- Count the frequency of each digit (0-9) in the number n.
- Find the digit with the smallest frequency (greater than 0).
- If multiple digits have the same frequency, choose the smallest digit.

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Memori:** O(1)