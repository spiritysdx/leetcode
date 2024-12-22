// https://leetcode.cn/problems/minimum-size-subarray-sum/description/
func minSubArrayLen(target int, nums []int) int {
    // 由于要求连续子数组，所以left要一直右移
    size := len(nums)
    ans := size + 1 // 初始化为一个大于可能的子数组长度的值
    s := 0
    left := 0
    for right, x := range nums {
        s += x // 增加当前右指针指向的值
        // 当子数组和大于等于目标值时，尝试缩小窗口
        for s >= target {
            ans = min(ans, right-left+1) // 更新最小长度
            s -= nums[left] // 缩小窗口
            left++
        }
    }
    // 如果 ans 没有被更新过，说明不存在满足条件的子数组
    if ans == size+1 {
        return 0
    }
    return ans
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
