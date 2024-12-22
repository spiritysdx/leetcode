func jewelleryValue(frame [][]int) int {
    // 获取矩阵的行数和列数
    rows := len(frame)
    cols := len(frame[0])
    // 创建一个二维切片 dp，用于存储到达每个位置的最大价值
    dp := make([][]int, rows)
    for i := 0; i < rows; i++ {
        dp[i] = make([]int, cols)
    }
    // 初始化左上角的价值
    dp[0][0] = frame[0][0]
    // 填充第一行的 dp 值
    for j := 1; j < cols; j++ {
        dp[0][j] = dp[0][j-1] + frame[0][j]
    }
    // 填充第一列的 dp 值
    for i := 1; i < rows; i++ {
        dp[i][0] = dp[i-1][0] + frame[i][0]
    }
    // 动态规划填充其余位置的 dp 值
    for i := 1; i < rows; i++ {
        for j := 1; j < cols; j++ {
            dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + frame[i][j]
        }
    }
    // 返回右下角的最大价值
    return dp[rows-1][cols-1]
}

// 工具函数：返回两个整数的最大值
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
