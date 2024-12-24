// https://leetcode.cn/problems/find-peak-element/

// 非二分法暴力解
func findPeakElement(nums []int) int {
    size := len(nums)
    if size == 1 {
        return 0
    }
    if size == 2 {
        if nums[0] > nums[1] {
            return 0
        }
        return 1
    }
    for i := 1; i < size-1; i++ {
        if nums[i-1] < nums[i] && nums[i+1] < nums[i] {
            return i
        } 
    }
    if nums[0] > nums[1] {
        return 0
    }
    return size-1
}

// 二分法
