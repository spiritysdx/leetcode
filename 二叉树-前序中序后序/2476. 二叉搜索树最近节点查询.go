// https://leetcode.cn/problems/closest-nodes-queries-in-a-binary-search-tree/description/
// 前序遍历每次都得dfs一次，对于大的查询太慢了会超时，此时需要考虑使用中序遍历得到切片后再进行二分查找
// func closestNodes(root *TreeNode, queries []int) [][]int {
//     var dfs func(*TreeNode, int, int) (int, int) // 返回值就是边界值
//     target := queries[0]
//     dfs = func(node *TreeNode, x, y int) (int, int) {
//         if node == nil {
//             return x, y
//         }
//         if node.Val <= target && target-node.Val < target-x {
//             x = node.Val
//         }
//         if node.Val >= target {
//             if y != -1 && node.Val-target < y-target {
//                 y = node.Val
//             } else if y == -1 {
//                 y = node.Val
//             }
//         }
//         x, y = dfs(node.Left, x, y)
//         x, y = dfs(node.Right, x, y)
//         return x, y
//     }
//     x, y := dfs(root, -1, -1)
//     ans := [][]int{[]int{x,y}}
//     if len(queries) > 1 {
//         for _, k := range queries[1:] {
//             target = k
//             x, y = dfs(root, -1, -1)
//             ans = append(ans, []int{x, y})
//         }
//     }
//     return ans
// }
