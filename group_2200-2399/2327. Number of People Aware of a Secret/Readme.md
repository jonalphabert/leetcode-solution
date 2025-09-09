## 📖 ID (Bahasa Indonesia)

### Deskripsi

Masalah ini berasal dari **LeetCode 2327 — Number of People Aware of a Secret**.

Pada hari ke-1, satu orang mengetahui sebuah rahasia.
Setiap orang:

* Mulai **membagikan rahasia** ke orang baru setelah `delay` hari sejak pertama tahu.
* Akan **lupa rahasia** setelah `forget` hari sejak pertama tahu.
* Orang yang lupa tidak bisa membagikan rahasia lagi.

Diberikan `n`, `delay`, dan `forget`, hitunglah **berapa banyak orang yang masih tahu rahasia di akhir hari ke-n**.
Jawaban dikembalikan dengan **modulo 10^9 + 7**.

---

### Intuisi

Alih-alih menyimpan total orang yang tahu rahasia setiap hari, kita cukup menyimpan:

* `new[day]` = jumlah orang yang **pertama kali tahu** pada hari tersebut.

Kenapa?
Karena seseorang yang tahu pada hari `j` hanya bisa membagikan pada rentang:

```
j + delay  ≤  t  ≤  j + forget - 1
```

Sehingga untuk hari `i`, orang baru = jumlah `new[j]` dalam window `[i - forget + 1, i - delay]`.

---

### Pendekatan

* Gunakan **sliding window** untuk menjaga jumlah orang yang bisa membagikan rahasia (`shareable`).
* Setiap hari:

  * Tambah orang yang baru boleh membagikan (`day - delay`).
  * Kurangi orang yang lupa (`day - forget`).
* `new[day] = shareable`.
* Pada akhirnya, jumlah orang yang masih tahu rahasia = `sum(new[n - forget + 1 .. n])`.

---

### Kompleksitas

* **Waktu:** O(n)
* **Memori:** O(n) (bisa dioptimasi jadi O(forget) dengan ring buffer)

---

### Contoh

Input:

```
n = 11, delay = 4, forget = 7
```

Output:

```
9
```

---

## 📖 EN (English)

### Description

This problem comes from **LeetCode 2327 — Number of People Aware of a Secret**.

On day 1, one person knows a secret.
Each person:

* Starts **sharing the secret** with a new person after `delay` days since discovering it.
* **Forgets the secret** after `forget` days since discovering it.
* Once forgotten, they cannot share anymore.

Given `n`, `delay`, and `forget`, return **the number of people who still remember the secret at the end of day `n`**, modulo 10^9 + 7.

---

### Intuition

Instead of tracking the total number of people who know the secret each day, we only track:

* `new[day]` = the number of people who **first learn the secret** on that day.

Why?
Because someone who learns on day `j` can share only within the range:

```
j + delay  ≤  t  ≤  j + forget - 1
```

So for day `i`, new people = sum of `new[j]` in the window `[i - forget + 1, i - delay]`.

---

### Approach

* Use a **sliding window** to maintain the number of people currently able to share (`shareable`).
* For each day:

  * Add people who just became eligible to share (`day - delay`).
  * Remove people who forgot the secret (`day - forget`).
* `new[day] = shareable`.
* Finally, the number of people who still remember = `sum(new[n - forget + 1 .. n])`.

---

### Complexity

* **Time:** O(n)
* **Space:** O(n) (can be optimized to O(forget) with a ring buffer)

---

### Example

Input:

```
n = 11, delay = 4, forget = 7
```

Output:

```
9
```
