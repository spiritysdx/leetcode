// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/description/
func twoSum(numbers []int, target int) []int {
    left := 0
    right := len(numbers)-1
    for left < right {
        if numbers[left] + numbers[right] == target {
            return []int{left+1, right+1}
        }
        if numbers[left] + numbers[right] < target {
            left++
        } else if numbers[left] + numbers[right] > target {
            right--
        }
    }
    return []int{}
}
