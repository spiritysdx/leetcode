// https://leetcode.cn/problems/insufficient-nodes-in-root-to-leaf-paths/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
    // 深度优先搜索，返回是否删除当前节点
    var dfs func(node *TreeNode, sum int) bool
    dfs = func(node *TreeNode, sum int) bool {
        // 如果当前节点为空，返回 true，表示不需要保留这个节点
        if node == nil {
            return false
        }
        // 更新当前路径的和
        sum += node.Val
        // 如果是叶子节点，判断路径和是否小于 limit
        if node.Left == nil && node.Right == nil {
            return sum >= limit
        }
        // 递归检查左子树和右子树
        leftValid := dfs(node.Left, sum)
        rightValid := dfs(node.Right, sum)
        // 如果左子树不合法，删除左子树
        if !leftValid {
            node.Left = nil
        }
        // 如果右子树不合法，删除右子树
        if !rightValid {
            node.Right = nil
        }
        // 如果左右子树都被删除，当前节点也需要删除
        return leftValid || rightValid
    }
    // 调用 DFS，从根节点开始
    if !dfs(root, 0) {
        return nil // 如果根节点也被删除，返回 nil
    }
    // 返回修改后的根节点
    return root
}
