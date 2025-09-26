# 📖 Readme

## [ID] Bahasa Indonesia - 129. Sum Root to Leaf Numbers

### Deskripsi 

Diberikan sebuah `root` dari sebuah binary tree yang terdiri dari digit `0` sampai `9`. Setiap angka dari `root` ke `leaf` merepresentasikan sebuah angka.

Misalkan ada tree yang berbentuk seperti `1->2->3` merepresentasikan angka `123`

Tugasmu adalah mengembalikan semua total penjumlahan dari semua bilangan yang dibentuk dari root ke leafnya.

--- 

### Intuisi 

- Soal ini dapat diselesaikan dengan metode DFS(Depth First Search) traverse.
- Ketika melakukan traverse ke leaf, perlu membawa `currentNumber` yang merepresentasikan bilangan yang dibentuk sampai leaf tersebut.
- Ketika kita menemukan `leaf` yang berada di kiri dan kanan adalah `null` maka kita kembalikan nilai `currentNumber` selain itu kita kembalikan penjumlahan nilai DFS ke arah kiri dan kanan

--- 

### Pendekatan 

- Lakukan DFS dari root
- Check apakah node adalah `null`.
    - Jika iya, kembalikan 0 (ini base case)
- Update `currentNumber` dengan rumus `currentNumber * 10 + Node.Val`
- Lakukan traverse ke kiri dan kanan => `dfs(left)` dan `dfs(right)`
- Cek apakah hasil traverse keduanya bernilai `0`
    - Jika iya, kembalikan `currentNumber`
    - Jika tidak, kembalikan `dfs(left) + dfs(right)`

--- 

### Kompleksitas 
- **Waktu:** O(n) 
- **Memori:** O(n) 
--- 

### Contoh 

- Input nilai parameter 

```
Input: root = [1,2,3]
```

- Output yang diharapkan
```
Output: 25

Penjelasan:
    Perjalanan root-to-left 1->2 merepresentasikan nilai 12.
    Perjalanan root-to-leaf path 1->3 merepresentasikan nilai 13.
    Sehingga, sum = 12 + 13 = 25.
```

### Psuedocode

```bash
def sumNumbers(root):
    def dfs(node, current):
        if not node:
            return 0
        current = current * 10 + node.val
        if not node.left and not node.right:
            return current
        return dfs(node.left, current) + dfs(node.right, current)
    return dfs(root, 0)
```

## [EN] English - 129. Sum Root to Leaf Numbers

### Problem Description 

You are given the root of a binary tree containing digits from 0 to 9.  
Each root-to-leaf path in the tree represents a number.  

For example, the root-to-leaf path `1 -> 2 -> 3` represents the number `123`.  
The task is to return the total sum of all root-to-leaf numbers.  
A leaf node is a node with no children.  

---

### Intuition 

- We can solve this problem using a Depth-First Search (DFS) traversal.  
- While traversing the tree, we carry a `currentNumber` that represents the number formed so far.  
- When we reach a leaf node, we return the `currentNumber`.  
- Otherwise, we return the sum of the recursive calls on the left and right subtrees.  

---

### Approach 

1. Start DFS from the root with `currentNumber = 0`.  
2. If the current node is `null`, return 0 (base case).  
3. Update `currentNumber` using the formula:
    ```
    currentNumber = currentNumber * 10 + node.val
    ```
4. If the node is a leaf (both left and right are null), return `currentNumber`.  
5. Otherwise, return the sum of the DFS results from the left and right children.  

---

### Complexity 
- **Time Complexity:** O(n), since we visit each node exactly once.  
- **Space Complexity:** O(n), due to the recursion stack in the worst case (a skewed tree).  

### Example 

- Input

```
Input: root = [1,2,3]
```

- Output yang diharapkan
```
Output: 25

Explanation:
    The root-to-leaf path 1->2 represents the number 12.
    The root-to-leaf path 1->3 represents the number 13.
    Therefore, sum = 12 + 13 = 25.
```

### Psuedocode

```bash
def sumNumbers(root):
    def dfs(node, current):
        if not node:
            return 0
        current = current * 10 + node.val
        if not node.left and not node.right:
            return current
        return dfs(node.left, current) + dfs(node.right, current)
    return dfs(root, 0)
```

