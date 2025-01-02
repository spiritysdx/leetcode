// https://leetcode.cn/problems/even-odd-tree/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isEvenOddTree(root *TreeNode) bool {
    // 偶数 取2余0 严格递增 节点值奇数
    // 奇数 取2余1 严格递减 节点值偶数
    flag := 0
    cur := []*TreeNode{root}
    for len(cur) > 0 {
        nxt := []*TreeNode{}
        temp := []int{}
        for _, c := range cur {
            if flag%2 == 1 && c.Val%2 == 1 {
                return false
            } else if flag%2 == 0 && c.Val%2 == 0 {
                return false
            }
            temp = append(temp, c.Val)
            if c.Left != nil {
                nxt = append(nxt, c.Left)
            }
            if c.Right != nil {
                nxt = append(nxt, c.Right)
            }
        }
        size := len(temp)
        for i := size-2; i>=0; i-- {
            if flag%2 == 1 && temp[i] <= temp[i+1] {
                return false
            } else if flag%2 == 0 && temp[i] >= temp[i+1] {
                return false
            }
        }
        cur = nxt
        flag += 1
    }
    return true
}
