// https://leetcode.cn/problems/range-sum-of-bst/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    var dfs func(*TreeNode)
    ans := 0
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        if node.Val < low {
            dfs(node.Right)
        } else if node.Val > high {
            dfs(node.Left)
        } else if node.Val >= low && node.Val <= high {
            ans += node.Val
            dfs(node.Right)
            dfs(node.Left)
        }
    }
    dfs(root)
    return ans
}
