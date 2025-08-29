/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  Recursive Solution
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
        // 1. Traversal ke kiri terus
        for curr != nil {
            stack = append(stack, curr)
            curr = curr.Left
        }

        // 2. Pop dari stack
        n := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        // 3. Tambahkan ke hasil
        res = append(res, n.Val)

        // 4. Geser ke kanan
        curr = n.Right
    }

    return res
}