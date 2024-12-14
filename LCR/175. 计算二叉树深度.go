/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func calculateDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    res := 0
    current := []*TreeNode{root}
    for len(current) > 0 {
        next := []*TreeNode{}
        for _, c := range current {
            if c.Left != nil {
                next = append(next, c.Left)
            }
            if c.Right != nil {
                next = append(next, c.Right)
            }
        }
        current = next
        res += 1
    }
    return res
}
