## Problem

**Indonesia:**  
Diberikan sebuah array integer `order` dengan panjang n dan sebuah array integer `friends`.

- `order` berisi setiap bilangan dari 1 sampai n tepat satu kali, merepresentasikan ID peserta lomba sesuai urutan mereka finish.
- `friends` berisi ID teman-temanmu yang ikut lomba, sudah terurut menaik. Setiap ID di `friends` pasti ada di `order`.

Kembalikan array berisi ID teman-temanmu sesuai urutan mereka finish.

**English:**  
You are given an integer array `order` of length n and an integer array `friends`.

- `order` contains every integer from 1 to n exactly once, representing the IDs of the participants of a race in their finishing order.
- `friends` contains the IDs of your friends in the race sorted in strictly increasing order. Each ID in `friends` is guaranteed to appear in the `order` array.

Return an array containing your friends' IDs in their finishing order.

---

## Solution

### Bahasa Indonesia

- Buat map dari ID ke indeks posisi finish berdasarkan array `order`.
- Urutkan array `friends` berdasarkan posisi finish menggunakan map tersebut.
- Kembalikan array hasil pengurutan.

**Kompleksitas Waktu:** O(n + m log m), n = panjang order, m = panjang friends  
**Kompleksitas Memori:** O(n)

### English

- Build a map from ID to finishing position index using the `order` array.
- Sort the `friends` array based on their finishing position using the map.
- Return the sorted array.

**Time Complexity:** O(n + m log m), n = length of order, m = length of friends  
**Space Complexity:** O(n)