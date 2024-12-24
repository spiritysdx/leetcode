// https://leetcode.cn/problems/minimize-maximum-of-array/description/ 

// func minimizeArrayValue(nums []int) int {
//     s := 0  // 初始化总和s为0
//     ans := 0  // 初始化最小化后的最大值ans为0
//     for i, x := range nums {  // 遍历nums数组，i为索引，x为元素值
//         s += x  // 将当前元素x加到总和s上
//         ans = max(ans, (s+i)/(i+1))  // 计算当前总和的平均值，并更新ans为最大平均值
//     }
//     return ans  // 返回最小化后的最大值
// }
// 算法通过逐步计算前缀子数组的平均值，并动态更新ans来确保返回一个最小的“最大平均值”，使得每个子数组的平均值尽可能均匀，从而达到优化的目的。

// 下面是二分法
func minimizeArrayValue(nums []int) int {
    // 二分法的范围设置
    left, right := 0, max(nums) // 最小值0，最大值为nums数组的最大值
    // 二分法循环
    for left < right {
        mid := (left + right) / 2 // 中间值作为猜测的limit
        if canAchieveLimit(nums, mid) {  // 如果mid作为limit可以满足条件
            right = mid  // 缩小范围，试更小的limit
        } else {
            left = mid + 1  // 否则增大limit
        }
    }
    return left
}

// 判断在给定的limit下是否能满足条件
func canAchieveLimit(nums []int, limit int) bool {
    extra := 0  // 额外的多余部分
    // 从后向前遍历数组，调整多余部分
    for i := len(nums) - 1; i >= 1; i-- {
        if nums[i] + extra > limit {  // 当前值加上多余部分超过limit
            extra = nums[i] + extra - limit  // 需要去掉多余的部分
        } else {
            extra = 0  // 如果当前值不超过limit，reset extra为0
        }
    }
    // 最后判断nums[0]是否满足条件
    return nums[0] + extra <= limit
}

// 求数组最大值
func max(nums []int) int {
    maxVal := nums[0]
    for _, num := range nums {
        if num > maxVal {
            maxVal = num
        }
    }
    return maxVal
}
