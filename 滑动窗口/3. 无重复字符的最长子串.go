// https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
func lengthOfLongestSubstring(s string) int {
    ans := 0 // 保存最长无重复子串的长度
    count := map[byte]int{} // 字节计数器，记录每个字符出现的次数
    sBytes := []byte(s) // 将字符串转换为字节切片以便处理
    start := 0 // 当前无重复子串的起始位置

    for index, x := range sBytes { // 遍历字符串的每个字符
        if _, ok := count[x]; ok { // 如果字符已存在于计数器中
            count[x] += 1 // 更新该字符的计数
            for count[x] > 1 { // 如果该字符的计数大于1，说明有重复
                count[sBytes[start]] -= 1 // 缩小窗口，减少起始位置字符的计数
                start += 1 // 移动起始位置
            }
        } else {
            count[x] = 1 // 如果字符不存在于计数器中，添加该字符并设置计数为1
        }
        ans = max(ans, index-start+1) // 更新最长无重复子串的长度
    }
    return ans // 返回结果
}

func max(a, b int) int { // 辅助函数，用于获取两个整数的最大值
    if a > b {
        return a
    }
    return b
}
