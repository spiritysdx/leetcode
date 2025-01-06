// https://leetcode.cn/problems/longest-consecutive-sequence/?envType=study-plan-v2&envId=top-100-liked
func longestConsecutive(nums []int) int {
    // 创建哈希表存储数字
    tempMap := map[int]int{}
    for _, c := range nums {
        tempMap[c] = 1
    }
    maxLength := 0
    // 遍历哈希表中的每个数字
    for num := range tempMap {
        // 如果当前数字是序列的起点（即num-1不存在于哈希表中）
        if tempMap[num-1] == 0 {
            currentNum := num
            currentLength := 1
            // 往后查找连续的数字
            for tempMap[currentNum+1] > 0 {
                currentNum++
                currentLength++
            }
            // 更新最长序列长度
            if currentLength > maxLength {
                maxLength = currentLength
            }
        }
    }
    return maxLength
}
