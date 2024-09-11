func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	var temp int
	for left < right {
		temp = numbers[left] + numbers[right]
		if target == temp {
			return []int{left, right}
		} else if temp > target {
			right--
		} else if temp < target {
			left++
		}
	}
	return []int{}
}
