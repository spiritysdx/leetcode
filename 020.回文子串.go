func countSubstrings(s string) int {
    size := len(s)
    if size == 0 {
        return 0
    }
    ans := 0
    // 判断回文的函数
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
    for tmpSize := 1; tmpSize <= size; tmpSize++ {
        for i, j := 0, tmpSize-1; j < size; i, j = i+1, j+1 {
            if isPalindrome(s, i, j) {
                ans++
            }
        }
    }
    return ans
}
