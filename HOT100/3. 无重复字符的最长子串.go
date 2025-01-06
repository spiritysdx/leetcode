// https://leetcode.cn/problems/longest-substring-without-repeating-characters/?envType=study-plan-v2&envId=top-100-liked
func lengthOfLongestSubstring(s string) int {
    sBytes := []byte(s)
    start := 0 
    ans := 0
    tempMap := map[byte]int{}
    for i, c := range sBytes {
        if p, ok := tempMap[c]; ok {
            start = max(start, p+1)  // 修改点1：确保start只能向前移动
        }
        ans = max(ans, i-start+1)    // 修改点2：每次都更新ans
        tempMap[c] = i
    }
    return ans
}
