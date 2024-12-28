// https://leetcode.cn/problems/find-mode-in-binary-search-tree/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
    tempMap := map[int]int{}
    maxNum := 1
    var dfs func(*TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        if _, ok := tempMap[node.Val]; ok {
            tempMap[node.Val] += 1
            maxNum = max(maxNum, tempMap[node.Val])
        } else {
            tempMap[node.Val] = 1
        }
        dfs(node.Left)
        dfs(node.Right)
    }
    dfs(root)
    ans := []int{}
    for key, value := range tempMap {
        if value == maxNum {
            ans = append(ans, key)
        }
    }
    return ans
}
