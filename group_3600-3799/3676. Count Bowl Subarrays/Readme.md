## 🇮🇩 Bahasa Indonesia

### 📌 Deskripsi Masalah

Diberikan sebuah array integer `nums` dengan elemen **unik**. Kita mendefinisikan sebuah subarray `nums[l...r]` sebagai **bowl** jika memenuhi dua syarat:

1. Panjang subarray minimal 3 → `r - l + 1 >= 3`
2. Nilai minimum dari kedua ujung lebih besar dari nilai maksimum elemen di tengahnya →
   `min(nums[l], nums[r]) > max(nums[l+1 ... r-1])`

Tujuannya adalah menghitung **jumlah bowl subarray** di dalam `nums`.

---

### 🔍 Abstraksi & Ide

Jika divisualisasikan, sebuah **bowl** menyerupai bentuk cekungan:

```
nums[l]   nums[r]
   \       /
    \_____/
```

* Kedua ujung lebih tinggi dibandingkan semua elemen di tengah.
* Dengan kata lain, setiap kali kita menemukan "lembah" di tengah, itu bisa membentuk bowl dengan beberapa ujung tertentu.

Untuk mengenali pola ini, kita bisa menggunakan **monotonic stack**:

* Stack menjaga indeks elemen dengan sifat tertentu (monoton menurun atau meningkat).
* Saat menemukan elemen yang lebih besar, kita mempop stack sampai kondisi valid.
* Setiap kali kondisi "cekungan" terdeteksi, kita tambahkan hasil.

---

### 💡 Penyelesaian

Implementasi menggunakan **monotonic stack**:

* Iterasi semua elemen dalam array.
* Gunakan stack untuk melacak kandidat batas bowl.
* Ketika menemukan elemen yang lebih besar dari stack top, lakukan pop dan hitung kemungkinan bowl.
* Ulangi sampai tidak ada lagi pelanggaran monotonic.
* Kembalikan total jumlah bowl.

Kode Go (ringkas):

```go
func bowlSubarrays(nums []int) int64 {
    stack := []int{}
    result := int64(0)

    for i := 0; i < len(nums); i++ {
        for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
            stack = stack[:len(stack)-1]
            if len(stack) == 0 {
                break
            }
            result++
        }
        stack = append(stack, i)
    }

    return result
}
```

---

### ⏱️ Kompleksitas

* **Waktu:** `O(n)`
  Setiap elemen dimasukkan dan dikeluarkan dari stack paling banyak sekali.
* **Ruang:** `O(n)`
  Stack menyimpan maksimal `n` indeks pada kasus terburuk.

---

## 🇬🇧 English

### 📌 Problem Description

You are given an integer array `nums` with **distinct elements**.
A subarray `nums[l...r]` is called a **bowl** if:

1. The length of the subarray is at least 3 → `r - l + 1 >= 3`
2. The minimum of its two ends is strictly greater than the maximum of all elements in between →
   `min(nums[l], nums[r]) > max(nums[l+1 ... r-1])`

The goal is to return the number of **bowl subarrays** in `nums`.

---

### 🔍 Abstraction & Idea

A bowl looks like a **valley shape**:

```
nums[l]   nums[r]
   \       /
    \_____/
```

* Both ends are higher than every element in between.
* This naturally suggests using a **monotonic stack** to track candidates for valid bowls.

---

### 💡 Solution

We apply a **monotonic stack** strategy:

* Iterate through the array.
* Maintain a stack of indices to represent decreasing order.
* When a larger element is encountered, pop from the stack and check possible bowls.
* Each valid "valley" contributes to the answer.
* Return the final count.

Go implementation:

```go
func bowlSubarrays(nums []int) int64 {
    stack := []int{}
    result := int64(0)

    for i := 0; i < len(nums); i++ {
        for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
            stack = stack[:len(stack)-1]
            if len(stack) == 0 {
                break
            }
            result++
        }
        stack = append(stack, i)
    }

    return result
}
```

---

### ⏱️ Complexity

* **Time:** `O(n)`
  Each element is pushed and popped at most once.
* **Space:** `O(n)`
  The stack stores up to `n` indices in the worst case.