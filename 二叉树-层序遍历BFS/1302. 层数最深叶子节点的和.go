// https://leetcode.cn/problems/deepest-leaves-sum/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
    cur := []*TreeNode{root}
    var ans int
    for len(cur) > 0 {
        nxt := []*TreeNode{}
        ans = 0
        for _, c := range cur {
            ans += c.Val
            if c.Left != nil {
                nxt = append(nxt, c.Left)
            }
            if c.Right != nil {
                nxt = append(nxt, c.Right)
            }
        }
        if len(nxt) == 0 {
            return ans
        }
        cur = nxt
    }
    return 0
}
