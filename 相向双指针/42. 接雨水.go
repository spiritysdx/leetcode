// https://leetcode.cn/problems/trapping-rain-water/description/
func trap(height []int) int {
    size := len(height)
    if size == 0 {
        return 0
    }
    pre_height := make([]int, size) // 前缀最大
    sur_height := make([]int, size) // 后缀最大
    // 构建前缀最大值数组
    pre_height[0] = height[0]
    for i := 1; i < size; i++ {
        pre_height[i] = max(pre_height[i-1], height[i])
    }
    // 构建后缀最大值数组
    sur_height[size-1] = height[size-1]
    for j := size - 2; j >= 0; j-- {
        sur_height[j] = max(sur_height[j+1], height[j])
    }
    // 计算储水量
    res := 0
    for i := 0; i < size; i++ {
        res += min(pre_height[i], sur_height[i]) - height[i]
    }
    return res
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
