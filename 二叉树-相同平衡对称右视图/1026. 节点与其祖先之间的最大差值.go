// https://leetcode.cn/problems/maximum-difference-between-node-and-ancestor/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(root *TreeNode) int {
    best := 0
    var dfs func(*TreeNode, []int)
    dfs = func(node *TreeNode, temp []int) {
        if node == nil {
            return
        }
        for _, v := range temp {
            best = max(best, abs(v-node.Val))
        }
        temp = append(temp, node.Val)
        dfs(node.Left, temp)
        dfs(node.Right, temp)
    }
    dfs(root, []int{})
    return best
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
