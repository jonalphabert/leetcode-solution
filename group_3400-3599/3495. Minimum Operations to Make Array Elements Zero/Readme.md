## Problem

**Indonesia:**  
Diberikan sebuah array 2D `queries`, di mana `queries[i]` berbentuk `[l, r]`.  
Setiap `queries[i]` mendefinisikan sebuah array `nums` yang berisi elemen dari `l` hingga `r` (inklusif).

Dalam satu operasi, kamu dapat:
- Memilih dua bilangan bulat `a` dan `b` dari array.
- Mengganti keduanya dengan `floor(a / 4)` dan `floor(b / 4)`.

Tugasmu adalah menentukan jumlah minimum operasi yang diperlukan untuk membuat semua elemen array menjadi nol untuk setiap query.  
Kembalikan jumlah total hasil dari semua query.

---

## Solution

### Bahasa Indonesia

**Ide Berpikir:**
- Untuk setiap angka `x`, jumlah langkah minimum untuk membuatnya nol adalah `ceil(log4(x))`.
- Untuk sebuah deret dari `l` hingga `r`, kita ingin menghitung total langkah minimum untuk semua elemen.
- Fungsi `F(n)` menghitung jumlah langkah untuk semua angka dari 1 hingga n.
- Untuk interval `[l, r]`, jumlah langkahnya adalah `F(r) - F(l-1)`.
- Karena setiap operasi bisa mengurangi dua angka sekaligus, jumlah operasi minimum adalah `ceil((F(r) - F(l-1)) / 2)`.

### English

**Thinking Approach:**
- For each number `x`, the minimum steps to reduce it to zero is `ceil(log4(x))`.
- For a range from `l` to `r`, we want to compute the total minimum steps for all elements.
- The function `F(n)` computes the total steps for all numbers from 1 to n.
- For interval `[l, r]`, the total steps is `F(r) - F(l-1)`.
- Since each operation can reduce two numbers at once, the minimum number of operations is `ceil((F(r) - F(l-1)) / 2)`.

---

**Time Complexity:**  
- O(log n) per query (karena perhitungan blok log4).

**Space Complexity:**  
- O(1) (hanya variabel perhitungan).

---

**Catatan:**  
- Solusi ini memanfaatkan sifat logaritma basis 4 dan pengelompokan blok untuk menghitung jumlah