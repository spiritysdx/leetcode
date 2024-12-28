// https://leetcode.cn/problems/search-in-a-binary-search-tree/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
    var dfs func(*TreeNode)
    var ans *TreeNode
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        if val == node.Val {
            ans = node
            return
        }
        if val < node.Val {
            dfs(node.Left)
        }
        if val > node.Val {
            dfs(node.Right)
        }
    }
    dfs(root)
    return ans
}
