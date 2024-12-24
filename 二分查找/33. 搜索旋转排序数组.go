// https://leetcode.cn/problems/search-in-rotated-sorted-array/description/
func search(nums []int, target int) int {
    left := 0                                    // 初始化左指针
    size := len(nums)                           // 获取数组长度
    right := size - 1                           // 初始化右指针
    // 找到数组的旋转点（最小值的索引）
    for left < right {
        mid := left + (right-left)/2            // 计算中间位置，避免溢出
        if nums[mid] > nums[right] {           // 如果中间值大于右边界值，最小值在右侧
            left = mid + 1
        } else {                               // 否则最小值在左侧或就是 mid
            right = mid
        }
    }
    minIndex := left                            // 最小值索引即为旋转点
    // 重新初始化左右边界，根据目标值所在范围选择子数组
    left, right = 0, size-1                     // 重新初始化左右指针
    if target >= nums[minIndex] && target <= nums[right] { // 如果目标值在右侧的递增段
        left = minIndex                         // 搜索范围从旋转点开始到数组末尾
    } else {                                   // 如果目标值在左侧的递增段
        right = minIndex - 1                   // 搜索范围从数组起点到旋转点之前
    }
    // 二分查找目标值
    for left <= right {
        mid := left + (right-left)/2            // 计算中间位置
        if nums[mid] == target {               // 找到目标值
            return mid
        } else if nums[mid] < target {         // 如果中间值小于目标值，移动左指针
            left = mid + 1
        } else {                               // 如果中间值大于目标值，移动右指针
            right = mid - 1
        }
    }
    return -1                                   // 未找到目标值返回 -1
}
