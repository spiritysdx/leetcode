// https://leetcode.cn/problems/most-frequent-subtree-sum/description/
func findFrequentTreeSum(root *TreeNode) []int {
    // 键值对 子树和：频率
    tempMap := map[int]int{}
    maxFreq := 0 // 记录最大频率
    // 递归函数，计算子树和
    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        // 计算左右子树和
        leftSum := dfs(node.Left)
        rightSum := dfs(node.Right)
        // 当前子树和
        sum := node.Val + leftSum + rightSum
        // 更新哈希表
        tempMap[sum]++
        maxFreq = max(maxFreq, tempMap[sum]) // 更新最大频率
        return sum
    }
    dfs(root)
    // 找出频率最高的子树和
    ans := []int{}
    for key, value := range tempMap {
        if value == maxFreq {
            ans = append(ans, key)
        }
    }
    return ans
}

// 辅助函数：计算两个数的最大值
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
