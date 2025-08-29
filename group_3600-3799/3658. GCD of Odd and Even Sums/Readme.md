## Problem

**Indonesia:**  
Diberikan sebuah bilangan bulat `n`. Hitung GCD (greatest common divisor / FPB) dari dua nilai berikut:

- `sumOdd`: jumlah dari n bilangan ganjil pertama.
- `sumEven`: jumlah dari n bilangan genap pertama.

Kembalikan GCD dari `sumOdd` dan `sumEven`.

**English:**  
You are given an integer `n`. Compute the GCD (greatest common divisor) of two values:

- `sumOdd`: the sum of the first n odd numbers.
- `sumEven`: the sum of the first n even numbers.

Return the GCD of `sumOdd` and `sumEven`.

**Contoh:**
```
Input: n = 3
sumOdd = 1 + 3 + 5 = 9
sumEven = 2 + 4 + 6 = 12
Output: 3
```

---

## Solution

### Bahasa Indonesia

- Jumlah n bilangan ganjil pertama: `sumOdd = n^2`
- Jumlah n bilangan genap pertama: `sumEven = n * (n + 1)`
- GCD dari `n^2` dan `n*(n+1)` adalah `gcd(n, n+1)` dikali n (karena n dan n+1 selalu coprime, maka hasilnya adalah n).

**Langkah:**
1. Hitung `sumOdd = n * n`
2. Hitung `sumEven = n * (n + 1)`
3. Kembalikan `gcd(sumOdd, sumEven) = n`

### English

- The sum of the first n odd numbers: `sumOdd = n^2`
- The sum of the first n even numbers: `sumEven = n * (n + 1)`
- The GCD of `n^2` and `n*(n+1)` is `n` (since n and n+1 are coprime).

**Steps:**
1. Compute `sumOdd = n * n`
2. Compute `sumEven = n * (n + 1)`
3. Return `gcd(sumOdd, sumEven) = n`

**Sample Implementation (Go):**
```go
func gcdOfOddEvenSums(n int) int {
    return n
}
```