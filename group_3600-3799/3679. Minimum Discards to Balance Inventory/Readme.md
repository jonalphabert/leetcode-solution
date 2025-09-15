# 🇮🇩 Bahasa Indonesia

## 📌 Deskripsi Masalah

Diberikan sebuah stream data harian dengan panjang window `w`. Setiap hari masuk sebuah nilai, dan kita perlu menjaga agar hanya nilai dalam rentang `w` hari terakhir yang dipertahankan. Kita juga diberi parameter `m` (misalnya: menghitung jumlah distinct atau melakukan validasi tertentu) yang digunakan di dalam window.

Tantangan utama adalah bagaimana mengelola window dengan efisien tanpa harus memangkas array asli (karena itu akan mahal secara waktu).

---

## 💡 Ide Solusi

* Gunakan **queue** (implementasi slice di Go) untuk menyimpan `(day, value)` dari tiap data.
* Setiap iterasi hari baru (`day++`), lakukan:

  1. Tambahkan `(day, value)` ke queue.
  2. Hapus elemen di depan queue jika sudah keluar dari jendela (`day - w`).
* Karena `day` **selalu naik +1 per iterasi**, elemen yang expired paling banyak cuma satu per langkah → cukup pakai **`if`** (tidak perlu `for`).

---

## 📝 Pseudocode

```text
def minArrivalsToDiscard(arrivals, w, m):
    queue = []              // menyimpan (val, day)
    freq = map()            // frekuensi per nilai dalam window
    discarded = 0

    for day, val in enumerate(arrivals):   // day mulai dari 0
        // 1. Buang item lama di luar window
        if queue not empty AND queue.front.day <= day - w:
            old = queue.pop_front()
            freq[old.val]--

        // 2. Cek apakah item baru boleh disimpan
        if freq[val] < m:
            queue.push_back((val, day))
            freq[val]++
        else:
            discarded++

    return discarded

```

---

## ⏱️ Kompleksitas

* **Waktu:** `O(n)` karena setiap elemen hanya sekali masuk & keluar dari queue.
* **Memori:** `O(w)` karena queue hanya menyimpan elemen dalam window aktif.

---

# 🇬🇧 English Version

## 📌 Problem Description

We are given a daily data stream with a window size `w`. Each day, a new value arrives, and we must keep only the values within the last `w` days. We also have a parameter `m` (e.g., to count distinct values or validate constraints) used within the window.

The main challenge is managing the sliding window efficiently without pruning the original array, which would be costly.

---

## 💡 Solution Idea

* Use a **queue** (slice in Go) to store `(day, value)` pairs.
* For each new day (`day++`):

  1. Insert `(day, value)` into the queue.
  2. Remove the front element if it has expired (`day - w`).
* Since `day` **always increases by 1**, at most one element expires per step → so a simple **`if`** check is sufficient (no need for `for`).

---

## 📝 Pseudocode

```text
def minArrivalsToDiscard(arrivals, w, m):
    queue = []              // stores (val, day)
    freq = map()            // frequency of each value in the window
    discarded = 0

    for day, val in enumerate(arrivals):   // day starts from 0
        // 1. Remove expired items outside the window
        if queue not empty AND queue.front.day <= day - w:
            old = queue.pop_front()
            freq[old.val]--

        // 2. Check if the new item can be kept
        if freq[val] < m:
            queue.push_back((val, day))
            freq[val]++
        else:
            discarded++

    return discarded


```

---

## ⏱️ Complexity

* **Time:** `O(n)` since each element enters and leaves the queue at most once.
* **Space:** `O(w)` since the queue stores only elements within the active window.
