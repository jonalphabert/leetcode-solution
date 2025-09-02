## Problem

**Indonesia:**  
Diberikan sebuah array `nums` berisi n bilangan bulat.  
Kembalikan semua quadruplet unik `[nums[a], nums[b], nums[c], nums[d]]` sehingga:
- 0 <= a, b, c, d < n
- a, b, c, dan d semuanya berbeda
- nums[a] + nums[b] + nums[c] + nums[d] == target

Jawaban dapat dikembalikan dalam urutan apa pun.

**English:**  
Given an array `nums` of n integers, return an array of all the unique quadruplets `[nums[a], nums[b], nums[c], nums[d]]` such that:
- 0 <= a, b, c, d < n
- a, b, c, and d are distinct
- nums[a] + nums[b] + nums[c] + nums[d] == target

You may return the answer in any order.

---

## Solution

### Bahasa Indonesia

- Urutkan array `nums` terlebih dahulu.
- Gunakan dua loop untuk memilih dua angka pertama.
- Gunakan dua pointer (`left` dan `right`) untuk mencari dua angka berikutnya sehingga jumlah keempat angka sama dengan `target`.
- Untuk menghindari quadruplet duplikat, gunakan set atau skip angka yang sama.
- Tambahkan quadruplet yang valid ke hasil.

### English

- Sort the `nums` array first.
- Use two loops to pick the first two numbers.
- Use two pointers (`left` and `right`) to find the next two numbers so that the sum of all four equals `target`.
- To avoid duplicate quadruplets, use a set or skip duplicate numbers.
- Add valid quadruplets to the result.

**Time Complexity:** O(n³)
**Space Complexity:** O(1) or O(k) depending on the storage of results.