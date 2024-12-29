// https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    temp := [][]int{}
    cur := []*TreeNode{root}
    for len(cur) > 0 {
        tp := []int{}
        nxt := []*TreeNode{}
        for _, c := range cur {
            tp = append(tp, c.Val)
            if c.Left != nil {
                nxt = append(nxt, c.Left)
            }
            if c.Right != nil {
                nxt = append(nxt, c.Right)
            }
        }
        temp = append(temp, tp)
        cur = nxt
    }
    ans := [][]int{}
    size := len(temp)
    for i := size-1; i>=0; i-- {
        ans = append(ans, temp[i])
    }
    return ans
}
