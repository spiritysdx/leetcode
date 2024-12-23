// https://leetcode.cn/problems/minimum-operations-to-reduce-x-to-zero/description/
func minOperations(nums []int, x int) int {
    target := -x // 将目标值初始化为-x，稍后会通过累加数组元素调整
    for _, x := range nums {
        target += x // 计算数组元素的总和减去x的值，最终目标是找到和为target的最长子数组
    }
    if target < 0 { // 如果全部移除也无法满足要求（总和小于x），返回-1
        return -1
    }
    ans, left, sum := -1, 0, 0 // 初始化结果值ans为-1，滑动窗口的左指针left，当前子数组的总和sum
    for right, x := range nums { // 遍历数组，right为滑动窗口的右指针
        sum += x // 将当前元素加到子数组总和中
        for sum > target { // 如果当前子数组总和大于目标值，则缩小子数组长度
            sum -= nums[left] // 从左边移除元素
            left++ // 左指针右移
        }
        if sum == target { // 如果找到和为目标值的子数组
            ans = max(ans, right-left+1) // 更新最长子数组的长度
        }
    }
    if ans < 0 { // 如果没有找到满足条件的子数组，返回-1
        return -1
    }
    return len(nums) - ans // 返回移除的最小操作数（数组总长度减去最长子数组长度）
}

// 辅助函数，用于求两个数的较大值
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
// 把问题转换成「从 nums 中移除一个最长的子数组，使得剩余元素的和为 x」。
// 换句话说，要从 nums 中找最长的子数组，其元素和等于 s−x，这里 s 为 nums 所有元素之和。
// 由于本题没有负数，我们可以用滑动窗口
// 最后答案为 nums 的长度减去最长子数组的长度。
