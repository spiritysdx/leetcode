// https://leetcode.cn/problems/count-the-number-of-fair-pairs/
func countFairPairs(nums []int, lower int, upper int) int64 {
	// 对数组进行升序排序
	sort.Ints(nums)
	res := int64(0)
	// 遍历每一个数 nums[i]，用双指针找到满足条件的范围
	for i := 0; i < len(nums); i++ {
		// 左指针找到 >= lower - nums[i] 的位置
		left := lowerBound(nums, i+1, len(nums)-1, lower-nums[i])
		// 右指针找到 <= upper - nums[i] 的位置
		right := upperBound(nums, i+1, len(nums)-1, upper-nums[i])
		if left <= right {
			res += int64(right - left + 1)
		}
	}
	return res
}

// lowerBound 找到数组中第一个 >= target 的索引
func lowerBound(nums []int, start, end, target int) int {
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start
}

// upperBound 找到数组中最后一个 <= target 的索引
func upperBound(nums []int, start, end, target int) int {
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] <= target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return end
}

// func countFairPairs(nums []int, lower, upper int) (ans int64) {
// 	sort.Ints(nums)
// 	for j, x := range nums {
// 		r := sort.SearchInts(nums[:j], upper-x+1) // <= upper-nums[j] 的 nums[i] 的个数
// 		l := sort.SearchInts(nums[:j], lower-x) // < lower-nums[j] 的 nums[i] 的个数
// 		ans += int64(r - l)
// 	}
// 	return
// }
