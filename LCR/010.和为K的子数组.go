func subarraySum(nums []int, k int) int {
	size := len(nums)
	ans := 0
	var start, end, sum int
	for ; start < size; start++ {
		sum = 0
		for end = start; end>=0; end-- {
			sum += nums[end]
			if sum == k {
				ans++
			}
		}
	}
	return ans
}
