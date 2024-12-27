// https://leetcode.cn/problems/balanced-binary-tree/description/
/**
 * Definition for a binary tree node.
 * 二叉树节点的定义
 * type TreeNode struct {
 *     Val int              // 节点值
 *     Left *TreeNode       // 左子节点
 *     Right *TreeNode      // 右子节点
 * }
 */
func isBalanced(root *TreeNode) bool {
    // 判断二叉树是否是平衡二叉树
    return getHeight(root) != -1 // 如果高度返回 -1，则表示不平衡
}

func getHeight(node *TreeNode) int {
    if node == nil {
        return 0 // 空节点的高度为 0
    }
    leftHeight := getHeight(node.Left) // 递归获取左子树高度
    if leftHeight == -1 {
        return -1 // 如果左子树不平衡，直接返回 -1
    }
    rightHeight := getHeight(node.Right) // 递归获取右子树高度
    if rightHeight == -1 || abs(leftHeight - rightHeight) > 1 {
        return -1 // 如果右子树不平衡或左右高度差超过 1，返回 -1
    }
    return max(leftHeight, rightHeight) + 1 // 返回当前节点的高度
}

func abs(a int) int {
    if a < 0 {
        return -a // 返回绝对值
    }
    return a
}
