// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-postorder-traversal/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(preorder, postorder []int) *TreeNode {
    n := len(preorder) // 获取前序遍历数组的长度
    if n == 0 { // 如果长度为0，说明是空节点
        return nil
    }
    if n == 1 { // 如果长度为1，说明是叶子节点
        return &TreeNode{Val: preorder[0]} // 创建叶子节点并返回
    }
    leftSize := slices.Index(postorder, preorder[1]) + 1 // 找到左子树的大小
    left := constructFromPrePost(preorder[1:1+leftSize], postorder[:leftSize]) // 构建左子树
    right := constructFromPrePost(preorder[1+leftSize:], postorder[leftSize:n-1]) // 构建右子树
    return &TreeNode{preorder[0], left, right} // 构建当前节点并返回
}
