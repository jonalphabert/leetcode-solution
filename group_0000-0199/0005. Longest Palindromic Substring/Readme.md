# Indonesia 
## Problem

Diberikan sebuah string `s`, kembalikan substring palindrom terpanjang dalam `s`.  
Palindrom adalah string yang dibaca sama baik dari depan maupun belakang.

**Contoh:**
```
Input: s = "babad"
Output: "bab"
Penjelasan: "aba" juga merupakan jawaban yang valid.
```

```
Input: s = "cbbd"
Output: "bb"
```

**Batasan:**
- `1 <= s.length <= 1000`
- `s` hanya terdiri dari digit dan huruf Inggris.

## Solution

Untuk menyelesaikan masalah ini, kita dapat menggunakan pendekatan **Expand Around Center**:

1. Iterasi setiap karakter dalam string, anggap setiap karakter (dan celah antar karakter) sebagai pusat kemungkinan palindrom.
2. Perluas ke kiri dan kanan dari pusat selama karakter di kedua sisi sama.
3. Simpan substring palindrom terpanjang yang ditemukan selama proses.

**Kompleksitas Waktu:** O(n²)  
**Kompleksitas

# English
## Problem

Given a string `s`, return the longest palindromic substring in `s`.  
A palindrome is a string that reads the same backward as forward.

**Example:**
```
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
```

```
Input: s = "cbbd"
Output: "bb"
```

**Constraints:**
- `1 <= s.length <= 1000`
- `s` consists of only digits and English letters.

## Solution

To solve this problem, we can use the **Expand Around Center** approach:

1. Iterate through each character in the string, treating each character (and the gap between characters) as the center of a possible palindrome.
2. Expand outward from the center as long as the characters on both sides are equal.
3. Keep track of the longest palindrome found during the process.

**Time Complexity:** O(n²)  
**Space Complexity:** O(1)