// https://leetcode.cn/problems/minimum-depth-of-binary-tree/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    var dfs func(*TreeNode, int)
    ans := -1
    dfs = func(node *TreeNode, depth int) {
        if node != nil && node.Left == nil && node.Right == nil {
            depth++
            if ans == -1 {
                ans = depth
            } else {
                ans = min(ans, depth)
            }
            return
        }
        depth++
        if node.Left != nil {
            dfs(node.Left, depth)
        }
        if node.Right != nil {
            dfs(node.Right, depth)
        }
    }
    dfs(root, 0)
    return ans
}

func min(a,b int) int {
    if a > b {
        return b
    }
    return a
}
