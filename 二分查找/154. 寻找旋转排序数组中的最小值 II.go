// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array-ii/
func findMin(nums []int) int {
    left, right := -1, len(nums)-1 // 开区间 (-1, n-1)
    for left+1 < right { // 开区间不为空
        mid := left + (right-left)/2
        if nums[mid] < nums[right] { // 最小值在左边
            right = mid
        } else if nums[mid] > nums[right] { // 最小值在右边
            left = mid
        } else {
            right--
        }
    }
    return nums[right]
}
