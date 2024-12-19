/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 主函数，判断 B 是否是 A 的子结构
func isSubStructure(A *TreeNode, B *TreeNode) bool {
    if A == nil || B == nil { // 空树不是子结构
        return false
    }
    // 检查当前节点和左右子树
    return isSameStructure(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

// 判断 B 是否是以 A 为根的子结构
func isSameStructure(A *TreeNode, B *TreeNode) bool {
    if B == nil { // 如果 B 遍历完了，说明匹配成功
        return true
    }
    if A == nil || A.Val != B.Val { // A 遍历完或值不同，匹配失败
        return false
    }
    // 递归比较左右子树
    return isSameStructure(A.Left, B.Left) && isSameStructure(A.Right, B.Right)
}
