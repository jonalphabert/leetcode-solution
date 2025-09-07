## 🇮🇩 Bahasa Indonesia

### 📌 Deskripsi Masalah

Diberikan sebuah bilangan bulat `n`. Tugas kita adalah mengembalikan sebuah array berisi `n` bilangan bulat **unik** sedemikian sehingga jumlah seluruh elemen sama dengan `0`.

Contoh:

* Input: `n = 5`
  Output: `[0, 1, -1, 2, -2]` (atau variasi urutan lain)
* Input: `n = 4`
  Output: `[1, -1, 2, -2]`

---

### 🔍 Abstraksi & Ide

Kita bisa memanfaatkan **pola simetri**:

* Jika `n` **ganjil**, maka kita **wajib menambahkan `0`** ke dalam hasil, agar jumlah akhirnya tetap seimbang.
* Untuk elemen lainnya, kita cukup memasangkan angka positif dan negatif dari `1` hingga `n/2`.

  * Misalnya, untuk `n = 6`, kita dapatkan pasangan `(1, -1), (2, -2), (3, -3)`
  * Total dijamin **0** karena setiap pasangan bernilai 0 dan jika ada `0`, tetap seimbang.

---

### 💡 Penyelesaian

Implementasi Go:

```go
func sumZero(n int) []int {
    res := make([]int, 0)

    if n%2 == 1 {
        res = append(res, 0) // tambahkan 0 jika n ganjil
    }

    for i := 1; i <= n/2; i++ {
        res = append(res, i)
        res = append(res, -i)
    }

    return res
}
```

---

### ⏱️ Kompleksitas

* **Waktu:** `O(n)` → kita melakukan iterasi sebanyak `n/2`.
* **Ruang:** `O(n)` → array hasil menyimpan `n` elemen unik.

---

## 🇬🇧 English

### 📌 Problem Description

Given an integer `n`, return an array containing `n` **unique integers** such that their sum equals `0`.

Example:

* Input: `n = 5`
  Output: `[0, 1, -1, 2, -2]` (or any permutation)
* Input: `n = 4`
  Output: `[1, -1, 2, -2]`

---

### 🔍 Abstraction & Idea

We can rely on **symmetry**:

* If `n` is **odd**, we must include `0` to balance the sum.
* For the rest, simply generate pairs of positive and negative numbers from `1` to `n/2`.

  * Example: for `n = 6`, pairs are `(1, -1), (2, -2), (3, -3)`
  * The sum is guaranteed to be `0`.

---

### 💡 Solution

Go implementation:

```go
func sumZero(n int) []int {
    res := make([]int, 0)

    if n%2 == 1 {
        res = append(res, 0) // add 0 if n is odd
    }

    for i := 1; i <= n/2; i++ {
        res = append(res, i)
        res = append(res, -i)
    }

    return res
}
```

---

### ⏱️ Complexity

* **Time:** `O(n)` → we iterate up to `n/2`.
* **Space:** `O(n)` → result array stores exactly `n` unique elements.
