# Indonesia
## Problem

Diberikan sebuah bilangan bulat `n`, kembalikan jumlah struktur BST (Binary Search Tree) yang unik yang dapat dibentuk dengan tepat `n` node bernilai unik dari 1 hingga `n`.

**Contoh:**
```
Input: n = 3
Output: 5
Penjelasan: Terdapat 5 struktur BST yang berbeda dengan 3 node unik.
```

## Solution

Masalah ini dapat diselesaikan menggunakan **rekursi dengan memoization** (Dynamic Programming).  
Jumlah BST unik untuk `n` node adalah jumlah kemungkinan memilih setiap node sebagai root, lalu mengalikan jumlah BST di kiri dan kanan root.

Rumus rekursif:
- Untuk setiap `i` dari 0 hingga `n-1`, jumlah BST adalah `numTrees(i) * numTrees(n-1-i)`.
- Gunakan memoization untuk menghindari perhitungan berulang.

**Kompleksitas Waktu:** O(n²)  
**Kompleksitas Memori:** O(n)

**Contoh Implementasi (Go):**
```go
var memo = map[int]int{0: 1}

func numTrees(n int) int {
    if val, ok := memo[n]; ok {
        return val
    }
    var res int = 0
    for i := 0; i < n; i++ {
        res += numTrees(i) * numTrees(n-1-i)
    }
    memo[n] = res
    return res
}
```

---

# English
## Problem

Given an integer `n`, return the number of structurally unique BST's (binary search trees) that store values 1 to `n`.

**Example:**
```
Input: n = 3
Output: 5
Explanation: There are 5 structurally unique BSTs with 3 unique nodes.
```

## Solution

This problem can be solved using **recursion with memoization** (Dynamic Programming).  
The number of unique BSTs for `n` nodes is the sum of all possible choices for the root, multiplied by the number of BSTs in the left and right subtrees.

Recursive formula:
- For each `i` from 0 to `n-1`, the number of BSTs is `numTrees(i) * numTrees(n-1-i)`.
- Use memoization to avoid redundant calculations.

**Time Complexity:** O(n²)  
**Space Complexity:** O(n)

**Sample Implementation (Go):**
```go
var memo = map[int]int{0: 1}

func numTrees(n int) int {
    if val, ok := memo[n]; ok {
        return val
    }
    var res int = 0
    for i := 0; i < n; i++ {
        res += numTrees(i) * numTrees(n-1-i)
    }
    memo[n] =