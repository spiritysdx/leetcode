// https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    ans := [][]int{}
    cur := []*TreeNode{root}
    isLeft := true
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
        if isLeft {
            ans = append(ans, temp)
        } else {
            ans = append(ans, reverse(temp))
        }
        isLeft = !isLeft
    }
    return ans
}

func reverse(temp []int) []int {
    tempR := []int{}
    size := len(temp)
    if size == 0 {
        return nil
    }
    for i := size-1; i>=0; i-- {
        tempR = append(tempR, temp[i])
    }
    return tempR
}
