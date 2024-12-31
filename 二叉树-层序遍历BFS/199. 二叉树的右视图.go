// https://leetcode.cn/problems/binary-tree-right-side-view/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func rightSideView(root *TreeNode) []int {
//     ans := []int{}
//     var find func(*TreeNode, int)
//     find = func(node *TreeNode, depth int) {
//         if node == nil {
//             return
//         }
//         if depth == len(ans) {
//             ans = append(ans, node.Val)
//         }
//         find(node.Right, depth+1)
//         find(node.Left, depth+1)
//     }
//     find(root, 0)
//     return ans
// }

func rightSideView(root *TreeNode) []int {
    ans := []int{}
    if root == nil {
        return ans
    }
    cur := []*TreeNode{root}
    for len(cur) > 0 {
        nxt := []*TreeNode{}
        for _, c := range cur {
            if c.Left != nil {
                nxt = append(nxt, c.Left)
            }
            if c.Right != nil {
                nxt = append(nxt, c.Right)
            }
        }
        ans = append(ans, cur[len(cur)-1].Val)
        cur = nxt
    }
    return ans
}
