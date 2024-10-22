func isAnagram(s string, t string) bool {
    if len(s) != len(t) || s == t {
        return false
    }
    // 使用map存储字符出现的次数
    temp := map[rune]int{}
    // 统计s中每个字符的出现次数
    for _, char := range s {
        temp[char]++
    }
    // 遍历t，减少对应字符的出现次数
    for _, char := range t {
        temp[char]--
        // 如果某个字符的次数小于0，说明t中的该字符比s中的多
        if temp[char] < 0 {
            return false
        }
    }
    // 检查map中的所有字符是否次数为0
    for _, count := range temp {
        if count != 0 {
            return false
        }
    }
    return true
}
