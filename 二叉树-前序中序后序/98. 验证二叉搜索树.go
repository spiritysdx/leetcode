// https://leetcode.cn/problems/validate-binary-search-tree/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int           // 节点的值
 *     Left *TreeNode    // 左子节点
 *     Right *TreeNode   // 右子节点
 * }
 */
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
