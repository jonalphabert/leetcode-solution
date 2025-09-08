# 🇮🇩 Bahasa Indonesia

## Deskripsi

Diberikan sebuah string yang hanya berisi karakter `'('` dan `')'`.
Tugasnya adalah mengembalikan panjang substring tanda kurung yang valid (well-formed) terpanjang.

Substring valid berarti setiap `'('` memiliki pasangan `')'` yang sesuai, dan urutannya benar.

## Contoh

### Contoh 1

```
Input: "(()"
Output: 2
Penjelasan: Substring valid terpanjang adalah "()"
```

### Contoh 2

```
Input: ")()())"
Output: 4
Penjelasan: Substring valid terpanjang adalah "()()"
```

### Contoh 3

```
Input: "((()))"
Output: 6
Penjelasan: Seluruh string adalah substring valid.
```

---

## Pendekatan

### Dynamic Programming (DP)

* Definisikan `dp[i]` sebagai panjang substring valid terpanjang yang **berakhir** di indeks `i`.
* Aturan transisi:

  1. Jika `s[i] == ')'` dan `s[i-1] == '('`:

     ```
     dp[i] = dp[i-2] + 2
     ```
  2. Jika `s[i] == ')'` dan `s[i-1] == ')'`:
     cek apakah ada `'('` yang berpasangan dengan `s[i]`:

     ```
     jika s[i - dp[i-1] - 1] == '(':
         dp[i] = dp[i-1] + 2 + dp[i - dp[i-1] - 2]
     ```
---

## Kompleksitas

* **Waktu**: `O(n)`
* **Memori**: `O(n)`

---

# 🇬🇧 English Version

## Description

Given a string containing only characters `'('` and `')'`,
return the length of the longest valid (well-formed) parentheses substring.

A valid substring means every `'('` has a matching `')'` in the correct order.

## Examples

### Example 1

```
Input: "(()"
Output: 2
Explanation: The longest valid substring is "()"
```

### Example 2

```
Input: ")()())"
Output: 4
Explanation: The longest valid substring is "()()"
```

### Example 3

```
Input: "((()))"
Output: 6
Explanation: The whole string is a valid substring.
```

---

## Approach

### Dynamic Programming (DP)

* Define `dp[i]` as the length of the longest valid substring **ending** at index `i`.
* Transition rules:

  1. If `s[i] == ')'` and `s[i-1] == '('`:

     ```
     dp[i] = dp[i-2] + 2
     ```
  2. If `s[i] == ')'` and `s[i-1] == ')'`:
     check if there's a matching `'('` before:

     ```
     if s[i - dp[i-1] - 1] == '(':
         dp[i] = dp[i-1] + 2 + dp[i - dp[i-1] - 2]
     ```

---

## Complexity

* **Time**: `O(n)`
* **Space**: `O(n)`

