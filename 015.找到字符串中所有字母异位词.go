func findAnagrams(s string, p string) []int {
    s1Byte, s2Byte := []byte(p), []byte(s)
    s1Size, s2Size := len(s1Byte), len(s2Byte)
    if s1Size > s2Size {
        return []int{}
    }

    // 字符计数器
    var i, matches int
    res := []int{}
    s1Count, s2Count := [26]int{}, [26]int{}

    // 初始化字符计数器
    for i = 0; i < s1Size; i++ {
        s1Count[s1Byte[i] - 'a']++
        s2Count[s2Byte[i] - 'a']++
    }

    // 计算初始的 matches
    for i = 0; i < 26; i++ {
        if s1Count[i] == s2Count[i] {
            matches++
        }
    }

    // 滑动窗口
    for i = 0; i < s2Size - s1Size; i++ {
        if matches == 26 {
            res = append(res, i)
        }

        leftChar := s2Byte[i] - 'a'
        rightChar := s2Byte[i + s1Size] - 'a'

        // 更新右边界字符计数
        s2Count[rightChar]++
        if s2Count[rightChar] == s1Count[rightChar] {
            matches++
        } else if s2Count[rightChar] == s1Count[rightChar] + 1 {
            matches--
        }

        // 更新左边界字符计数
        s2Count[leftChar]--
        if s2Count[leftChar] == s1Count[leftChar] {
            matches++
        } else if s2Count[leftChar] == s1Count[leftChar] - 1 {
            matches--
        }
    }
    // 检查最后一个窗口
    if matches == 26 {
        res = append(res, s2Size - s1Size)
    }
    return res
}
