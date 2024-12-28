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


func closestNodes(root *TreeNode, queries []int) [][]int {
    // 用中序遍历构建非递减序列
    temp := []int{}
    var dfs func(*TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        dfs(node.Left)
        temp = append(temp, node.Val)
        dfs(node.Right)
    }
    dfs(root)
    // 二分查找最近的值
    ans := [][]int{}
    for _, target := range queries {
        l, r := 0, len(temp)-1
        for l <= r {
            mid := l + (r-l)/2
            if temp[mid] < target {
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
        // 当查找结束时：
        // r 指向目标值左边的最近一个值，或越界到 -1。
        // l 指向目标值右边的最近一个值，或越界到 len(temp)。
        // 计算左边和右边的最近值
        leftClosest := -1
        if r >= 0 {
            leftClosest = temp[r]
        }
        rightClosest := -1
        if l < len(temp) {
            rightClosest = temp[l]
        }
        // 添加结果
        if leftClosest == target || rightClosest == target {
            ans = append(ans, []int{target, target})
        } else {
            ans = append(ans, []int{leftClosest, rightClosest})
        }
    }
    return ans
}
