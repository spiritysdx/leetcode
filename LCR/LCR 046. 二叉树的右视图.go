/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    cur := []*TreeNode{root}
    res := []int{}
    if root == nil {
        return res
    }
    for len(cur) > 0 {
        next := []*TreeNode{}
        size := len(cur)
        for index, c := range cur {
            if c.Left != nil {
                next = append(next, c.Left)
            }
            if c.Right != nil {
                next = append(next, c.Right)
            }
            if index == size-1 {
                res = append(res, c.Val)
            }
        }
        cur = next
    }
    return res
}
