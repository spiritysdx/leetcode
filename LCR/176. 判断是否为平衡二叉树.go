/**
 * 二叉树节点定义
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 不能使用层序遍历，需要递归计算左右子树深度并实时判断是否为非平衡二叉树
// 不用递归不行，因为必须要算子树本身是否符合条件

// 判断二叉树是否为平衡二叉树
func isBalanced(root *TreeNode) bool {
    // 调用辅助函数，返回树是否平衡的布尔值
    _, isBal := checkBalance(root)
    return isBal
}

// 辅助函数：返回树的高度和是否平衡
func checkBalance(root *TreeNode) (int, bool) {
    if root == nil { 
        // 空节点高度为0，且是平衡的
        return 0, true
    }
    // 递归检查左子树的高度和是否平衡
    leftHeight, leftBalanced := checkBalance(root.Left)
    if !leftBalanced { 
        // 如果左子树不平衡，直接返回
        return 0, false
    }
    // 递归检查右子树的高度和是否平衡
    rightHeight, rightBalanced := checkBalance(root.Right)
    if !rightBalanced { 
        // 如果右子树不平衡，直接返回
        return 0, false
    }
    // 如果左右子树高度差大于1，则当前节点不平衡
    if abs(leftHeight-rightHeight) > 1 {
        return 0, false
    }
    // 当前节点的高度为左右子树高度的较大值加1
    return max(leftHeight, rightHeight) + 1, true
}

// 计算绝对值
func abs(x int) int {
    if x < 0 { 
        return -x // 如果x小于0，返回其相反数
    }
    return x // 否则返回x本身
}

// 返回两个数中的较大值
func max(a, b int) int {
    if a > b { 
        return a // 如果a大于b，返回a
    }
    return b // 否则返回b
}
