// https://leetcode.cn/problems/move-zeroes/?envType=study-plan-v2&envId=top-100-liked
func moveZeroes(nums []int) {
    // 使用双指针，left指向需要被替换的0的位置，right寻找非0数字
    left := 0
    // 遍历数组，将非零数字与左边的零交换位置
    for right := 0; right < len(nums); right++ {
        if nums[right] != 0 {
            // 交换 left 和 right 位置的元素
            nums[left], nums[right] = nums[right], nums[left]
            left++
        }
    }
}
