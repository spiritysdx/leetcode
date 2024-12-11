/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // 依旧往递归方向考虑，对比一个树是否是轴对称的，就要证明 root 的左右子节点中，left 节点值与 right 节点值相等，同时还要求 left 左子节点值与 right 右子结点值相等 & left 右子结点值与 right 左子节点值相等：
// left.val == right.val
// left.left.val == right.right.val
// left.right.val == right.left.val
func checkSymmetricTree(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return dfs(root.Left, root.Right)
}

func dfs(node1, node2 *TreeNode) bool {
    if node1 == nil && node2 == nil {
        return true
    }
    if node1 == nil || node2 == nil {
        return false
    }
    if node1.Val != node2.Val {
        return false
    }
    return dfs(node1.Left, node2.Right) && dfs(node1.Right, node2.Left)
}
