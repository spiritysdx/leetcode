// https://leetcode.cn/problems/successful-pairs-of-spells-and-potions/
func successfulPairs(spells []int, potions []int, success int64) []int {
    sizeA, sizeB := len(spells), len(potions)
    res := []int{}
    if sizeA == 0 || sizeB == 0 {
        return res
    }
    slices.Sort(potions)
    for _, k := range spells {
        res = append(res, calculate(potions, int64(k), success))
    }
    return res
}

func calculate(temp []int, k int64, target int64) int {
    left := 0
    right := len(temp)
    for left < right {
        middle := left + (right-left)/2
        if k*int64(temp[middle]) < target {
            left = middle+1
        } else {
            right = middle
        }
    }
    start := left // 从这开始后面的都可以
    if start == len(temp) {
        return 0
    }
    return len(temp) - start
}
