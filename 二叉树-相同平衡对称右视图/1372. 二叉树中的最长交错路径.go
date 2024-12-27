// https://leetcode.cn/problems/longest-zigzag-path-in-a-binary-tree/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {
    var maxLen int
    var dfs func(node *TreeNode, isLeft bool, depth int)
    dfs = func(node *TreeNode, isLeft bool, depth int) {
        if node == nil {
            return
        }
        // 更新最大长度
        if depth > maxLen {
            maxLen = depth
        }
        if isLeft {
            // 当前方向为左，下一步走右
            dfs(node.Left, false, depth+1) // 继续左子树，重置深度
            dfs(node.Right, true, 1)      // 切换方向，从右子树开始
        } else {
            // 当前方向为右，下一步走左
            dfs(node.Right, true, depth+1) // 继续右子树，重置深度
            dfs(node.Left, false, 1)      // 切换方向，从左子树开始
        }
    }
    // 从根节点的两种方向开始搜索
    dfs(root.Left, false, 1)
    dfs(root.Right, true, 1)
    return maxLen
}

// 路径出发的节点可以从根节点也可以从子节点出发
