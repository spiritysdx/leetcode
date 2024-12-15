// https://leetcode.cn/problems/maximum-count-of-positive-integer-and-negative-integer/
func maximumCount(nums []int) int {
    // 分别找到小于0的数和大于0的数的数目 即找0的位置
    target := 0
    left := 0
    right := len(nums)
    for left < right {
        mid := left + (right-left)/2
        if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    start := left
    left = 0
    right = len(nums)
    for left < right {
        mid := left + (right-left)/2
        if nums[mid] <= target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    end := left
    return max(start, len(nums)-end)
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}
