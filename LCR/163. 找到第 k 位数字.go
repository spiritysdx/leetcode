func findKthNumber(k int) int {
	// 1. 确定第 k 位属于几位数 - digit
	digit := 1                       // 当前数字的位数，从 1 位数开始
	count := 9                       // 1 位数的总个数（1 到 9 有 9 个数字）
	for k > digit*count {            // 如果 k 超过当前位数的所有数字总位数
		k -= digit * count          // 减去当前位数数字的所有位数
		digit++                     // 增加位数，进入下一位数范围
		count *= 10                 // 当前位数的数字个数乘以 10，例如 2 位数有 90 个
	}
	// 2. 确定第 k 位属于哪个具体数字
    // 先找位数对应区间起始的数，消耗掉部分K
    // 然后根据剩下的K(前面算k有几位数的时候已经减去了区间外的值了)，找它对应的是当前范围内的第几个数字
	start := int(math.Pow(10, float64(digit-1))) // 当前位数的起始数字，例如 2 位数的起始数字是 10
	num := start + (k-1)/digit                  // 找到第 k 位所在的数字
	// 3. 确定第 k 位是这个数字的第几位
	pos := (k - 1) % digit                      // 计算 k 是这个数字的第几位
	return int(fmt.Sprintf("%d", num)[pos] - '0') // 提取该数字的第 pos 位并返回
}
