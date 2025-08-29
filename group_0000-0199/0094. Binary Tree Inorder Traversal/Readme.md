# Indonesia
## Problem

Diberikan root dari sebuah binary tree, kembalikan urutan nilai node berdasarkan **inorder traversal**.

**Inorder traversal** berarti mengunjungi node kiri terlebih dahulu, lalu node saat ini, kemudian node kanan.

**Contoh:**
```
Input: root = [1,null,2,3]
Output: [1,3,2]
Penjelasan: Urutan kunjungan adalah kiri → root → kanan.
```

## Solution

Terdapat dua pendekatan untuk melakukan inorder traversal:

### 1. Rekursif

- Kunjungi node kiri secara rekursif.
- Tambahkan nilai node saat ini.
- Kunjungi node kanan secara rekursif.

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Memori:** O(n) (karena stack rekursi)

### 2. Iteratif (Menggunakan Stack)

- Gunakan stack untuk menelusuri node kiri hingga habis.
- Pop node dari stack, tambahkan nilainya ke hasil.
- Geser ke node kanan dan ulangi proses.

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Memori:** O(n)

**Contoh Implementasi (Go):**
```go
// Recursive Solution
func inorderTraversal(root *TreeNode) []int {
    res := []int{}
    if root == nil {
        return res
    }
    res = append(res, inorderTraversal(root.Left)...)
    res = append(res, root.Val)
    res = append(res, inorderTraversal(root.Right)...)
    return res
}

// Iterative Solution
func inorderTraversal(root *TreeNode) []int {
    res := []int{}
    stack := []*TreeNode{}
    curr := root
    for curr != nil || len(stack) > 0 {
        for curr != nil {
            stack = append(stack, curr)
            curr = curr.Left
        }
        n := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        res = append(res, n.Val)
        curr = n.Right
    }
    return res
}
```

---

# English
## Problem

Given the root of a binary tree, return the inorder traversal of its nodes' values.

**Inorder traversal** means visiting the left node first, then the current node, then the right node.

**Example:**
```
Input: root = [1,null,2,3]
Output: [1,3,2]
Explanation: The visiting order is left → root → right.
```

## Solution

There are two approaches to perform inorder traversal:

### 1. Recursive

- Recursively visit the left node.
- Add the current node's value.
- Recursively visit the right node.

**Time Complexity:** O(n)  
**Space Complexity:** O(n) (due to recursion stack)

### 2. Iterative (Using Stack)

- Use a stack to traverse left nodes until none remain.
- Pop nodes from the stack, add their values to the result.
- Move to the right node and repeat.

**Time Complexity:** O(n)  
**Space Complexity:** O(n)

**Sample Implementation (Go):**
```go
// Recursive Solution
func inorderTraversal(root *TreeNode) []int {
    res := []int{}
    if root == nil {
        return res
    }
    res = append(res, inorderTraversal(root.Left)...)
    res = append(res, root.Val)
    res = append(res, inorderTraversal(root.Right)...)
    return res
}

// Iterative Solution
func inorderTraversal(root *TreeNode) []int {
    res := []int{}
    stack := []*TreeNode{}
    curr := root
    for curr != nil || len(stack) > 0 {
        for curr != nil {
            stack = append(stack, curr)
            curr = curr.Left
        }
        n := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        res = append(res, n.Val)
        curr = n.Right
    }
    return res