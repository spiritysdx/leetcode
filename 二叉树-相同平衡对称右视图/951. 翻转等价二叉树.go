// https://leetcode.cn/problems/flip-equivalent-binary-trees/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// flipEquiv 函数判断两棵二叉树是否是翻转等价的
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
    // 如果两个节点都为 nil，返回 true
    if root1 == nil && root2 == nil {
        return true
    }
    // 如果只有一个节点为 nil 或者节点值不相等，返回 false
    if root1 == nil || root2 == nil || root1.Val != root2.Val {
        return false
    }
    // 判断不翻转的情况和翻转的情况是否满足条件
    // 不翻转的情况：左子树对比左子树，右子树对比右子树
    // 翻转的情况：左子树对比右子树，右子树对比左子树
    return (flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)) ||
           (flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left))
}
