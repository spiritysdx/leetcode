// https://leetcode.cn/problems/minimum-time-to-complete-trips/description/
func minimumTime(time []int, totalTrips int) int64 {
    slices.Sort(time)  // 对时间数组进行升序排序
    // step1 找到左右指针的端点值，结果检索的区间找出来
    left := 0  // 左指针初始化为0
    right := time[0]*totalTrips  // 右指针初始化为完成最大旅途所需的时间，即最短时间乘以总旅途数
    for left < right {  // 当左指针小于右指针时，继续进行二分查找
        mid := left + (right-left)/2  // 计算中间值，当前循环检索的是当前mid值的时长下行驶的路程长度
        trips := 0  // 统计在mid时间内可以完成的总旅途数
        for _, t := range time {  // 遍历所有的时间数组元素
            if mid < t {  // 如果当前时间mid小于该元素时间，则跳出循环，因为该时间无法完成任何旅途
                break
            }
            trips += mid/t  // 在当前时间内完成的旅途数
        }
        if trips >= totalTrips {  // 如果在mid时间内可以完成的旅途数大于等于目标旅途数
            right = mid  // 将右指针更新为mid，增大范围
        } else {
            left = mid + 1  // 否则，将左指针更新为mid + 1，缩小范围
        }
    }
    return int64(left)  // 返回左指针的值，即最少时间
}

