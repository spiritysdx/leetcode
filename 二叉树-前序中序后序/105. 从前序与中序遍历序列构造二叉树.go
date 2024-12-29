// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
func buildTree(preorder, inorder []int) *TreeNode {
    n := len(preorder)
    if n == 0 { // 空节点
        return nil
    }
    leftSize := slices.Index(inorder, preorder[0]) // 左子树的大小
    left := buildTree(preorder[1:1+leftSize], inorder[:leftSize])
    right := buildTree(preorder[1+leftSize:], inorder[1+leftSize:])
    return &TreeNode{preorder[0], left, right}
}
