// https://leetcode.cn/problems/group-anagrams/?envType=study-plan-v2&envId=top-100-liked
// 关键点，异位词的顺序位置的形式一致
func groupAnagrams(strs []string) [][]string {
    // 创建 map 用于存储分组，key 是排序后的字符串，value 是原始字符串的数组
    groups := make(map[string][]string)
    for _, str := range strs {
        // 将字符串转换为字节数组并排序
        bytes := []byte(str)
        sort.Slice(bytes, func(i, j int) bool {
            return bytes[i] < bytes[j]
        })
        // 将排序后的字节数组转回字符串，作为 key
        sortedStr := string(bytes)
        // 将原始字符串添加到对应的组中
        groups[sortedStr] = append(groups[sortedStr], str)
    }
    // 将 map 转换为需要的返回格式 [][]string
    result := make([][]string, 0, len(groups))
    for _, group := range groups {
        result = append(result, group)
    }
    return result
}
