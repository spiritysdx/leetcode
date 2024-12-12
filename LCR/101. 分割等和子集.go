func canPartition(nums []int) bool {
    size := len(nums)
    if size < 2 { // 数组只有一个元素时，无法分割成两个子集
        return false
    }
    sumNum := 0
    maxNum := 0
    for _, k := range nums {
        sumNum += k // 计算数组的总和
        maxNum = max(maxNum, k) // 找到数组中的最大值
    }
    if sumNum%2 != 0 { // 如果总和为奇数，无法分为两个和相等的子集
        return false
    }
    target := sumNum / 2 // 目标值是总和的一半
    if maxNum > target { // 如果最大元素大于目标值，剩下元素之和必然无法等于最大元素
        return false
    }
    // 动态规划求解
    // dp[i] 表示是否可以从数组中选出若干个数，其和恰好为 i
    dp := make([]bool, target+1)
    dp[0] = true // 初始化：和为 0 是永远可以达到的（选择空集）
    for _, num := range nums { // 遍历数组中的每个数
        for j := target; j >= num; j-- { // 从后往前更新 dp 数组
            dp[j] = dp[j] || dp[j-num] // 如果能通过前面的状态转移到达当前状态，则更新为 true
        }
    }
    return dp[target] // 返回是否可以达到目标值
}

// 辅助函数：返回两个整数中的较大值
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
