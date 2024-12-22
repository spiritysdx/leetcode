// https://leetcode.cn/problems/3sum-closest/
func threeSumClosest(nums []int, target int) int {
	// 排序数组
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2] // 初始化为前三个数的和，存储与target最接近的和
	size := len(nums)
	for i := 0; i < size-2; i++ {
		left := i + 1
		right := size - 1
		for left < right {
			s := nums[i] + nums[left] + nums[right]
			// 如果当前和更接近目标值，则更新结果
			if abs(target-s) < abs(target-res) {
				res = s
			}
			// 根据与目标值的比较调整左右指针
			if s < target {
				left++
			} else if s > target {
				right--
			} else {
				// 如果正好等于目标值，直接返回
				return s
			}
		}
	}
	return res
}

// 返回较小值的辅助函数
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 取绝对值的辅助函数
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
