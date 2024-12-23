// https://leetcode.cn/problems/count-subarrays-where-max-element-appears-at-least-k-times/
func countSubarrays(nums []int, k int) int64 {
    if len(nums) == 0 || k <= 0 {
        return 0
    }
    // 找到数组中的最大值
    maxNum := 0
    for _, v := range nums {
        maxNum = max(maxNum, v)
    }
    // 滑动窗口计数
    ans := 0
    start := 0
    temp := 0
    size := len(nums)
    for end, x := range nums {
        if x == maxNum {
            temp += 1
        }
        for temp >= k {
            // 累加以 start 开头、以 end 结尾的子数组数量
            ans += size - end // 当前数组可以，那么比这个数组长的后续的数组也可以
            if nums[start] == maxNum {
                temp -= 1
            }
            start += 1
        }
    }
    return int64(ans)
}

// 比较两个整数的最大值
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
