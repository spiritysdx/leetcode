// https://leetcode.cn/problems/h-index-ii/description/
func hIndex(citations []int) int {
    left := 0 // 左指针初始化为数组的起始位置
    size := len(citations) // 获取数组的长度
    right := size // 右指针初始化为数组的长度
    for left < right { // 当左指针小于右指针时继续循环
        mid := left + (right-left)/2 // 计算中间位置，避免溢出
        if citations[mid] < size-mid { // 如果当前引用数小于从中间到末尾的论文数
            left = mid + 1 // 移动左指针到中间位置右侧
        } else {
            right = mid // 移动右指针到中间位置
        }
    }
    return size - left // 返回符合 h 指数定义的值
}
