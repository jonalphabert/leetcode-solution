# 🇮🇩 Bahasa Indonesia

### 📌 Deskripsi Masalah

Alice dan Bob memainkan sebuah permainan dengan sebuah string.

Diberikan sebuah string `s`, aturan permainan adalah sebagai berikut:

* Alice selalu bermain pertama.
* Pada gilirannya, **Alice** harus menghapus substring tidak kosong yang mengandung jumlah huruf vokal sebanyak **bilangan ganjil**.
* Pada gilirannya, **Bob** harus menghapus substring tidak kosong yang mengandung jumlah huruf vokal sebanyak **bilangan genap**.
* Pemain yang tidak dapat melakukan langkah akan kalah.

Asumsikan kedua pemain bermain secara **optimal**.
Tentukan apakah Alice bisa menang.

---

### 🔍 Intuisi & Ide

Masalah ini berkaitan dengan teori bilangan (paritas ganjil/genap).

Beberapa fakta penting:

* **Bilangan genap** dapat dibentuk dari `ganjil + ganjil` atau `genap + genap`.
* **Bilangan ganjil** dapat dibentuk dari `ganjil + genap`.

Dari sini terlihat bahwa **bilangan ganjil berperan penting** karena menjadi kunci perubahan paritas.

* Jika jumlah huruf vokal dalam string **genap**, Alice bisa menang.
  Karena Alice bermain pertama, ia dapat memilih langkah sehingga jumlah vokal tersisa selalu ganjil. Pada kondisi itu, Bob tidak bisa mengambil semua substring tersisa (karena sisa vokal ganjil).

* Jika jumlah huruf vokal dalam string **ganjil**, satu-satunya cara Alice untuk menang adalah dengan langsung menghapus seluruh string `s`.
  Jika ia hanya mengambil sebagian, maka Bob bisa langsung mengambil sisa string dan menang.

* Alice **kalah** jika jumlah huruf vokal adalah **0**.
  Karena ia bermain pertama, ia tidak bisa melakukan langkah apapun.

---

### 💡 Penyelesaian

Penyelesaiannya sederhana:

* Iterasi semua karakter dalam string `s`.
* Jika terdapat huruf vokal, Alice **pasti bisa menang**.
* Jika tidak ada huruf vokal sama sekali, Alice **kalah**.

Kode Go (ringkas):

```go
func doesAliceWin(s string) bool {
    for _, char := range s {
        if char == 'a' || char == 'i' || char == 'u' || char == 'e' || char == 'o' {
            return true
        }
    }
    return false
}
```

---

### ⏱️ Kompleksitas

* **Waktu:** `O(n)`
  Karena setiap karakter string `s` diiterasi sekali.
* **Ruang:** `O(1)`
  Tidak ada penggunaan memori tambahan selain variabel iterasi.

---

# 🇬🇧 English

### 📌 Problem Description

Alice and Bob are playing a game with a string.

Given a string `s`, the game rules are as follows:

* Alice always plays first.
* On her turn, **Alice** must remove a non-empty substring that contains an **odd** number of vowels.
* On his turn, **Bob** must remove a non-empty substring that contains an **even** number of vowels.
* The player who cannot make a move loses the game.

Assume both players play **optimally**.
Determine whether Alice can win.

---

### 🔍 Intuition & Idea

This problem is related to number theory (odd/even parity).

Key observations:

* **Even numbers** can be formed as `odd + odd` or `even + even`.
* **Odd numbers** can be formed as `odd + even`.

Thus, **odd numbers play a central role**, since they determine the parity transitions.

* If the total number of vowels in the string is **even**, Alice can win.
  Because she plays first, she can make a move such that the remaining vowel count is always odd, preventing Bob from removing the whole remaining string.

* If the total number of vowels is **odd**, the only way Alice can win is by removing the **entire string** on her first move.
  If she removes only part of it, Bob can immediately take the rest and win.

* Alice **loses** if the total number of vowels is **0**.
  Since she plays first, she has no valid move.

---

### 💡 Solution

The solution is straightforward:

* Iterate through all characters in `s`.
* If there is at least one vowel, Alice **can win**.
* If there are no vowels at all, Alice **loses**.

Go code (concise):

```go
func doesAliceWin(s string) bool {
    for _, char := range s {
        if char == 'a' || char == 'i' || char == 'u' || char == 'e' || char == 'o' {
            return true
        }
    }
    return false
}
```

---

### ⏱️ Complexity

* **Time:** `O(n)`
  We iterate through each character of the string once.
* **Space:** `O(1)`
  Only a few variables are used for iteration.