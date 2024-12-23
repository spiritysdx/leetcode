// https://leetcode.cn/problems/find-the-longest-semi-repetitive-substring/description/
func longestSemiRepetitiveSubstring(s string) int {
    // 至多有一对相邻字符是相等的
    ans := 0
    start := 0
    prev := -1 // 上一次出现重复字符的位置
    sBytes := []byte(s)
    if len(sBytes) == 1 {
        return 1
    }
    for i := 1; i < len(sBytes); i++ {
        if sBytes[i] == sBytes[i-1] {
            if prev != -1 {
                // 更新 start 到上一次重复的字符位置
                start = prev + 1
            }
            prev = i - 1 // 记录当前重复字符的位置
        }
        // 更新最大长度
        ans = max(ans, i-start+1)
    }
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
