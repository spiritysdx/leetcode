/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func decorateRecord(root *TreeNode) [][]int {
    res := [][]int{}
    if root == nil {
        return res
    }
    res = append(res, []int{root.Val})
    flag := true
    cur := []*TreeNode{root}
    for len(cur) > 0 {
        next := []*TreeNode{}
        nextVals := []int{}
        for _, c := range cur {
            if c.Left != nil {
                next = append(next, c.Left)
                nextVals = append(nextVals, c.Left.Val)
            }
            if c.Right != nil {
                next = append(next, c.Right)
                nextVals = append(nextVals, c.Right.Val)
            }
        }
        if flag {
            left, right := 0, len(nextVals)-1
            for left < right {
                nextVals[left], nextVals[right] = nextVals[right], nextVals[left]
                left++
                right--
            }
        }
        if len(nextVals) > 0 {
            res = append(res, nextVals)
        }
        flag = !flag
        cur = next
    }
    return res
}
