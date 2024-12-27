// https://leetcode.cn/problems/sum-root-to-leaf-numbers/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    ans := 0
    // 深度优先搜索函数
    var dfs func(*TreeNode, int)
    dfs = func(node *TreeNode, currentSum int) {
        if node == nil {
            return
        }
        // 更新当前路径的数字
        currentSum = currentSum*10 + node.Val
        // 如果到达叶子节点，将当前路径数字加入总和
        if node.Left == nil && node.Right == nil {
            ans += currentSum
            return
        }
        // 递归访问左子树和右子树
        dfs(node.Left, currentSum)
        dfs(node.Right, currentSum)
    }
    dfs(root, 0)
    return ans
}

// func sumNumbers(root *TreeNode) int {
//     return dfs(root, 0)
// }

// func dfs(root *TreeNode, x int) int {
//     if root == nil {
//         return 0
//     }
//     x = x*10 + root.Val
//     if root.Left == root.Right { // root 是叶子节点
//         return x
//     }
//     return dfs(root.Left, x) + dfs(root.Right, x)
// }
