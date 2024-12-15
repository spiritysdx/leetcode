func reverseMessage(message string) string {
    // 去掉首尾空格，并按空格分割单词
    words := strings.Fields(message)
    // 反转切片中的单词顺序
    for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
        words[i], words[j] = words[j], words[i]
    }
    // 将单词用空格连接并返回
    return strings.Join(words, " ")
}
