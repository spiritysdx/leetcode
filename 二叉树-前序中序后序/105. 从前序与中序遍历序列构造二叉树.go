// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
func buildTree(preorder, inorder []int) *TreeNode {
    n := len(preorder) // 获取前序遍历数组的长度
    if n == 0 { // 如果长度为 0，说明是空节点，返回 nil
        return nil
    }
    leftSize := slices.Index(inorder, preorder[0]) // 找到根节点在中序遍历中的索引，计算左子树的大小
    left := buildTree(preorder[1:1+leftSize], inorder[:leftSize]) // 递归构建左子树
    right := buildTree(preorder[1+leftSize:], inorder[1+leftSize:]) // 递归构建右子树
    return &TreeNode{preorder[0], left, right} // 构造当前节点并返回
}
