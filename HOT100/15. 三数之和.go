// https://leetcode.cn/problems/3sum/description/?envType=study-plan-v2&envId=top-100-liked
func threeSum(nums []int) [][]int {
    slices.Sort(nums)
    size := len(nums)
    ans := [][]int{}
    for i := 0; i < size-2; i++ {
        // 跳过重复的第一个数
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        left := i + 1
        right := size - 1
        for left < right {
            sum := nums[i] + nums[left] + nums[right]
            if sum < 0 {
                left++
            } else if sum > 0 {
                right--
            } else {
                // 找到一组解
                ans = append(ans, []int{nums[i], nums[left], nums[right]})
                // 跳过重复的左边界
                for left < right && nums[left] == nums[left+1] {
                    left++
                }
                // 跳过重复的右边界
                for left < right && nums[right] == nums[right-1] {
                    right--
                }
                // 移动左右指针
                left++
                right--
            }
        }
    }
    return ans
}
