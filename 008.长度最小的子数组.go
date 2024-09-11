func minSubArrayLen(target int, nums []int) int {
    size := len(nums) // 获取数组长度
    if size == 0 {
        return 0 // 如果数组为空，返回0
    }
    ans := size + 1 // 初始化最小长度为比数组长度大1的值，表示不存在的情况
    sum := 0 // 当前子数组的和
    left := 0 // 左指针初始化为0
    for right := 0; right < size; right++ { // 右指针从0开始遍历数组
        sum += nums[right] // 将当前元素加入到子数组的和中
        // 当子数组的和大于等于目标值时，尝试缩小窗口
        for sum >= target {
            ans = min(ans, right-left+1) // 更新最小长度
            sum -= nums[left] // 减去左指针指向的元素
            left++ // 左指针右移，缩小窗口
        }
    }
    if ans == size+1 {
        return 0 // 如果最小长度没有被更新，返回0
    }
    return ans // 返回最小长度
}

// 辅助函数，用于计算两个数的最小值
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
