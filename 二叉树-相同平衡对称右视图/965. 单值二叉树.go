// https://leetcode.cn/problems/univalued-binary-tree/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
    ans := true
    target := root.Val
    var dfs func(*TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        if node.Val != target {
            ans = false
        }
        dfs(node.Left)
        dfs(node.Right)
    }
    dfs(root)
    return ans
}
