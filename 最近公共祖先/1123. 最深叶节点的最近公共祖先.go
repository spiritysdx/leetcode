/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
    s, _ := dfs(root)
    return s
}

func dfs(root *TreeNode) (*TreeNode, int) {
    if root == nil {
        return nil, 0
    }
    left, leftHeight := dfs(root.Left)
    right, rightHeight := dfs(root.Right)
    if leftHeight > rightHeight { // 左子树高
        return left, leftHeight + 1
    }
    if leftHeight < rightHeight { // 右子树高
        return right, rightHeight + 1
    }
    return root, leftHeight + 1 // 高度相等，返回当前节点
}
