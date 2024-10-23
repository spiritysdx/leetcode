func findMinDifference(timePoints []string) int {
    // 如果时间点超过 1440 个（一天的分钟数），一定有重复，最小差为 0
    if len(timePoints) > 1440 {
        return 0
    }
    // 将时间点转换为分钟数
    minutes := make([]int, len(timePoints))
    for i, time := range timePoints {
        parts := strings.Split(time, ":")
        hours, _ := strconv.Atoi(parts[0])
        mins, _ := strconv.Atoi(parts[1])
        minutes[i] = hours*60 + mins
    }
    // 排序时间点 - 因为题目中要求差值为正数
    sort.Ints(minutes)
    // 初始化最小差值为一天的分钟数
    minDiff := 1440
    // 计算相邻时间点的最小差值
    for i := 1; i < len(minutes); i++ {
        minDiff = min(minDiff, minutes[i]-minutes[i-1])
    }
    // 比较首尾时间的差值，处理跨天的情况 - 要考虑循环时间
    minDiff = min(minDiff, 1440+minutes[0]-minutes[len(minutes)-1])
    return minDiff
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
