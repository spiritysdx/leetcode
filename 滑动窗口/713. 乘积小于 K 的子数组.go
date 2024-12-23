// https://leetcode.cn/problems/subarray-product-less-than-k/description/
func numSubarrayProductLessThanK(nums []int, k int) int {
    // 如果 k 小于等于 1，乘积不可能小于 k，直接返回 0
    if k <= 1 {
        return 0
    }
    temp := 1 // 当前子数组的乘积
    left := 0 // 左指针
    ans := 0  // 符合条件的子数组数量
    // 遍历数组，right 为右指针
    for right, x := range nums {
        temp *= x // 更新当前子数组的乘积
        // 当乘积大于等于 k 时，移动左指针，缩小子数组范围
        for temp >= k {
            temp /= nums[left] // 除去左指针位置的元素
            left++             // 左指针右移
        }
        // 计算以 right 为右端点的子数组数量，并累加到结果中
        ans += right - left + 1
    }
    return ans // 返回最终的子数组数量
}
