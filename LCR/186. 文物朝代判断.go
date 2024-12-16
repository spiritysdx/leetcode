func checkDynasty(places []int) bool {
	// 排序数组
	sort.Ints(places)
	// 统计 0 的个数
	temp := 0
	for _, v := range places {
		if v == 0 {
			temp++
		}
	}
	// 检查间隙是否可以由 0 填补，同时检查是否有重复的非零数字
	pre := -1
	for _, v := range places {
		if v == 0 {
			continue
		}
		if pre == -1 {
			pre = v
		} else {
			if v == pre { // 存在重复的非零数字，直接返回 false
				return false
			}
			bias := v - pre - 1
			if bias > temp { // 不足以填补间隙
				return false
			}
			temp -= bias
			pre = v
		}
	}
	return true
}
