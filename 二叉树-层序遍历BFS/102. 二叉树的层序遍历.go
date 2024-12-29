// https://leetcode.cn/problems/binary-tree-level-order-traversal/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 两个切片
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

// 一个队列
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    ans := [][]int{}
    cur := []*TreeNode{root}
    for len(cur) > 0 {
        size := len(cur)
        temp := []int{}
        for i := 0; i < size; i++ {
            node := cur[0]
            cur = cur[1:]
            temp = append(temp, node.Val)
            if node.Left != nil {
                cur = append(cur, node.Left)
            }
            if node.Right != nil {
                cur = append(cur, node.Right)
            }
        }
        ans = append(ans, temp)
    }
    return ans
}
