// https://leetcode.cn/problems/evaluate-boolean-binary-tree/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int            // 节点的值
 *     Left *TreeNode     // 左子节点
 *     Right *TreeNode    // 右子节点
 * }
 */
func evaluateTree(root *TreeNode) bool {
	// 如果当前节点是叶子节点（没有左子节点）
	if root.Left == nil {
		// 叶子节点的值为1表示true，否则为false
		return root.Val == 1
	}
	// 如果当前节点的值是2，表示逻辑或运算
	if root.Val == 2 {
		// 返回左子树或右子树的逻辑或结果
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	}
	// 否则（当前节点的值是3），表示逻辑与运算
	// 返回左子树和右子树的逻辑与结果
	return evaluateTree(root.Left) && evaluateTree(root.Right)
}
