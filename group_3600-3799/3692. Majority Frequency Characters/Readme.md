# 📖 Readme

## [Bahasa Indonesia] - [3692. Majority Frequency Characters](https://leetcode.com/problems/majority-frequency-characters/description/)

### Deskripsi 

Diberikan sebuah string `s` yang terdiri dari lowercase English letters.

Kelompok frekuensi untuk nilai `k` adalah himpunan karakter yang muncul tepat `k` kali dalam `s`.

Kelompok frekuensi mayoritas adalah kelompok frekuensi yang berisi jumlah karakter berbeda terbanyak.

Kembalikan sebuah string yang berisi semua karakter dalam kelompok frekuensi mayoritas, dalam urutan apa pun. Jika dua atau lebih kelompok frekuensi memiliki ukuran terbesar yang sama, pilih kelompok dengan frekuensi `k` yang lebih besar.

--- 

### Intuisi 

- Hitung frekuensi frekuensi tiap karakternya
- Lalu dari perhitungan frekuensi, kita kelompokkan karakter dengan frekuensi yang sama banyaknya
- Dari pengelompokan frekuensi, maka kita lakukan iterasi untuk mencari group mana yang lebih banyak membernya. Kalau membernya sama, maka kita ambil frekuensi yang lebih besar

--- 

### Pendekatan 

- Definisikan `freq` dengan struktur data map
- Iterasi setiap karakter yang ada di string `s` untuk update nilai `freq` dari karakter tersebut
- Definisikan `group` dengan struktur data map
- Iterasi setiap key yang terbentuk dari `freq`. Kelompokkan berdasarkan banyaknya kemunculan sebuah karakter di string `s`
- Definisikan result, maxSizeGroup, dan maxFreq
- Iterasi setiap freq, value `group`
    - Jika ditemukan `len(value) > maxSizeGroup` atau `len(value) == maxSizeGroup && maxFreq < freq` maka update maxSizeGroup, result, dan maxFreq.
- Return result

--- 

### Kompleksitas 
- **Waktu:** O(N) 
- **Memori:** O(N) + catatan optimasi 
--- 


## [English] - [3692. Majority Frequency Characters](https://leetcode.com/problems/majority-frequency-characters/description/)

### Description 

You are given a string s consisting of lowercase English letters.

The frequency group for a value k is the set of characters that appear exactly k times in s.

The majority frequency group is the frequency group that contains the largest number of distinct characters.

Return a string containing all characters in the majority frequency group, in any order. If two or more frequency groups tie for that largest size, pick the group whose frequency k is larger.

--- 

### Intuition 

- Count frequency of each character
- Then from the frequency count, we group characters with the same frequency
- From the frequency grouping, we iterate to find which group has the most members. If the members are the same, we take the higher frequency

--- 

### Approach 

- Define `freq` as a map structure
- Iterate each character in the string `s` to update the `freq` value of that character
- Define `group` as a map structure
- Iterate each key formed from `freq`. Group based on how many times a character appears in the string `s`
- Define result, maxSizeGroup, and maxFreq
- Iterate each freq, value `group`
    - If found `len(value) > maxSizeGroup` or `len(value) == maxSizeGroup && maxFreq < freq` then update maxSizeGroup, result, and maxFreq.
- Return result

--- 

### Complexity 
- **Time Complexity:** O(N) 
- **Auxiliary Space:** O(N)
--- 