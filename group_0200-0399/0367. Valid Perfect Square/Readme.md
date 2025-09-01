## Problem

**Indonesia:**  
Diberikan sebuah bilangan bulat positif `num`, kembalikan `true` jika `num` adalah bilangan kuadrat sempurna, atau `false` jika bukan.

Bilangan kuadrat sempurna adalah bilangan bulat yang merupakan hasil perkalian suatu bilangan bulat dengan dirinya sendiri.

Tidak boleh menggunakan fungsi pustaka bawaan seperti `sqrt`.

**English:**  
Given a positive integer `num`, return `true` if `num` is a perfect square or `false` otherwise.

A perfect square is an integer that is the square of an integer. In other words, it is the product of some integer with itself.

You must not use any built-in library function, such as `sqrt`.

---

## Solution

### Bahasa Indonesia

Solusi umum adalah menggunakan **binary search**:
- Cari nilai tengah antara 1 dan `num`.
- Jika `mid * mid == num`, maka `num` adalah kuadrat sempurna.
- Jika `mid * mid < num`, cari di bagian kanan.
- Jika `mid * mid > num`, cari di bagian kiri.
- Ulangi sampai ditemukan atau tidak.

### English

A common solution is to use **binary search**:
- Search for the middle value between 1 and `num`.
- If `mid * mid == num`, then `num` is a perfect square.
- If `mid * mid < num`, search the right half.
- If `mid * mid > num`, search the left half.
- Repeat until found or not.

**Time Complexity:** O(log num)
**Space Complexity:** O(1)