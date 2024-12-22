// https://leetcode.cn/problems/valid-triangle-number/description/
func triangleNumber(nums []int) int {
    // 排序数组
    sort.Ints(nums) // 使用标准库排序
    res := 0
    size := len(nums)
    // 枚举三角形的最长边 nums[k]
    for k := 2; k < size; k++ {
        i, j := 0, k-1 // 初始化两指针
        for i < j {
            if nums[i] + nums[j] > nums[k] {
                // 如果满足条件，说明 [i, j-1] 范围内的所有组合都满足条件
                res += j - i
                j-- // 收缩右边界
            } else {
                i++ // 增加左边界
            }
        }
    }
    return res
}
