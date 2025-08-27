# Problem

## Inggris
You are given a 2D integer matrix grid of size n x m, where each element is either 0, 1, or 2.

A V-shaped diagonal segment is defined as:

- The segment starts with 1.
- The subsequent elements follow this infinite sequence: 2, 0, 2, 0, ....
- The segment 
    - Starts along a diagonal direction (top-left to bottom-right, bottom-right to top-left, top-right to bottom-left, or bottom-left to top-right).
    - Continues the sequence in the same diagonal direction.
    - Makes at most one clockwise 90-degree turn to another diagonal direction while maintaining the sequence.


Return the length of the longest V-shaped diagonal segment. If no valid segment exists, return 0.

## Indonesia
Diberikan sebuah matriks 2D grid berukuran n x m, di mana setiap elemen bernilai 0, 1, atau 2.

Sebuah segmen diagonal berbentuk V didefinisikan sebagai berikut:

- Segmen selalu dimulai dengan angka 1.
- Elemen-elemen berikutnya mengikuti pola urutan tak hingga: `2, 0, 2, 0, 2, 0, ...`
- Segmen:
    - Berawal dari salah satu arah diagonal:
        - kiri-atas ke kanan-bawah
        - kanan-bawah ke kiri-atas
        - kanan-atas ke kiri-bawah
        - kiri-bawah ke kanan-atas
    - Meneruskan urutan angka sesuai pola pada arah diagonal tersebut.
    - Boleh melakukan paling banyak satu belokan searah jarum jam sebesar 90 derajat ke arah diagonal lain, tetapi tetap harus menjaga pola urutan angkanya.

Tugas:
Kembalikan panjang segmen diagonal berbentuk V yang paling panjang.
Jika tidak ada segmen valid, kembalikan 0.


# 📌 Intuisi Penyelesaian

Penyelesaian masalah ini dapat dilakukan dengan DFS(Depth First Search). Jadi kita traverse peta tersebut dengan kondisi yang ditentukan.

## Dynamic Programming Approach

Untuk optimasinya, ada kemungkinan sebuah petak dilewati lebih dari satu kali. Disini DP dapat berperan penting. DP approach yang dapat digunakan dalam problem ini adalah jarak terjauh di suatu titik dengan direction tertentu dan kondisi sudah berpindah arah berlawanan jarum jam atau belum.

```
dirs = {1, 1}, {1, -1}, {-1, -1}, {-1, 1}

dp(r, c, dir, turnUsed)

/*
    dp => Jumlah step pada titik r dan c dengan arah dir dan juga jumlah belokan yang terjadi
    dir={0,1,2,3} => index dari dirs
    turnUsed= {0,1} => Jumlah belokan yang telah terjadi hingga di titik r, c
*/
```