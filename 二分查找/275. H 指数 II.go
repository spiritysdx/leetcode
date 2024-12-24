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

/**
在二分查找的过程中：

    如果 citations[mid] < size - mid:
        意味着从 mid 开始到数组末尾，有 size - mid 篇论文的引用次数比当前值 citations[mid] 更大。
        换句话说，citations[mid] 太小，无法满足 h 指数 的条件。
        因此，我们需要在更右侧寻找可能的答案，所以移动左指针 (left = mid + 1)。

    如果 citations[mid] >= size - mid:
        表示当前 mid 对应的引用次数足够大，可能是一个合法的 h 指数 候选值。
        但我们还需要继续往左侧查找，看看是否存在更小的 mid 也满足条件。
        因此，我们移动右指针 (right = mid)。
**/
