func fileCombination(target int) [][]int {
	left := 1 // 左边界，初始值为1
	right := 2 // 右边界，初始值为2
	var res [][]int // 用于存储所有满足条件的连续序列的结果
	for left < right { // 保证左边界始终小于右边界
		sum := (left + right) * (right - left + 1) / 2 // 计算从left到right的连续序列的和
		if sum == target { // 如果当前序列的和等于目标值
			var list []int // 创建一个切片，用于存储当前的连续序列
			for i := left; i <= right; i++ { // 遍历从left到right的所有整数
				list = append(list, i) // 将每个整数加入当前序列
			}
			res = append(res, list) // 将当前序列加入结果集
			left++ // 左边界右移，继续寻找下一个序列
		} else if sum < target { // 如果当前序列的和小于目标值
			right++ // 右边界右移，扩大范围
		} else { // 如果当前序列的和大于目标值
			left++ // 左边界右移，缩小范围
		}
	}
	return res // 返回所有满足条件的连续序列
}
