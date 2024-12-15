func searchRange(nums []int, target int) []int {
    // 找等于target的(排除掉小于target的)
    left := 0
    right := len(nums)
    var middle int
    for left < right {
        middle = left + (right-left)/2
        if nums[middle] < target {
            left = middle+1
        } else {
            right = middle
        }
    }
    start := left
    // 找大于target的(排除掉小于等于target的)
    left = 0
    right = len(nums)
    for left < right {
        middle = left + (right-left)/2
        if nums[middle] <= target {
            left = middle + 1
        } else {
            right = middle
        }
    }
    end := left
    if start == len(nums) || start == end { // 都大于target或都小于target或大小等于一致
        return []int{-1,-1}
    }
    return []int{start, end-1}
}
