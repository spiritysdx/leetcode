/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	if root == nil || root == q || root == q {
//         return root
//     }
//     left := lowestCommonAncestor(root, p, q)
//     right := lowestCommonAncestor(root, p, q)
//     if left != nil && right != nil {
//         return root
//     }
//     if left != nil {
//         return left
//     }
//     return right
// }

// 二叉树
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    current := root.Val
	if p.Val < current && q.Val < current { // 都小于当前节点，证明节点都在左子树中
		return lowestCommonAncestor(root.Left, p, q)

	}
	if p.Val > current && q.Val > current { // 都大于当前节点，证明节点都在右子树中
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
