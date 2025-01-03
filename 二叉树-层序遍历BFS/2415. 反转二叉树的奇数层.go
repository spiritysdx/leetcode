// https://leetcode.cn/problems/reverse-odd-levels-of-binary-tree/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func reverseOddLevels(root *TreeNode) *TreeNode {
    // 只反转值还行
    if root == nil {
        return nil
    }
    flag := 0 // 标记层数，根节点为第 0 层
    cur := []*TreeNode{root}
    for len(cur) > 0 {
        nxt := []*TreeNode{}
        // 奇数层时反转值
        if flag%2 == 1 {
            for i, j := 0, len(cur)-1; i < j; i, j = i+1, j-1 {
                cur[i].Val, cur[j].Val = cur[j].Val, cur[i].Val
            }
        }
        // 准备下一层的节点
        for _, node := range cur {
            if node.Left != nil {
                nxt = append(nxt, node.Left)
            }
            if node.Right != nil {
                nxt = append(nxt, node.Right)
            }
        }
        cur = nxt
        flag++ // 层数加一
    }
    return root
}
