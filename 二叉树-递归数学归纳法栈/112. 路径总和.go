// https://leetcode.cn/problems/path-sum/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    // 自顶向下
    var ans bool
    var dfs func(*TreeNode, int)
    dfs = func(node *TreeNode, cal int) {
        if node == nil {
            return
        }
        cal += node.Val
        if node.Left != nil {
            dfs(node.Left, cal)
        }
        if node.Right != nil {
            dfs(node.Right, cal)
        }
        if node.Left == nil && node.Right == nil && cal == targetSum {
            ans = true
        }
    }
    dfs(root, 0)
    return ans
}

func hasPathSum(root *TreeNode, targetSum int) bool {
    // 自底向上
    if root == nil {
        return false
    }
    targetSum -= root.Val
    if root.Left == root.Right {
        return targetSum == 0
    }
    return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}
