func dismantlingAction(arr string) int {
    res := 0 // 存储最长的连续子串的长度
    arrs := []byte(arr)
    size := len(arrs)
    if size == 0 {
        return 0
    }
    tempMap := map[byte]int{} // 存储字符的最新索引
    start := 0 // 当前无重复子串的起始索引
    for i := 0; i < size; i++ {
        if index, ok := tempMap[arrs[i]]; ok && index >= start {
            // 如果字符已存在且在当前子串范围内，更新起始位置
            start = index + 1
        }
        tempMap[arrs[i]] = i // 更新字符的最新索引
        res = max(res, i-start+1) // 更新最长子串的长度
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
