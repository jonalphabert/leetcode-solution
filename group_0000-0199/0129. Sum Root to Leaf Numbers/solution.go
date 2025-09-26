/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    var dfs func(current int, node *TreeNode) int
    dfs = func(current int, node *TreeNode) int{
        if node == nil {
            return 0
        }
        currentUpdated := current*10+node.Val
        
        if node.Left == nil && node.Right == nil {
            return currentUpdated
        }

        return dfs(currentUpdated, node.Left) + dfs(currentUpdated, node.Right)
    }

    return dfs(0, root)
}