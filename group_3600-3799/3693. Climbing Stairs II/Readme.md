# 📖 Readme

## [Bahasa] - [Judul Problem] 

### Deskripsi 

Kamu menaiki sebuah tangga dengan `n+1` langkah, dimana langkah tersebut dimulai dari 0 hingga `n`

Kamu juga diberikan sebuah array `costs` yang berindeks 1 dengan panjang `n`, dimana `costs[i]` adalah biaya dari langkah `i`.

Dari langkah `i`, kamu hanya dapat melompat ke langkah `i + 1`, `i + 2`, atau `i + 3`. Biaya untuk melompat dari langkah `i` ke langkah `j` dide

finisikan sebagai: `costs[j] + (j - i)^2`

Kamu memulai dari langkah 0 dengan biaya = 0.

Kembalikan biaya total minimum untuk mencapai langkah `n`.

--- 

### Intuisi 

- Kita mau sampai ke step n dengan biaya seminimal mungkin.
- Dari setiap step i, kita bisa loncat ke i+1, i+2, atau i+3. Tapi setiap loncatan ada biayanya:
    - biaya berdiri di step tujuan (costs[j])
    - ditambah penalti besar lompatan (j - i)².
- Jadi, kalau kita sudah tahu biaya minimum untuk sampai ke step lebih kecil (misalnya i-1, i-2, i-3), kita bisa hitung biaya ke i dengan memilih loncatan terakhir yang paling murah.
- Intinya: biaya minimum ke step i = biaya minimum ke salah satu step sebelumnya + biaya loncat ke i.

--- 

### Pendekatan 

- Definisikan array `bestRoute[n+1]`
- Inisiasikan `bestRoute[0] = 0`
- Iterasi dari `1` hingga `n`
  - Untuk setiap `i`, iterasi dari `1` hingga `3`
    - Jika `i - j >= 0`, maka hitung `bestRoute[i] = min(bestRoute[i], bestRoute[i-j] + costs[i] + j*j)`
- Kembalikan `bestRoute[n]`

--- 

### Kompleksitas 
- **Waktu:** O(N) 
- **Memori:** O(N)
--- 

### Contoh 

- Input

- Output

## [English] - [Judul Problem] 

### Description 

You are climbing a staircase with n + 1 steps, numbered from 0 to n.

You are also given a 1-indexed integer array costs of length n, where costs[i] is the cost of step i.

From step i, you can jump only to step i + 1, i + 2, or i + 3. The cost of jumping from step i to step j is defined as: costs[j] + (j - i)2

You start from step 0 with cost = 0.

Return the minimum total cost to reach step n.

--- 

### Intuition 

- We need to step `n` with optimal cost
- From each step `i`, we can jump to `i+1`, `i+2`, or `i+3`. But each jump has a cost:
    - the cost of standing on the destination step (`costs[j]`)
    - plus a big penalty for the jump distance (`(j - i)²`).
- So, if we already know the minimum cost to reach smaller steps (like `i-1`, `i-2`, `i-3`), we can compute the cost to `i` by picking the cheapest last jump.
- The idea: minimum cost to step `i` = minimum cost to one of the previous steps + cost to jump to `i`.

--- 

### Approach 

- Define array `bestRoute[n+1]`
- Set `bestRoute[0] =0`
- Iterate `i` from 1 to `n`
    - Iterate `j` step from 1 to 3
        - If `i-h` > 0, do compute `bestRoute[i] = min(bestRoute[i], bestRoute[i-j] + costs[i] + j*j)`
- Return `bestRoute[n]`

--- 

### Complexity 
- **Time Complexity:** O(N) 
- **Auxiliary Space:** O(N)
---