/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil || root == p || root == q {
        return root // 如果当前节点为空，或当前节点是 p 或 q，返回当前节点
    }
    left := lowestCommonAncestor(root.Left, p, q)   // 在左子树中查找
    right := lowestCommonAncestor(root.Right, p, q) // 在右子树中查找
    if left != nil && right != nil {
        return root // 如果左右子树都找到了 p 或 q，说明当前节点是最近公共祖先
    }
    if left != nil {
        return left // 如果左子树找到 p 或 q，返回左子树结果
    }
    return right    // 如果右子树找到 p 或 q，返回右子树结果
}
