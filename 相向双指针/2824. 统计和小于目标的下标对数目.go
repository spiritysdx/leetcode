// https://leetcode.cn/problems/count-pairs-whose-sum-is-less-than-target/
func countPairs(nums []int, target int) (ans int) {
	// 对 nums 数组进行排序
	slices.Sort(nums)
	// 初始化左右指针
	left, right := 0, len(nums)-1
	// 当左指针小于右指针时进行循环
	for left < right {
		// 如果当前两个数的和小于目标值
		if nums[left]+nums[right] < target {
			// 统计从 left 到 right 之间的所有对数
			ans += right - left
			// 移动左指针
			left++
		} else {
			// 如果当前两个数的和大于等于目标值，则移动右指针
			right--
		}
	}
	// 返回最终的计数结果
	return
}
