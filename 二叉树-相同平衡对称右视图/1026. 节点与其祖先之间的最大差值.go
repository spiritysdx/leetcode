// https://leetcode.cn/problems/maximum-difference-between-node-and-ancestor/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(root *TreeNode) int {
    best := 0
    var dfs func(*TreeNode, []int)
    dfs = func(node *TreeNode, temp []int) {
        if node == nil {
            return
        }
        for _, v := range temp {
            best = max(best, abs(v-node.Val))
        }
        temp = append(temp, node.Val)
        dfs(node.Left, temp)
        dfs(node.Right, temp)
    }
    dfs(root, []int{})
    return best
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

// 优化方法：在递归过程中只传递当前路径中的最大值和最小值，而不是整个路径的值数组。这样可以显著降低空间复杂度和运行时间。
// func maxAncestorDiff(root *TreeNode) int {
//     if root == nil {
//         return 0
//     }
//     var dfs func(node *TreeNode, minVal, maxVal int) int
//     dfs = func(node *TreeNode, minVal, maxVal int) int {
//         if node == nil {
//             return maxVal - minVal
//         }
//         // 更新当前路径的最小值和最大值
//         minVal = min(minVal, node.Val)
//         maxVal = max(maxVal, node.Val)
//         // 递归计算左右子树的结果
//         left := dfs(node.Left, minVal, maxVal)
//         right := dfs(node.Right, minVal, maxVal)
//         return max(left, right)
//     }
//     return dfs(root, root.Val, root.Val)
// }
// func min(a, b int) int {
//     if a < b {
//         return a
//     }
//     return b
// }
// func max(a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }
