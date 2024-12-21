
func sockCollocation(sockets []int) []int {
    // 题目要求时间复杂度 O(N) ，空间复杂度 O(1) ，因此首先排除 暴力法 和 哈希表统计法 。
	t := 0 // t 为 两个不相等值的异或值
	for _, v := range sockets {
		t ^= v
	}
	m := 1 // m 为第一个为 1 的那一位
	for m&t == 0 {
		m <<= 1
	}
	x, y := 0, 0
	for _, v := range sockets {
		if v&m == 0 { //如果此数字这一位 为 0，就和 x 异或
			x ^= v
		} else {
			//如果此数字这一位 为 1，就和 y 异或
			y ^= v
		}
	}
	// 正是因为 x 和 y 不同，才导致了 t 的产生
	return []int{x, y}
}
