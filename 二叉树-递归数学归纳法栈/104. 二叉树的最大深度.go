// https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil { // 边界条件
        return 0
    }
    L := maxDepth(root.Left) // 递
    R := maxDepth(root.Right)
    return max(L, R)+1 // 归
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}
