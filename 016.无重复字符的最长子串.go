func lengthOfLongestSubstring(s string) int {
    n := len(s)
    if n == 0 {
        return 0 // 如果字符串长度为0，直接返回0
    }
    // 使用map来记录字符及其最新位置
    charIndex := make(map[byte]int)
    maxLen := 0
    left := 0 // 窗口的左边界
    for right := 0; right < n; right++ { // 窗口的右边界从0开始遍历
        if index, found := charIndex[s[right]]; found && index >= left {
            // 如果当前字符已经在map中存在且位置在left右边
            left = index + 1 // 移动left到重复字符的下一个位置
        }
        charIndex[s[right]] = right // 更新当前字符的位置到map中
        if currentLen := right - left + 1; currentLen > maxLen {
            // 计算当前窗口的长度，并更新maxLen
            maxLen = currentLen
        }
    }
    return maxLen // 返回最长无重复字符子串的长度
}
