func crackPassword(password []int) string {
	// 转换为字符串数组
	strs := make([]string, len(password))
	for i, num := range password {
		strs[i] = strconv.Itoa(num)
	}
	// 自定义排序规则
	sort.Slice(strs, func(i, j int) bool {
		// 比较两种拼接方式的大小
		return strs[i]+strs[j] < strs[j]+strs[i]
	})
	// 拼接结果
	result := strings.Join(strs, "")
	return result
}
