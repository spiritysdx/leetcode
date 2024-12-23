// https://leetcode.cn/problems/length-of-longest-subarray-with-at-most-k-frequency/description/
func maxSubarrayLength(nums []int, k int) int {
    ans := 0 // 记录最长子数组的长度
    start := 0 // 滑动窗口的起始位置
    count := map[int]int{} // 用于记录当前窗口中每个数字的出现次数
    for index, x := range nums {
        if _, ok := count[x]; ok { // 如果数字 x 已经存在于 map 中
            count[x] += 1 // 增加数字 x 的计数

            // 如果数字 x 的出现次数超过了 k
            for count[x] > k {
                count[nums[start]] -= 1 // 减少窗口起始位置对应数字的计数
                start += 1 // 将窗口起始位置右移
            }
        } else {
            count[x] = 1 // 如果数字 x 不在 map 中，将其加入并设置计数为 1
        }

        // 到这里保证 map 中的值都符合好数组的要求
        ans = max(ans, index-start+1) // 更新最长子数组的长度
    }
    return ans // 返回结果
}

// 工具函数，用于返回两个整数中的较大值
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
