## Problem

**Indonesia:**  
Diberikan sebuah string `s` dan sebuah pola `p`, implementasikan pencocokan regular expression dengan dukungan untuk karakter khusus berikut:
- `.` mencocokkan satu karakter apa saja.
- `*` mencocokkan nol atau lebih dari elemen sebelumnya.

Pencocokan harus meliputi seluruh string input (bukan sebagian).

**English:**  
Given an input string `s` and a pattern `p`, implement regular expression matching with support for:
- `.` Matches any single character.
- `*` Matches zero or more of the preceding element.

The matching should cover the entire input string (not partial).

---

## Solution

### Bahasa Indonesia

Solusi menggunakan **rekursi dengan memoization (DP)**:
- Fungsi rekursif `dfs(i, j)` memeriksa apakah substring `s[0..i]` cocok dengan pola `p[0...j]`.
- Karena fungsi ini bisa dipanggil berulang dengan state yang sama, kita gunakan memoization dengan state `(i, j)` agar tidak menghitung ulang.
- Untuk setiap posisi:
  - Jika karakter berikutnya di pola adalah `*`, ada dua pilihan:
    - Abaikan pasangan karakter dan `*` (`dfs(i, j+2)`).
    - Jika karakter cocok, lanjutkan ke karakter berikutnya di string (`dfs(i+1, j)`).
  - Jika tidak ada `*`, lanjutkan pencocokan satu karakter dan rekursi ke posisi berikutnya.
- Hasil akhir adalah apakah seluruh string dan pola bisa dicocokkan.

### English

The solution uses **recursion with memoization (DP)**:
- The recursive function `dfs(i, j)` checks if the substring `s[0..i]` matches the pattern `p[0..j]`.
- Since this function can be called multiple times with the same state, we use memoization with state `(i, j)` to avoid redundant calculations.
- For each position:
  - If the next character in the pattern is `*`, there are two options:
    - Skip the character and `*` (`dfs(i, j+2)`).
    - If the character matches, move to the next character in the string (`dfs(i+1, j)`).
  - If there is no `*`, match one character and recurse to the next position.
- The final result is whether the entire string and pattern can be matched.

**Time Complexity:** O(len(s) * len(p))  
**Space Complexity:** O(len(s))