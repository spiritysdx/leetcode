/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func decorateRecord(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    // 层序遍历
    cur := []*TreeNode{root}
    res := []int{root.Val}
    for len(cur) > 0 {
        next := []*TreeNode{}
        for _, c := range cur {
            if c.Left != nil {
                next = append(next, c.Left)
                res = append(res, c.Left.Val)
            }
            if c.Right != nil {
                next = append(next, c.Right)
                res = append(res, c.Right.Val)
            }
        }
        cur = next
    }
    return res
}
