## 🇮🇩 Bahasa Indonesia

### 📌 Deskripsi Masalah

Diberikan sebuah string `s` yang berisi karakter inggris dengan lowercase

Kamu dapat melakukan operasi sebanyak apapun dengan cara melakukan pergantian suatu karakter `c` di dalam string `s` dan mengubah semua kemunculan `c` ke huruf selanjutnya

Kamu diminta untuk mengembalikan nilai terkecil dari operasi yang kamu jalankan untuk mengubah semua karakter yang ada di string `s` menjadi karakter 'a' semua

Catatan: Dalam masalah ini anggap saja kita dapat melakukan perulangan dari karakter 'z' ke karakter 'a'

---

### 🔍 Abstraksi & Ide

Untuk mendapatkan nilai minimum operasi yang dilakukan, maka kita harus lakukan operasi yang diminta mulai dari karakter yang paling jauh dari karakter 'a'.

Karena operasi yang dilakukan hanya bisa dengan mengubah suatu karakter ke karakter selanjutnya (dalam order ASCII) maka dapat kita simpulkan 'b' adalah karakter terjauh sedangkan 'z' adalah karakter terdekat karena cyclicnya

Sehingga rumus minimum dari operasinya ada
```
    min_ops = max(min_ops, (26 - (s[i] - 'a'))%26)
```
---

### 💡 Penyelesaian

Implementasi dapat dilakukan dengan cara seperti berikut:

* Iterasi semua elemen dalam array.
* Lakukan update variabel `result` operasi max terhadap variabel `result` dan rumus yang kita dapatkan tadi
* Kembalikan nilai `result`

Kode Go (ringkas):

```go
func minOperations(s string) int {
    res := -1
    for _, val := range s {
        res = max(res, (26 - int(val - 'a'))%26)
    }
    return res
}
```

---

### ⏱️ Kompleksitas

* **Waktu:** `O(n)`
  Setiap elemen pada variabel `s` akan diiterasi sekali
* **Ruang:** `O(1)`
  Memory tambahan yang dibutuhkan pada solusi ini hanyalah tipe data primitif berupa int saja

---


## 🇬🇧 English

### 📌 Problem Description

You are given a string `s` consisting of lowercase English letters.

You can perform the following operation any number of times:

* Select a character `c` in the string `s` and replace **all occurrences** of `c` with the **next character** in the alphabet (cyclic, meaning `'z'` wraps around to `'a'`).

Your task is to return the **minimum number of operations** required to transform the entire string `s` into a string containing only the character `'a'`.

---

### 🔍 Abstraction & Idea

To minimize the number of operations, we need to consider the **farthest character from `'a'`** in cyclic order.

* Since the operation only allows moving forward to the next character, `'b'` is actually the farthest character (requires 25 steps to become `'a'`).
* On the other hand, `'z'` is the closest character (requires just 1 step to become `'a'` because of the wrap-around).

Thus, the minimum number of operations can be expressed as:

```
min_ops = max(min_ops, (26 - (s[i] - 'a')) % 26)
```

---

### 💡 Solution

We can implement the solution as follows:

1. Iterate over each character in the string `s`.
2. For each character, calculate the distance to `'a'` using the formula above.
3. Keep track of the maximum distance (this will be the minimum operations required).
4. Return the result.

#### Go Implementation

```go
func minOperations(s string) int {
    res := -1
    for _, val := range s {
        res = max(res, (26 - int(val - 'a')) % 26)
    }
    return res
}
```

---

### ⏱️ Complexity Analysis

* **Time Complexity:** `O(n)`
  We iterate through the string once.

* **Space Complexity:** `O(1)`
  Only a few integer variables are used as extra space.


