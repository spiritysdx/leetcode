// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    size := len(postorder)
    if size == 0 {
        return nil
    }
    index := slices.Index(inorder, postorder[size-1]) // 左子树大小
    node := &TreeNode{Val: postorder[size-1]}
    node.Left = buildTree(inorder[:index], postorder[:index])
    node.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])
    return node
}
