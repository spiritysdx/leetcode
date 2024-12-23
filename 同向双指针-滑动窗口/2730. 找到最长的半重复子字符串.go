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
        // 更新最大长度，此时需要保证这之前的start到index至多有一次重复字符
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

// func longestSemiRepetitiveSubstring(s string) int {
// 	ans, left, same := 1, 0, 0
// 	for right := 1; right < len(s); right++ {
// 		if s[right] == s[right-1] {
// 			same++
// 		}
//         if same > 1 { // same == 2
//             left++
//             for s[left] != s[left-1] {
//                 left++
//             }
//             same = 1
//         }
// 		ans = max(ans, right-left+1)
// 	}
// 	return ans
// }
// 移动右指针 right，并统计相邻相同的情况出现了多少次，记作 same。
// 如果 same>1，则不断移动左指针 left 直到 s[left]=s[left−1]，此时将一对相同的字符移到窗口之外。然后将 same 置为 1。
// 然后统计子串长度 right−left+1 的最大值。
// 思想：保证检索之前区间内无重复字符，遇到重复字符就排空到无重复字符
