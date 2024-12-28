// https://leetcode.cn/problems/maximum-sum-bst-in-binary-tree/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxSumBST(root *TreeNode) (ans int) {
    var dfs func(*TreeNode) (int, int, int)
    dfs = func(node *TreeNode) (int, int, int) {
        if node == nil {
            return math.MaxInt, math.MinInt, 0
        }

        lMin, lMax, lSum := dfs(node.Left) // 递归左子树
        rMin, rMax, rSum := dfs(node.Right) // 递归右子树
        x := node.Val
        if x <= lMax || x >= rMin { // 不是二叉搜索树
            return math.MinInt, math.MaxInt, 0
        }

        s := lSum + rSum + x // 这棵子树的所有节点值之和
        ans = max(ans, s)

        return min(lMin, x), max(rMax, x), s
    }
    dfs(root)
    return
}

func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
