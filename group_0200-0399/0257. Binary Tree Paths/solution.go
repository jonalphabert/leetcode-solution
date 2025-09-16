/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    if root == nil {
        return []string{}
    }

    var res []string

    var dfs func(node *TreeNode, path string)
    dfs = func(node *TreeNode, path string) {
        if node.Left == nil && node.Right == nil {
            res = append(res, path+strconv.Itoa(node.Val))
            return
        }

        if node.Left != nil {
            dfs(node.Left, path+strconv.Itoa(node.Val)+"->")
        }
        if node.Right != nil {
            dfs(node.Right, path+strconv.Itoa(node.Val)+"->")
        }
    }

    dfs(root, "")
    return res
}
