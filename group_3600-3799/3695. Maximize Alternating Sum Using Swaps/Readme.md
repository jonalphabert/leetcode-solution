# 📖 Readme

## [Bahasa Indonesia] - [3695. Maximize Alternating Sum Using Swaps](https://leetcode.com/problems/maximize-alternating-sum-using-swaps/description/)

### Deskripsi 

Diberikan array nums dan daftar swaps (pair index yang bisa ditukar berulang kali). Cari alternating sum maksimum

```
    nums[0] - nums[1] + nums[2] - nums[3] + ...
```
--- 

### Intuisi 

- Dari array `swaps` kita dapat membuat struktur data graph yang dimana ini nanti memunculkan semua connected graph.
- Dari connected graph ini, maka semua element di dalamnya dapat ditukar. - Karena kita akan memaksimalkan nilai alternating sum, maka kita perlu menaruh bilangan terbesar pada connected graph dipindahkan ke index genap.

--- 

### Pendekatan 

- Buat `graph` dengan pendekatan adj list.
- Generate `graph` dengan melakukan iterasi semua element di swaps. Untuk graph bersifat undirect graph jadi a->b = b->a
- Inisiasi map `visited` untuk menentukan bahwa node/index tersebut telah dikunjungi ketika dilakukan dfs.
- Iterasi `i` dari 0 hingga `n`
    - Jika ditemukan `visited[i] == true` maka angka di index `i` telah dijumlahkan. Jadi lanjut ke iterasi selanjutnya
    - Jika tidak ditemukan maka lakukan 
        - `dfs(i, comp)` dimana `comp` adalah berisi node/index yang dilalui dfs
        - Setelah dfs selesai, maka masukkan nilai `nums` dari index yang dilalui oleh proses dfs ke array `vals`, dan hitung ada berapa index genap yang dilalui dalam variable `countEven`
        - Urutkan `vals` dengan urutan dari besar ke kecil
        - Iterasi `vals` 
            - Jika `index < countEven` maka lakukan komputasi `total = total + val`
            - Jika tidak maka lakukan komputasi `total = total - val`
- Kembalikan `total`

--- 

### Kompleksitas 
- **Waktu:** O(N log N + E)
- **Memori:** O(N + E) (karena adjacency list).
--- 

Oke 👍 aku tuliskan versi bahasa Inggrisnya yang clean, singkat, dan rapi untuk README:

---

## [English] - 📖 [3695. Maximize Alternating Sum Using Swaps](https://leetcode.com/problems/maximize-alternating-sum-using-swaps/description/)

### Description

You are given an integer array `nums` and a list of `swaps` where each element `[p, q]` represents indices that can be swapped any number of times.

The task is to maximize the **alternating sum** defined as:

```
nums[0] - nums[1] + nums[2] - nums[3] + ...
```

---

### Intuition

* The `swaps` define an **undirected graph** between indices.
* Within each connected component, indices can be freely rearranged.
* To maximize the alternating sum, the **largest values should be placed on even indices**, and the smaller ones on odd indices.

---

### Approach

1. Build an undirected graph from the `swaps`.
2. Use DFS/BFS to find each connected component.
3. Collect the values of `nums` within the component.
4. Count how many indices in the component are even/odd.
5. Sort the values in descending order.
6. Assign the top `countEven` values to even positions (add to total), and the rest to odd positions (subtract from total).

---

### Complexity

* **Time:** `O(N log N + E)` → DFS over the graph plus sorting values in each component.
* **Space:** `O(N + E)` → adjacency list + visited set.

---