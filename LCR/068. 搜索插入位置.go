func searchInsert(nums []int, target int) int {
    // 通过一次二分查找找到第一个大于等于 target 的索引
    left, right := 0, len(nums)
    for left < right {
        mid := left + (right-left)/2
        if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}
