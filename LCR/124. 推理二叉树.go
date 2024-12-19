/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deduceTree(preorder []int, inorder []int) *TreeNode {
    // 如果先序遍历或中序遍历为空，直接返回 nil
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }
    // 从先序遍历中获取根节点（第一个元素）
    rootVal := preorder[0]
    // 创建根节点
    root := &TreeNode{Val: rootVal}
    // 在中序遍历中找到根节点的位置
    var rootIndex int
    for i, val := range inorder {
        if val == rootVal {
            rootIndex = i
            break
        }
    }
    // 获取左子树的中序遍历数组（根节点左边的部分）
    leftInorder := inorder[:rootIndex]
    // 获取右子树的中序遍历数组（根节点右边的部分）
    rightInorder := inorder[rootIndex+1:]
    // 获取左子树的先序遍历数组（跳过根节点，长度与左子树的中序遍历数组相同）
    leftPreorder := preorder[1 : 1+len(leftInorder)]
    // 获取右子树的先序遍历数组（剩余部分）
    rightPreorder := preorder[1+len(leftInorder):]
    // 递归构造左子树
    root.Left = deduceTree(leftPreorder, leftInorder)
    // 递归构造右子树
    root.Right = deduceTree(rightPreorder, rightInorder)
    // 返回构造完成的根节点
    return root
}
