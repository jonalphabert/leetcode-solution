/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    maxSum := math.MinInt

    var bfs func(node *TreeNode) int 
    bfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

		// Cari penjumlahan terbesar di kiri dan dikanan
        leftMaxGain := max(bfs(node.Left),0)
        rightMaxGain := max(bfs(node.Right),0)

		// Ada kemungkinan penjumlahan root ke kiri dan kanan ini bakal menjadi yang terbesar
        maxSum = max(maxSum, node.Val + leftMaxGain + rightMaxGain)

		// Ini kenapa memilih max(leftMaxGain, rightMaxGain) karena hanya bisa 1 kali traverse untuk suatu node dilewati
        return node.Val + max(leftMaxGain, rightMaxGain)
    } 

    bfs(root)

    return maxSum
}