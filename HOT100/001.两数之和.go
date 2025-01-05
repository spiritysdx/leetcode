// https://leetcode.cn/problems/two-sum/description/
func twoSum(nums []int, target int) []int {
    // 需要先进行数组排序 - 冒泡排序
    sorted := bubbleSort(nums)
    // 由于后续需要返回位置，需要使用哈希表记录每个值对应原有的数组的位置
    numMap := make(map[int][]int) // 键-初始值，值-所有初始值所在数组的多个位置
    for i := 0; i < len(nums); i++ {
        numMap[nums[i]] = append(numMap[nums[i]], i)
    }
    var ans []int
    // 需要对有序号数组进行相向双指针查找目标
    for i, j := 0, len(sorted)-1; i < j; {
        sum := sorted[i] + sorted[j]
        if sum == target {
            ans = append(ans, sorted[i], sorted[j])
            break
        } else if sum < target {
            i++
        } else {
            j--
        }
    }
    // 如果找到答案，通过哈希表返回原始索引
    if len(ans) == 2 {
        index1 := numMap[ans[0]][0]
        index2 := numMap[ans[1]][0]
        // 如果两个数相同，确保获取不同的索引
        if ans[0] == ans[1] && len(numMap[ans[0]]) > 1 {
            index2 = numMap[ans[1]][1]
        }
        return []int{index1, index2}
    }
    return []int{}
}

func bubbleSort(nums []int) []int {
    n := len(nums)
    // 对原数组进行复制，防止排序修改原数组
    sorted := make([]int, n)
    copy(sorted, nums)
    // 进行冒泡排序
    for i := 0; i < n-1; i++ {
        swapped := false
        for j := 0; j < n-i-1; j++ {
            if sorted[j] > sorted[j+1] {
                // 交换相邻元素
                sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
                swapped = true
            }
        }
        // 如果没有发生交换，说明已经排序完成
        if !swapped {
            break
        }
    }
    return sorted
}

// 哈希表存储中间检索过的值，下次需要用的时候检索一下
func twoSum(nums []int, target int) []int {
    tempMap := map[int]int{} // 键值对： 检索过的值，位置
    for index, value := range nums {
        if v, ok := tempMap[target-value]; ok {
            return []int{index, v}
        }
        tempMap[value] = index
    }
    return nil
}
