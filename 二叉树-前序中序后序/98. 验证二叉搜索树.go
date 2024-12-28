// https://leetcode.cn/problems/validate-binary-search-tree/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int           // 节点的值
 *     Left *TreeNode    // 左子节点
 *     Right *TreeNode   // 右子节点
 * }
 */

// 前序遍历
func dfs(node *TreeNode, left, right int) bool {
    if node == nil { 
        return true // 如果节点为空，返回 true（空树是有效的二叉搜索树）
    }
    x := node.Val 
    // 检查当前节点的值是否在有效范围 (left, right) 内
    return left < x && x < right && 
        dfs(node.Left, left, x) &&   // 递归检查左子树，更新右边界为当前节点值
        dfs(node.Right, x, right)   // 递归检查右子树，更新左边界为当前节点值
}

func isValidBST(root *TreeNode) bool {
    // 检查以 root 为根的树是否是有效的二叉搜索树
    return dfs(root, math.MinInt, math.MaxInt) // 初始化有效范围为整个整数范围

// 中序遍历
func isValidBST(root *TreeNode) bool {
    pre := math.MinInt // 用于记录中序遍历中前一个节点的值，初始化为最小整数
    var dfs func(*TreeNode) bool // 定义一个递归函数 dfs
    dfs = func(node *TreeNode) bool {
        if node == nil { 
            return true // 如果节点为空，返回 true（空树是有效的二叉搜索树）
        }
        // 递归检查左子树是否是有效的二叉搜索树
        if !dfs(node.Left) || node.Val <= pre { 
            return false // 如果左子树无效，或者当前节点值小于等于前一个节点值，返回 false
        }
        pre = node.Val // 更新 pre 为当前节点的值
        return dfs(node.Right) // 递归检查右子树是否是有效的二叉搜索树
    }
    return dfs(root) // 调用递归函数检查整个树
}

// 后序遍历
func dfs(node *TreeNode) (int, int) {
    if node == nil {
        // 如果节点为空，返回最大整数作为最小值，最小整数作为最大值（空子树的边界）
        return math.MaxInt, math.MinInt
    }
    // 递归计算左子树的最小值和最大值
    lMin, lMax := dfs(node.Left)
    // 递归计算右子树的最小值和最大值
    rMin, rMax := dfs(node.Right)
    x := node.Val // 当前节点的值
    if x <= lMax || x >= rMin { 
        // 如果当前节点值小于等于左子树的最大值，或者大于等于右子树的最小值，说明不是二叉搜索树
        // 返回一个无效范围，用最小整数和最大整数表示
        return math.MinInt, math.MaxInt
    }
    // 返回当前子树的最小值和最大值
    // 当前子树的最小值是左子树最小值和当前节点值的较小者
    // 当前子树的最大值是右子树最大值和当前节点值的较大者
    return min(lMin, x), max(rMax, x)
}

func isValidBST(root *TreeNode) bool {
    _, mx := dfs(root) // 调用 dfs 检查整棵树的有效性
    return mx != math.MaxInt // 如果返回的最大值等于 math.MaxInt，说明树不是有效的二叉搜索树
}

// 辅助函数，用于计算两个整数的最小值
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// 辅助函数，用于计算两个整数的最大值
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
