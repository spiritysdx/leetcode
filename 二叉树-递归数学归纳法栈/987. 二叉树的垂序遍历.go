// https://leetcode.cn/problems/vertical-order-traversal-of-a-binary-tree/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 还需要熟悉 slices.SortFunc 的使用
func verticalTraversal(root *TreeNode) [][]int {
    // 每个三元组表示 (row, col, val)，你需要把这些三元组按照 col 分组，也就是把 col 相同的分到同一组，
    // 每组只保留 val，每组的 val 按照 row 从小到大排序，row 相同的按照 val 从小到大排序。分组后的结果就是答案 [[4],[2],[1,5,6],[3],[7]]。
    // 为了获取每个节点的信息，我们可以用 DFS，除了参数 node 外，还需要参数 row 和 col 表示当前节点的行号和列号。
    // 每往下递归一层，就把 row 加一。如果往左儿子递归，就把 col 减一；如果往右儿子递归，就把 col 加一。
    // 定义 DFS 函数，负责遍历每个节点，同时记录其行号和列号。
    var dfs func(*TreeNode, int, int)
    // 定义一个结构 pair，用于存储节点的行号和值。
    type pair struct {
        row, val int
    }
    // 使用一个 map 存储每列的节点数据，键是列号，值是 pair 数组，表示 (row, val)。
    groups := map[int][]pair{}
    // 实现 DFS 函数
    dfs = func(node *TreeNode, row, col int) {
        if node == nil {
            return
        }
        // 将当前节点加入对应列的组
        groups[col] = append(groups[col], pair{row, node.Val})
        // 遍历左子节点，列号减一，行号加一
        dfs(node.Left, row+1, col-1)
        // 遍历右子节点，列号加一，行号加一
        dfs(node.Right, row+1, col+1)
    }
    // 从根节点开始 DFS 遍历，初始行号和列号均为 0
    dfs(root, 0, 0)
    // 收集所有列号，并对列号排序，以便按照从左到右的顺序输出
    keys := make([]int, 0, len(groups))
    for k := range groups {
        keys = append(keys, k)
    }
    // 排序列号
    slices.Sort(keys)
    // 准备答案数组
    ans := make([][]int, len(keys))
    for i, key := range keys {
        // 获取当前列的所有节点
        g := groups[key]
        // 对当前列的节点进行排序，首先按行号排序，行号相同则按值排序
        slices.SortFunc(g, func(a, b pair) int {
            if a.row != b.row {
                return a.row - b.row // 按行号升序排列
            }
            return a.val - b.val // 如果行号相同，按值升序排列
        })
        // 提取排序后的节点值，存入答案
        ans[i] = make([]int, len(g))
        for j, p := range g {
            ans[i][j] = p.val
        }
    }
    // 返回最终结果
    return ans
}
