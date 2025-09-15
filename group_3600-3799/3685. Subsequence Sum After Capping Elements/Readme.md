# 🇮🇩 README (Bahasa Indonesia)

## 📌 Deskripsi Masalah

Diberikan array `nums` berukuran `n` dan sebuah bilangan bulat positif `k`.
Untuk setiap nilai `x` dari 1 hingga `n`, bentuk **array capped** dengan aturan:

```
capped[i] = min(nums[i], x)
```

Kemudian tentukan apakah ada subsequence dari `capped` yang jumlah elemennya tepat `k`.

Hasil berupa array boolean berukuran `n`, di mana `ans[x-1] = true` jika mungkin dengan cap = `x`, dan `false` jika tidak.

---

## 💡 Ide & Pendekatan

1. **Sort array** supaya mudah memproses elemen secara bertahap.
2. Iterasi cap `x` dari 1..n:

   * Untuk setiap `nums[i] < x`, masukkan ke dalam **DP subset sum** (knapsack klasik).
   * Elemen yang **≥ x** dianggap bernilai `x` (capped).
   * Sisa elemen ≥ x bisa digunakan beberapa kali, sehingga kita hanya perlu cek apakah target `k` bisa dicapai dengan menambahkan `t * x` (`0 ≤ t ≤ jumlah elemen ≥ x`) ke salah satu sum yang sudah mungkin dari DP.
3. Jika ada `dp[k - t*x] = true` → berarti `ans[x-1] = true`.

---

## 🧩 Pseudocode

```
sort(nums)
dp[0] = true
i = 0

for x in 1..n:
    # Tambahkan semua elemen < x ke DP
    while i < n and nums[i] < x:
        for s from k down to nums[i]:
            dp[s] = dp[s] or dp[s - nums[i]]
        i++

    # Hitung jumlah elemen ≥ x
    count = n - i

    # Cek apakah k bisa dicapai dengan dp + kelipatan x
    ans[x-1] = false
    for t in 0..count:
        if k - t*x >= 0 and dp[k - t*x]:
            ans[x-1] = true
            break
```

---

## ⏱️ Kompleksitas

* **Sorting**: `O(n log n)`
* **DP update**: setiap elemen `nums[i]` diproses sekali, biaya total `O(n * k)`
* **Cek kelipatan `x`**: per cap maksimal `n`, jadi `O(n^2)` dalam kasus terburuk.
* **Total**: `O(n * k + n^2)` → cukup efisien untuk `n, k ≤ 4000`.

---

# 🇬🇧 README (English)

## 📌 Problem Description

You are given an array `nums` of size `n` and an integer `k`.
For each `x` in the range `[1..n]`, construct a **capped array**:

```
capped[i] = min(nums[i], x)
```

Then check if there exists a subsequence of `capped` whose sum is exactly `k`.

Return a boolean array `ans` of length `n`, where `ans[x-1] = true` if possible for cap `x`, otherwise `false`.

---

## 💡 Idea & Approach

1. **Sort nums** to process elements incrementally.
2. Iterate cap `x = 1..n`:

   * Add all `nums[i] < x` into the **subset sum DP** (classic 0/1 knapsack).
   * Remaining elements (`nums[i] ≥ x`) are capped to `x`.
   * Let `count = number of such elements`.
     Check whether `k` can be formed as `dp[s] + t*x` for some `t ≤ count`.
3. If any such combination exists, mark `ans[x-1] = true`.

---

## 🧩 Pseudocode

```
sort(nums)
dp[0] = true
i = 0

for x in 1..n:
    # Add all nums[i] < x into dp
    while i < n and nums[i] < x:
        for s from k down to nums[i]:
            dp[s] = dp[s] or dp[s - nums[i]]
        i++

    count = n - i  # number of elements >= x

    ans[x-1] = false
    for t in 0..count:
        if k - t*x >= 0 and dp[k - t*x]:
            ans[x-1] = true
            break
```

---

## ⏱️ Complexity

* **Sorting**: `O(n log n)`
* **DP updates**: `O(n * k)`
* **Check multiples per cap**: `O(n^2)` worst case.
* **Total**: `O(n * k + n^2)` → feasible for `n, k ≤ 4000`.