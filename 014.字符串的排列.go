// func checkInclusion(s1 string, s2 string) bool {
//     s1Bytes := []byte(s1) // 将s1转换为字节切片
//     s2Bytes := []byte(s2) // 将s2转换为字节切片
//     s1Len := len(s1Bytes) // 获取s1的长度
//     s2Len := len(s2Bytes) // 获取s2的长度
//     if s1Len > s2Len { // 如果s1比s2长，直接返回false
//         return false
//     }
//     // 初始化频率计数数组，大小为26，分别记录s1和s2中前s1Len长度的字符频率
//     s1Count := make([]int, 26)
//     s2Count := make([]int, 26)
//     for i := 0; i < s1Len; i++ {
//         s1Count[s1Bytes[i]-'a']++ // 计算s1中每个字符的频率
//         s2Count[s2Bytes[i]-'a']++ // 计算s2中前s1Len长度的字符频率
//     }
//     matches := 0 // 初始化匹配计数
//     for i := 0; i < 26; i++ {
//         if s1Count[i] == s2Count[i] { // 如果字符频率相等，增加匹配计数
//             matches++
//         }
//     }
//     // 通过滑动窗口检查 s2 中所有长度为 s1Len 的子字符串
//     for i := 0; i < s2Len-s1Len; i++ {
//         if matches == 26 { // 如果所有字符频率匹配，返回true
//             return true
//         }
//         leftChar := s2Bytes[i] - 'a' // 滑动窗口移出的字符
//         rightChar := s2Bytes[i+s1Len] - 'a' // 滑动窗口加入的字符
//         s2Count[rightChar]++ // 增加滑动窗口加入的字符频率
//         if s2Count[rightChar] == s1Count[rightChar] { // 如果频率相等，增加匹配计数
//             matches++
//         } else if s2Count[rightChar] == s1Count[rightChar]+1 { // 如果频率多1，减少匹配计数
//             matches--
//         }
//         s2Count[leftChar]-- // 减少滑动窗口移出的字符频率
//         if s2Count[leftChar] == s1Count[leftChar] { // 如果频率相等，增加匹配计数
//             matches++
//         } else if s2Count[leftChar] == s1Count[leftChar]-1 { // 如果频率少1，减少匹配计数
//             matches--
//         }
//     }
//     return matches == 26 // 最后检查是否所有字符频率匹配
// }



func checkInclusion(s1 string, s2 string) bool {
    s1Bytes, s2Bytes := []byte(s1), []byte(s2)
    s1Size, s2Size := len(s1Bytes), len(s2Bytes)
    if s1Size > s2Size {
        return false
    }
    var i int
    s1Count, s2Count, matches := [26]int{}, [26]int{}, 0
    for ;i < s1Size; i++{
        s1Count[s1Bytes[i]-'a']++
        s2Count[s2Bytes[i]-'a']++
    }
    for i=0; i < 26; i++{
        if s1Count[i] == s2Count[i] {
            matches++
        }
    }
    var leftChar,rightChar byte
    for i=0; i<s2Size-s1Size;i++{
        if matches == 26 {
            return true
        }
        leftChar = s2Bytes[i]-'a' // 移出
        rightChar = s2Bytes[i+s1Size]-'a' // 加入
        s2Count[rightChar]++
        if s2Count[rightChar] == s1Count[rightChar] {
            matches++
        } else if s2Count[rightChar] == s1Count[rightChar]+1{
            matches--
        }
        s2Count[leftChar]--
        if s2Count[leftChar] == s1Count[leftChar] {
            matches++
        } else if s2Count[leftChar] == s1Count[leftChar]-1{
            matches--
        }
    }
    return matches == 26
}
