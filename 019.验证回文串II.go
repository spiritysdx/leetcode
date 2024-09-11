func validPalindrome(s string) bool {
    size := len(s)  // 获取字符串的长度
    if size == 0 || size == 1 {  // 如果字符串长度为0或1，直接返回true
        return true
    }
    // 判断是否是回文串的辅助函数
    isPalindrome := func(str string, left, right int) bool {
        for left < right {
            if str[left] != str[right] {
                return false
            }
            left++
            right--
        }
        return true
    }
    left, right := 0, size-1  // 初始化左右指针
    for left < right {  // 当左指针小于右指针时，继续循环
        if s[left] == s[right] {  // 如果左右字符相等，继续向中间移动指针
            left++
            right--
        } else {  // 如果字符不相等
            // 尝试跳过左指针的字符或右指针的字符，看是否能成为回文
            return isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
        }
    }
    return true  // 如果循环完成，说明字符串可以通过最多删除一个字符成为回文
}
