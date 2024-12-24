// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/description/
func findMin(nums []int) int {
    // 定义左右指针
    left, right := 0, len(nums)-1
    // 当左指针小于右指针时继续查找
    for left < right {
        // 计算中间位置
        mid := left + (right-left)/2
        // 如果中间值大于右边界的值
        // 说明最小值在右半部分，因为正常升序数组不会出现这种情况
        if nums[mid] > nums[right] {
            // 将左指针移到中间位置的右边
            left = mid + 1
        } else {
            // 如果中间值小于等于右边界的值
            // 说明最小值在左半部分（包括中间值）
            right = mid
        }
    }
    // 返回左指针位置的值，此时就是最小值
    return nums[left]
}
