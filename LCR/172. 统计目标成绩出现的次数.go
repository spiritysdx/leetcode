func countTarget(scores []int, target int) int {
    // 找到第一个等于 target 的位置
    left := 0
    right := len(scores)
    for left < right {
        mid := left + (right-left)/2
        if scores[mid] < target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    start := left

    // 找到第一个大于 target 的位置
    left = 0
    right = len(scores)
    for left < right {
        mid := left + (right-left)/2
        if scores[mid] <= target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    end := left

    // 计算出现次数
    return end - start
}
