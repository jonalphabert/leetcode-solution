# 📖 Readme

## [Bahasa Indonesia] - [2221. Find Triangular Sum of an Array](https://leetcode.com/problems/find-triangular-sum-of-an-array/description/?envType=daily-question&envId=2025-09-30)

### Deskripsi 

Anda diberikan array bilangan bulat `nums` yang diindeks dari 0, di mana `nums[i]` adalah digit antara 0 dan 9 (inklusif).

Triangular sum dari `nums` adalah nilai dari elemen tunggal yang tersisa di `nums` setelah proses berikut selesai:
1. Misalkan `nums` terdiri dari `n` elemen. Jika `n == 1`, proses berakhir. Jika tidak, buat array bilangan bulat baru `newNums` yang diindeks dari 0 dengan panjang `n - 1`.
2. Untuk setiap indeks `i`, di mana `0 <= i < n - 1`, tetapkan nilai `newNums[i]` sebagai `(nums[i] + nums[i+1]) % 10`, di mana `%` adalah operator modulo.
3. Ganti array `nums` dengan `newNums`.
4. Ulangi seluruh proses mulai dari langkah 1.

Kembalikan triangular sum dari `nums`.

--- 

### Intuisi 

- Simulasi seperti perintah saja
- Ada solusi O(N) dengan menggunakan binomial coefficient. Mungkin akan ada di lain waktu

--- 

### Pendekatan 

- Iterasi `i` dari 0 sampai `n`
- Iterasi `j` dari 0 hingga `n-i-1`
- Update array` nums[j] = nums[j] + nums[j+1]`
- Kembalikan `nums[0]` sebagai hasil

--- 

### Kompleksitas 
- **Waktu:** O(N^2) 
- **Memori:** O(1) karena kita update langsung dari array tanpa tambahan array lain 
--- 

## [English] - [2221. Find Triangular Sum of an Array](https://leetcode.com/problems/find-triangular-sum-of-an-array/description/?envType=daily-question&envId=2025-09-30)

### Description 

You are given a 0-indexed integer array nums, where nums[i] is a digit between 0 and 9 (inclusive).

The triangular sum of nums is the value of the only element present in nums after the following process terminates:

Let nums comprise of n elements. If n == 1, end the process. Otherwise, create a new 0-indexed integer array newNums of length n - 1.
For each index i, where 0 <= i < n - 1, assign the value of newNums[i] as (nums[i] + nums[i+1]) % 10, where % denotes modulo operator.
Replace the array nums with newNums.
Repeat the entire process starting from step 1.
Return the triangular sum of nums.

--- 

### Intuition 

- Do simulation as explanation
- There is any solution with O(N) time complexity with binomial coefficient. Later will be added

--- 

### Approach 

- Iterate `i` from 0 to `n`
- Iterate `j` from 0 to `n-i-1`
- Update array` nums[j] = nums[j] + nums[j+1]`
- Return `nums[0]` as result

--- 

### Complexity 
- **Time Complexity:** O(N^2) 
- **Auxiliary Space:** O(1) because we update the array itself without using any addition array
---