## Problem

**Indonesia:**  
Kamu adalah seorang product manager yang memimpin tim pengembangan produk baru. Sayangnya, versi terbaru produkmu gagal dalam quality check. Karena setiap versi dikembangkan dari versi sebelumnya, maka semua versi setelah versi yang buruk juga pasti buruk.

Diberikan n versi `[1, 2, ..., n]`, kamu ingin mencari tahu versi pertama yang buruk, karena versi inilah yang menyebabkan semua versi setelahnya menjadi buruk.

Tersedia API `bool isBadVersion(version)` yang mengembalikan apakah suatu versi buruk.  
Implementasikan fungsi untuk menemukan versi buruk pertama dengan meminimalkan jumlah pemanggilan API.

**English:**  
You are a product manager leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check. Since each version is developed based on the previous version, all the versions after a bad version are also bad.

Given n versions `[1, 2, ..., n]`, you want to find out the first bad one, which causes all the following ones to be bad.

You are given an API `bool isBadVersion(version)` which returns whether a version is bad.  
Implement a function to find the first bad version. You should minimize the number of calls to the API.

---

## Solution

### Bahasa Indonesia

Solusi menggunakan **binary search**:
- Inisialisasi batas kiri (`left = 1`) dan kanan (`right = n`).
- Selama `left < right`, lakukan:
  - Hitung nilai tengah (`mid`).
  - Jika `isBadVersion(mid)` bernilai `true`, berarti versi buruk pertama ada di kiri atau di `mid`, geser batas kanan ke `mid`.
  - Jika `false`, berarti versi buruk pertama ada di kanan, geser batas kiri ke `mid + 1`.
- Ketika loop selesai, `left` adalah versi buruk pertama.

### English

The solution uses **binary search**:
- Initialize left boundary (`left = 1`) and right boundary (`right = n`).
- While `left < right`, do:
  - Compute the middle value (`mid`).
  - If `isBadVersion(mid)` is `true`, the first bad version is at `mid` or to the left, move right boundary to `mid`.
  - If `false`, the first bad version is to the right, move left boundary to `mid + 1`.
- When the loop ends, `left` is the first bad version.

**Time Complexity:** O(log N)
**Space Complexity:** O(1)