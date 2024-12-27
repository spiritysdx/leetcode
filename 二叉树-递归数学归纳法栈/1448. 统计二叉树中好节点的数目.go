// https://leetcode.cn/problems/count-good-nodes-in-binary-tree/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    ans := 0
    var dfs func(*TreeNode, []int)
    dfs = func(node *TreeNode, temp []int) {
        if node == nil {
            return
        }
        temp = append(temp, node.Val)
        size := len(temp)
        if size > 1 && temp[size-1] < temp[size-2] {
            temp = temp[:size-1]
        } else {
            ans += 1
        }
        if node.Left != nil {
            dfs(node.Left, temp)
        }
        if node.Right != nil {
            dfs(node.Right, temp)
        }
    }
    dfs(root, []int{})
    return ans
}

// func dfs(root *TreeNode, mx int) int {
//     if root == nil {
//         return 0
//     }
//     left := dfs(root.Left, max(mx, root.Val))
//     right := dfs(root.Right, max(mx, root.Val))
//     if mx <= root.Val {
//         return left + right + 1
//     }
//     return left + right
// }

// func goodNodes(root *TreeNode) int {
//     return dfs(root, math.MinInt) // 也可以写 root.Val
// }

// func max(a, b int) int { if b > a { return b }; return a }
