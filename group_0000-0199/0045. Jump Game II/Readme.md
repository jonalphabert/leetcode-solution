## Problem

**Indonesia:**  
Diberikan sebuah array integer `nums` dengan panjang n dan diindeks mulai dari 0.  
Awalnya kamu berada di indeks 0.  
Setiap elemen `nums[i]` menunjukkan panjang maksimum lompatan maju dari indeks i.  
Artinya, jika kamu berada di indeks i, kamu bisa melompat ke indeks mana pun `(i + j)` dengan:
- `0 <= j <= nums[i]` dan
- `i + j < n`

Kembalikan jumlah minimum lompatan untuk mencapai indeks `n - 1`.  
Test case dijamin selalu bisa mencapai indeks terakhir.

**English:**  
You are given a 0-indexed array of integers `nums` of length n.  
You are initially positioned at index 0.  
Each element `nums[i]` represents the maximum length of a forward jump from index i.  
In other words, if you are at index i, you can jump to any index `(i + j)` where:
- `0 <= j <= nums[i]` and
- `i + j < n`

Return the minimum number of jumps to reach index `n - 1`.  
The test cases are generated such that you can reach index `n - 1`.

---

## Solution

### Bahasa Indonesia

#### 1. BFS + DP (O(n^2))

- Buat array DP untuk menyimpan minimum lompatan ke setiap indeks.
- Iterasi setiap indeks, dan untuk setiap kemungkinan lompatan dari indeks tersebut, update DP untuk indeks tujuan.
- Jika sudah mencapai indeks terakhir, berhenti.
- Kompleksitas waktu: O(n^2)
- Kompleksitas memori: O(n)

#### 2. Greedy (Optimized, O(n))

- Lacak indeks terjauh yang bisa dicapai (`farthest`) dan batas akhir lompatan saat ini (`end`).
- Setiap kali mencapai batas akhir lompatan, tambahkan jumlah lompatan dan perbarui batas ke indeks terjauh.
- Ulangi hingga mencapai indeks terakhir.
- Kompleksitas waktu: O(n)
- Kompleksitas memori: O(1)

---

### English

#### 1. BFS + DP (O(n^2))

- Create a DP array to store the minimum jumps to reach each index.
- For each index, iterate through all possible jumps from that index and update the DP value for the destination index.
- Stop if the last index is reached.
- Time complexity: O(n^2)
- Space complexity: O(n)

#### 2. Greedy (Optimized, O(n))

- Track the farthest index that can be reached (`farthest`) and the end of the current jump (`end`).
- Each time you reach the end of the current jump, increment the jump count and update the end to the farthest index.
- Repeat until the last index is reached.
- Time complexity: O(n)
- Space complexity: O(1)

---

**Pseudocode Greedy:**
```
Initialize farthest = 0, jumps = 0, end = 0
For i from 0 to n-2:
    farthest = max(farthest, i + nums[i])
    If i == end:
        jumps++
        end = farthest
Return jumps
```

**Pseudocode BFS + DP:**
```
Initialize dp[0] = 0, dp[1..n-1] = INF
For i from 0 to n-1:
    For j from 1 to nums[i]:
        If i + j < n:
            dp[i + j] = min(dp[i + j], dp[i] + 1)
Return dp[n-1]
```