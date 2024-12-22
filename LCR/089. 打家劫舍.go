func rob(nums []int) int {
    size := len(nums)
    if size == 0 {
        return 0
    }
    if size == 1 {
        return nums[0]
    }
    // 动态规划数组
    dp := make([]int, size)
    // 初始化
    dp[0] = nums[0]
    dp[1] = max(nums[0], nums[1])
    // 填充 dp 数组
    for i := 2; i < size; i++ {
        dp[i] = max(dp[i-1], dp[i-2]+nums[i])
    }
    // 返回最大金额
    return dp[size-1]
}

// 辅助函数
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
