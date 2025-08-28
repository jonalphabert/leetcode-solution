# Indonesia
## Problem

Diberikan sebuah bilangan bulat bertanda 32-bit `x`, kembalikan `x` dengan digit-digitnya dibalik.  
Jika pembalikan digit menyebabkan nilai keluar dari rentang bilangan bulat bertanda 32-bit `[-2^31, 2^31 - 1]`, maka kembalikan 0.

Asumsikan lingkungan tidak mengizinkan penyimpanan bilangan bulat 64-bit (signed maupun unsigned).

**Contoh:**
```
Input: x = 123
Output: 321
```
```
Input: x = -123
Output: -321
```
```
Input: x = 120
Output: 21
```
```
Input: x = 1534236469
Output: 0
Penjelasan: Hasil pembalikan melebihi batas integer 32-bit.
```

**Batasan:**
- `-2^31 <= x <= 2^31 - 1`

## Solution

Untuk membalik digit bilangan bulat, lakukan iterasi selama `x` tidak sama dengan 0:
1. Ambil digit terakhir dengan operasi modulus 10.
2. Tambahkan digit tersebut ke hasil yang sudah dikalikan 10.
3. Periksa apakah hasil pembalikan melebihi batas integer 32-bit sebelum menambah digit berikutnya.
4. Jika melebihi batas, kembalikan 0.

**Kompleksitas Waktu:** O(log₁₀ n)  
**Kompleksitas Memori:** O(1)



# English
## Problem

Given a signed 32-bit integer `x`, return `x` with its digits reversed.  
If reversing `x` causes the value to go outside the signed 32-bit integer range `[-2^31, 2^31 - 1]`, then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

**Example:**
```
Input: x = 123
Output: 321
```
```
Input: x = -123
Output: -321
```
```
Input: x = 120
Output: 21
```
```
Input: x = 1534236469
Output: 0
Explanation: The reversed result exceeds the 32-bit integer limit.
```

**Constraints:**
- `-2^31 <= x <= 2^31 - 1`

## Solution

To reverse the digits of an integer, iterate while `x` is not zero:
1. Take the last digit using modulus 10.
2. Add the digit to the result after multiplying the result by 10.
3. Check if the reversed result exceeds the 32-bit integer limit before adding the next digit.
4. If it exceeds the limit, return 0.

**Time Complexity:** O(log₁₀ n)  
**Space Complexity:** O(1)