func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	// 定义dp数组，dp[i]表示到达第i阶的最小花费
	dp := make([]int, n+1)
	// 动态规划计算
	// 初始状态：
	// 到达第 0 个阶梯不需要花费：dp[0] = 0；
        // 到达第 1 个阶梯也不需要花费：dp[1] = 0。
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2]) // 状态转移方程
	}
	return dp[n]
}

// 求两个数的最小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
