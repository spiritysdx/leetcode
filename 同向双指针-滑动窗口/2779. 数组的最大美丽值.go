// https://leetcode.cn/problems/maximum-beauty-of-an-array-after-applying-operation/description/
func maximumBeauty(nums []int, k int) (ans int) {
    // 1. 先对数组进行升序排序
    // 排序是为了方便后续比较数字之间的差距，因为排序后数字的顺序有序，
    // 可以用滑动窗口快速计算满足条件的最长子序列长度。
    slices.Sort(nums)
    // 定义左指针，初始化为 0
    left := 0
    // 2. 遍历数组，用 right 指针表示当前窗口的右端
    for right, x := range nums {
        // 对于当前的数字 x 和窗口左端的数字 nums[left]，判断它们是否满足以下条件：
        // x + k ≥ nums[left] - k
        // 即：nums[left] - x ≤ 2 * k
        // 如果不满足条件，说明当前窗口的范围不合法，需要移动左指针。
        for x - nums[left] > k*2 {
            left++ // 左指针右移，缩小窗口
        }
        // 3. 更新答案
        // 当前窗口的长度为 right - left + 1，更新最大长度。
        ans = max(ans, right-left+1)
    }
    
    // 4. 返回最终的最大长度
    return
}

// 辅助函数：计算两个数的最大值
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
