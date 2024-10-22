func groupAnagrams(strs []string) [][]string {
    // 创建一个映射，用于存储排序后的字符串和其对应的原始字符串数组
    mp := map[string][]string{}
    // 遍历输入的字符串数组
    for _, str := range strs {
        // 将字符串转换为字节数组
        s := []byte(str)
        // 对字节数组进行排序
        sort.Slice(s, func(i, j int) bool { return s[i] < s[j] }) 
        /**
        通过这行代码，字节切片 s 中的字符会根据字典顺序进行排序。例如：
        输入："bat" → []byte{'b', 'a', 't'}
        排序后：[]byte{'a', 'b', 't'}
        最终转换为字符串："abt"
        **/
        // 将排序后的字节数组转换回字符串
        sortedStr := string(s)
        // 将原始字符串添加到对应排序后的字符串的切片中
        mp[sortedStr] = append(mp[sortedStr], str)
    }
    // 创建一个切片用于存储结果，初始容量为映射的长度
    ans := make([][]string, 0, len(mp))
    // 遍历映射的值，将其添加到结果切片中
    for _, v := range mp {
        ans = append(ans, v)
    }
    // 返回结果切片
    return ans
}

/**
func groupAnagrams(strs []string) [][]string {
    // 创建一个映射，用于存储字符计数和对应的原始字符串数组
    mp := map[[26]int][]string{}

    // 遍历输入的字符串数组
    for _, str := range strs {
        // 创建一个长度为 26 的数组，用于计数每个字母出现的次数
        cnt := [26]int{}
        
        // 遍历字符串中的每个字符
        for _, b := range str {
            // 计算字母的索引并增加对应计数
            cnt[b-'a']++
        }

        // 将原始字符串添加到对应字符计数的切片中
        mp[cnt] = append(mp[cnt], str)
    }

    // 创建一个切片用于存储结果，初始容量为映射的长度
    ans := make([][]string, 0, len(mp))

    // 遍历映射的值，将其添加到结果切片中
    for _, v := range mp {
        ans = append(ans, v)
    }

    // 返回结果切片
    return ans
}

**/
