# 📖 Readme

## [Bahasa Indonesia] - [3694. Distinct Points Reachable After Substring Removal](https://leetcode.com/problems/distinct-points-reachable-after-substring-removal/description/)

### Deskripsi 

Kamu diberikan sebuah string `s` yang berisikan karakter `U`, `L`, `R`, dan `D` yang dimana merupakan repesentasi arah gerak sebuah karakter di 2 dimensi Catersian

`U`: Bergerak dari `(x, y)` ke `(x, y + 1)`
`D`: Bergerak dari `(x, y)` ke `(x, y - 1)`
`L`: Bergerak dari `(x, y)` ke `(x - 1, y)`
`R`: Bergerak dari `(x, y)` ke `(x + 1, y)`

Kamu juga diberikan bilangan positif `k`

Kamu ditugaskan untuk menghilangkan `k` substring lanjutan dari `s`. Semua karakter akan berjalan melalui titik (0,0)

Tentukan berapa titik akhir yang berbeda setelah melakukan penghapusan substring sebanyak `k` karakter

--- 

### Intuisi 

- Kita simpan prefix sum untuk perpindahan x dan y untuk setiap iterasi
- Untuk query kita bisa gunakan inclusion + exclusion

--- 

### Pendekatan 

- Buat prefix sum array prefX dan prefY
- Iterasi setiap karakter di string `s`
    - Kalkulasi `prefX[i], prefY[i] = prefX[i-1] + deltaX, prefY[i-1] + deltaY`
- Iterasi `i` dari 0 hingga `n-k`
    - Kalkulasi perubahan perpindahan dengan cara
        ```
            dx := px[l+k] - px[l]
            dy := py[l+k] - py[l]
            final := [2]int{X - dx, Y - dy}
        ```
    - Simpan nilai akhir ke map `seen` sehingga mencegah adanya nilai akhir yang sama
- Kembalikan panjang map `seen` untuk banyaknya titik akhir seen.

--- 

### Kompleksitas 
- **Waktu:** O(N) 
- **Memori:** O(N)
--- 

## [English] - [3694. Distinct Points Reachable After Substring Removal](https://leetcode.com/problems/distinct-points-reachable-after-substring-removal/description/)

### Description 

You are given a string s consisting of characters 'U', 'D', 'L', and 'R', representing moves on an infinite 2D Cartesian grid.

'U': Move from (x, y) to (x, y + 1).
'D': Move from (x, y) to (x, y - 1).
'L': Move from (x, y) to (x - 1, y).
'R': Move from (x, y) to (x + 1, y).
You are also given a positive integer k.

You must choose and remove exactly one contiguous substring of length k from s. Then, start from coordinate (0, 0) and perform the remaining moves in order.

Return an integer denoting the number of distinct final coordinates reachable.

--- 

### Intuition 

- We save prefix sum for each movement, so prefix sum hold the final position after `i` step
- We need query for `k` length substring removal using inclusion and exclusion

--- 

### Approach 

- Build prefX and prefY
- For each character in `s`
    - Calculate the prefix sum array with `prefX[i], prefY[i] = prefX[i-1] + deltaX, prefY[i-1] + deltaY`
- Iterate `l` from 0 to `n-k` 
    - Calculate the changes final position after `l` to `l+k` substring removal using this formula
         ```
            dx := px[l+k] - px[l]
            dy := py[l+k] - py[l]
            final := [2]int{X - dx, Y - dy}
        ```
    - Save final position to map `seen` for preventing same value of final position duplicate
- Return the length of map `seen` for distinct final position count

--- 

### Complexity 
- **Time Complexity:** O(N) 
- **Auxiliary Space:** O(N)
---