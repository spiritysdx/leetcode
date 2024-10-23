func isAlienSorted(words []string, order string) bool {
    // 构建字典映射，记录每个字母在字母表中的顺序
    sMap := map[byte]int{}
    for i := 0; i < len(order); i++ {
        sMap[order[i]] = i
    }
    // 比较两个单词的字典序
    compare := func(word1, word2 string) bool {
        minLength := min(len(word1), len(word2))
        for i := 0; i < minLength; i++ {
            // 比较相同位置的字母
            if sMap[word1[i]] < sMap[word2[i]] {
                return true
            } else if sMap[word1[i]] > sMap[word2[i]] {
                return false
            }
        }
        // 如果前缀相同，较短的单词应该排在前面
        return len(word1) <= len(word2)
    }
    // 检查每对相邻的单词是否符合字典序
    for i := 1; i < len(words); i++ {
        if !compare(words[i-1], words[i]) {
            return false
        }
    }
    return true
}

// 辅助函数：返回两个数中的最小值
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
