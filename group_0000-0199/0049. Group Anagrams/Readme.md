# 📖 Readme

## [Bahasa Indonesia] - [49. Group Anagrams](https://leetcode.com/problems/group-anagrams/description/)

### Deskripsi 

Diberikan sebuah array string `strs`, tugasmu adalah melakukan grouping anagram yang ada. Kamu bisa mengembalikan return dalam urutan apapun.

Note
```
Anagram adalah kata atau frasa yang huruf-hurufnya disusun ulang untuk membentuk kata atau frasa baru, namun tetap menggunakan huruf yang sama persis dan dalam jumlah yang sama.
```

--- 

### Intuisi 

- Masalah ini dapat diselesaikan dengan menggunakan teknik sorting dan struktur data hash map.
- Sorting disini digunakan untuk mengurutkan semua penyusun sebuah string sehingga jika terdapat anagram yang sama maka 2 string ini kalau kita sort akan menghasilkan string yang sama
- Hash map disini digunakan untuk menyimpan semua string yang memiliki anagram. Untuk key nya bakal menggunakan hasil sorting.

--- 

### Pendekatan 

- Definisikan struktur data map
- Lakukan iterasi pada setiap elemen dalam array `strs`.
- Untuk setiap elemen, lakukan sorting terhadap karakter di dalam string.
- Hasil sorting akan digunakan `key` pada struktur data map. Dan masukkan element tersebut dalam `key` tersebut

--- 

### Kompleksitas 
- **Waktu:** O(N * s log s), dengan N = jumlah string, s = panjang rata-rata string.
- **Memori:** O(N)
--- 

### Pseudocode

```
func solution(strs):
    mapAnagram = map[string]string[]
    for idx, str = range strs:
        sorted = sort(str)
        mapAnagramp[sorted] = append(s)
    
    res = string[][]

    for key = range mapAnagram:
        res = append(mapAnagram[key])

    return res
```

## [English] - [49. Group Anagrams](https://leetcode.com/problems/group-anagrams/description/)

### Description 

Given an array of strings strs, group the anagrams together. You can return the answer in any order.

--- 

### Intuition 

- This problem can be solved using sorting and map data structure
- We use sorting to rearrange the characters of each string. Strings that are anagrams will produce the same sorted result
- We use a map to store groups of strings that share the same sorted key.

--- 

### Approach 

- Define map data structure
- Loop for every element `s` on `strs`
- Sort every element `s`
- Add element `s` on map with key of sorted version of `s`

--- 

### Complexity 
- **Time Complexity:** O(N * s log s) => Note: s log(s) used for sorting each string on `strs` 
- **Memory Complexity:** O(N)
--- 

### Pseudocode

```
func solution(strs):
    mapAnagram = map[string]string[]
    for idx, str = range strs:
        sorted = sort(str)
        mapAnagramp[sorted] = append(s)
    
    res = string[][]

    for key = range mapAnagram:
        res = append(mapAnagram[key])

    return res
```