// https://leetcode.cn/problems/kth-smallest-element-in-a-bst/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    // 中序遍历然后取值即可
    temp := []int{}
    var dfs func(*TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        dfs(node.Left)
        temp = append(temp, node.Val)
        dfs(node.Right)
    }
    dfs(root)
    return temp[k-1]
}
