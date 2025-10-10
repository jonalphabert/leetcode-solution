# 📖 Readme

## [Bahasa Indonesia] - [3147. Taking Maximum Energy From the Mystic Dungeon](https://leetcode.com/problems/taking-maximum-energy-from-the-mystic-dungeon/description/?envType=daily-question&envId=2025-10-10)

### Deskripsi 

Terdapat dungeon dengan `n` penyihir. Penyihir tersebut dapat menyerap atau memberikan mana ke dirimu.

Ajaibnya, ketika kamu datang ke penyihir `i` maka kamu akan langsung berpindah ke penyihir `i+k`. Proses ini akan berlanjut hingga kita tidak menemukan penyihir di `i+k`

Kamu akan diberikan array `energy` yang merepresentasikan jumlah mana yang diserap atau diberikan dari penyihir dan juga `k` yang merupakan step teleportasi.

Kamu bebas memilih titik awal manapun untuk memulai perjalananmu. Tujuanmu adalah mengumpulkan jumlah mana maksimum selama menjelajahi dungeon tersebut.

--- 

### Intuisi 

- Konsep utama dalam penyelesaian problem disini adalah prefix sum
- Prefix sum dilakukan dengan formula `preSum[i] = energy[i] + preSum[i-k]`

--- 

### Pendekatan 

- Iterasi array `energy` untuk mendapatkan prefix sum dengan formula `preSum[i] = energy[i] + preSum[i-k]`
- Cari maximum value dari array prefix sum

--- 

### Kompleksitas 
- **Kompleksitas Waktu:** O(N) 
- **Memori Tambahan:** O(N) atau O(1) jika kamu menggunakan array `energy` untuk menyimpan prefix sumnya
--- 

## [English] - [3147. Taking Maximum Energy From the Mystic Dungeon](https://leetcode.com/problems/taking-maximum-energy-from-the-mystic-dungeon/description/?envType=daily-question&envId=2025-10-10)

### Description 

In a mystic dungeon, n magicians are standing in a line. Each magician has an attribute that gives you energy. Some magicians can give you negative energy, which means taking energy from you.

You have been cursed in such a way that after absorbing energy from magician i, you will be instantly transported to magician (i + k). This process will be repeated until you reach the magician where (i + k) does not exist.

In other words, you will choose a starting point and then teleport with k jumps until you reach the end of the magicians' sequence, absorbing all the energy during the journey.

You are given an array energy and an integer k. Return the maximum possible energy you can gain.

Note that when you are reach a magician, you must take energy from them, whether it is negative or positive energy.

--- 

### Intuition 

- The key concept to solve this problem was prefix sum
- Because of the teleportation will do with `k` step, so the formula of prefSum is `preSum[i] = energy[i] + preSum[i]`

--- 

### Approach 

- Iterate array `energy` to get the prefix sum array with the formula `preSum[i] = energy[i] + preSum[i]`
- Find the maximum value from the prefix sum array

--- 

### Complexity 
- **Time Complexity:** O(N) 
- **Auxiliary Space:** O(N) or O(1) if you rewrite the array energy it self
---