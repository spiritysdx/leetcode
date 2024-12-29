// https://leetcode.cn/problems/binary-tree-level-order-traversal/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    ans := [][]int{}
    cur := []*TreeNode{root}
    for len(cur) > 0 {
        temp := []int{}
        nxt := []*TreeNode{}
        for _, c := range cur {
            temp = append(temp, c.Val)
            if c.Left != nil {
                nxt = append(nxt, c.Left)
            }
            if c.Right != nil {
                nxt = append(nxt, c.Right)
            }
        }
        cur = nxt
        ans = append(ans, temp)
    }
    return ans
}
