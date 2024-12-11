/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flipTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    flipTree(root.Left)
    flipTree(root.Right)
    root.Left,root.Right=root.Right,root.Left
    return root
}
