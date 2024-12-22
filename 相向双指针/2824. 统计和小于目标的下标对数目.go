// https://leetcode.cn/problems/count-pairs-whose-sum-is-less-than-target/
func countPairs(nums []int, target int) (ans int) {
	slices.Sort(nums)
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] < target {
			ans += right - left
			left++
		} else {
			right--
		}
	}
	return
}
