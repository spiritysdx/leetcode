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
func findPeakElement(nums []int) int {
    left, right := -1, len(nums)-1 // 初始化左右边界为开区间 (-1, n-1)
    for left+1 < right { // 当开区间不为空时继续循环
        mid := left + (right-left)/2 // 计算中间位置，防止溢出
        if nums[mid] > nums[mid+1] { // 如果当前元素大于右侧元素
            right = mid // 峰值可能在左侧或就是当前元素，更新右边界
        } else { // 如果当前元素小于或等于右侧元素
            left = mid // 峰值在右侧，更新左边界
        }
    }
    return right // 返回右边界，即峰值元素的索引
}
