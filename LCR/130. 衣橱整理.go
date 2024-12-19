func wardrobeFinishing(m int, n int, cnt int) int {
    // 计算数字各位之和的辅助函数
    digitSum := func(x int) int {
        sum := 0
        for x > 0 {
            sum += x % 10 // 获取当前数字的个位并加到总和中
            x /= 10       // 去掉个位
        }
        return sum
    }
    // 创建一个二维布尔数组，用于标记每个格子是否已经访问过
    visited := make([][]bool, m)
    for i := 0; i < m; i++ {
        visited[i] = make([]bool, n) // 初始化每一行的访问记录
    }
    // 深度优先搜索函数
    var dfs func(i, j int) int
    dfs = func(i, j int) int {
        // 边界检查、是否访问过的检查，以及条件是否满足
        if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || digitSum(i)+digitSum(j) > cnt {
            return 0 // 如果越界、已访问或条件不满足，返回0
        }
        // 标记当前格子已访问
        visited[i][j] = true
        // 当前格子满足条件，递归访问向下和向右的格子
        return 1 + dfs(i+1, j) + dfs(i, j+1)
    }
    // 从左上角开始递归
    return dfs(0, 0)
}
