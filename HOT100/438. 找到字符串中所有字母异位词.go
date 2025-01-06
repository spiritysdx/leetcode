// https://leetcode.cn/problems/find-all-anagrams-in-a-string/?envType=study-plan-v2&envId=top-100-liked
func findAnagrams(s string, p string) []int {
    countS := [26]byte{}
    countP := [26]byte{}
    for _, c := range p {
        countP[c-'a'] += 1
    }
    ans := []int{}
    size := len(p)
    for right, c := range s {
        countS[c-'a'] += 1
        left := right - size + 1
        if left < 0 {
            continue
        }
        if countS == countP {
            ans = append(ans, left)
        }
        countS[s[left]-'a'] -= 1
    }
    return ans
}
