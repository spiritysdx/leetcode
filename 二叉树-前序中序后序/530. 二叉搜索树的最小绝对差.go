// https://leetcode.cn/problems/minimum-absolute-difference-in-bst/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 中序遍历
func getMinimumDifference(root *TreeNode) int {
    var dfs func(*TreeNode)
    var pre *int // 使用指针以便处理初始值
    ans := math.MaxInt
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        // 中序遍历：先左子树
        dfs(node.Left)
        // 计算当前节点与前一个节点的差值
        if pre != nil {
            ans = min(ans, abs(node.Val-*pre))
        }
        // 更新 pre 为当前节点值
        pre = &node.Val
        // 中序遍历：后右子树
        dfs(node.Right)
    }
    dfs(root)
    return ans
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
