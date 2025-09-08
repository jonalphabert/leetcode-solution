# 🇮🇩 (Bahasa Indonesia)

## 📌 Deskripsi Masalah

**No-Zero Integer** adalah bilangan bulat positif yang **tidak mengandung angka 0** dalam representasi desimalnya.

Diberikan sebuah bilangan bulat `n`, temukan dua bilangan bulat `[a, b]` yang memenuhi:

1. `a` dan `b` adalah **No-Zero Integers**.
2. `a + b = n`.

Dijamin bahwa selalu ada setidaknya satu solusi yang valid. Jika terdapat banyak solusi, kembalikan salah satu dari mereka.

---

## 💡 Ide Penyelesaian

1. Lakukan iterasi dari `1` hingga `n/2`.
2. Untuk setiap bilangan `i`, periksa apakah `i` dan `n-i` tidak mengandung angka 0.
3. Jika keduanya valid, kembalikan `[i, n-i]`.
4. Karena soal menjamin solusi ada, maka loop pasti berhenti dengan hasil.

---

## 🚀 Implementasi Go

```go
func checkAnyZeroDigit(num int) bool {
    for num > 0 {
        if num % 10 == 0 {
            return false
        }
        num /= 10
    }
    return true
}

func getNoZeroIntegers(n int) []int {
    for i := 1; i <= n/2; i++ {
        if checkAnyZeroDigit(i) && checkAnyZeroDigit(n-i) {
            return []int{i, n-i}
        }
    }
    return []int{}
}
```

---

## ⏱️ Kompleksitas

* **Waktu:** `O(n log n)`
  Iterasi hingga `n/2`, dengan pengecekan digit setiap bilangan.
* **Ruang:** `O(1)`
  Hanya menggunakan variabel integer tambahan.

---

## 📝 Contoh

```
Input: n = 11
Output: [2, 9]

Penjelasan: 
2 dan 9 adalah No-Zero Integers, dan 2 + 9 = 11.
```

---

# 🇬🇧 (English)

## 📌 Problem Description

A **No-Zero Integer** is a positive integer that **does not contain the digit 0** in its decimal representation.

Given an integer `n`, find two integers `[a, b]` such that:

1. Both `a` and `b` are **No-Zero Integers**.
2. `a + b = n`.

It is guaranteed that at least one valid solution exists. If there are multiple valid solutions, return any of them.

---

## 💡 Approach

1. Iterate from `1` to `n/2`.
2. For each number `i`, check whether both `i` and `n-i` contain no zero digits.
3. If both are valid, return `[i, n-i]`.
4. Since the problem guarantees a solution, the loop will always find a result.

---

## 🚀 Go Implementation

```go
func checkAnyZeroDigit(num int) bool {
    for num > 0 {
        if num % 10 == 0 {
            return false
        }
        num /= 10
    }
    return true
}

func getNoZeroIntegers(n int) []int {
    for i := 1; i <= n/2; i++ {
        if checkAnyZeroDigit(i) && checkAnyZeroDigit(n-i) {
            return []int{i, n-i}
        }
    }
    return []int{}
}
```

---

## ⏱️ Complexity

* **Time Complexity:** `O(n log n)`
  Iterates up to `n/2`, checking digits for each number.
* **Space Complexity:** `O(1)`
  Uses only a few integer variables.

---

## 📝 Example

```
Input: n = 11
Output: [2, 9]

Explanation: 
2 and 9 are No-Zero Integers, and 2 + 9 = 11.
```