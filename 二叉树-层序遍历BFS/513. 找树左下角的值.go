// https://leetcode.cn/problems/find-bottom-left-tree-value/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
    cur := []*TreeNode{root}
    ans := 0
    for len(cur) > 0 {
        nxt := []*TreeNode{}
        for _, c := range cur {
            ans = c.Val
            if c.Right != nil {
                nxt = append(nxt, c.Right)
            }
            if c.Left != nil {
                nxt = append(nxt, c.Left)
            }
        }
        cur = nxt
    }
    return ans
}

func findBottomLeftValue(root *TreeNode) int {
    node := root
    q := []*TreeNode{root}
    for len(q) > 0 {
        node, q = q[0], q[1:]
        if node.Right != nil {
            q = append(q, node.Right)
        }
        if node.Left != nil {
            q = append(q, node.Left)
        }
    }
    return node.Val
}
