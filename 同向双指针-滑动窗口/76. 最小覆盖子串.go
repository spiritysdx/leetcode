// https://leetcode.cn/problems/minimum-window-substring/description/
// 判断当前窗口的字符频率是否覆盖目标字符串 t 的字符频率
func isCovered(cntS, cntT []int) bool {
    for i := 'A'; i <= 'Z'; i++ { // 遍历大写字母
        if cntS[i] < cntT[i] { // 如果当前窗口中的字符数量小于目标字符数量
            return false // 不满足覆盖条件
        }
    }
    for i := 'a'; i <= 'z'; i++ { // 遍历小写字母
        if cntS[i] < cntT[i] { // 如果当前窗口中的字符数量小于目标字符数量
            return false // 不满足覆盖条件
        }
    }
    return true // 所有字符都满足条件
}

// 寻找字符串 s 中的最小窗口子串，使得它包含字符串 t 中所有字符
func minWindow(s, t string) string {
    ansLeft, ansRight := -1, len(s) // 初始化最优解的左右端点，默认无解
    var cntS, cntT [128]int // 用于存储 ASCII 范围内的字符频率（包含大小写字母和其他符号）
    
    for _, c := range t { // 遍历字符串 t
        cntT[c]++ // 记录 t 中每个字符的频率
    }

    left := 0 // 初始化滑动窗口的左端点
    for right, c := range s { // 移动滑动窗口的右端点
        cntS[c]++ // 将当前字符加入窗口并更新窗口内字符的频率
        for isCovered(cntS[:], cntT[:]) { // 判断当前窗口是否完全覆盖 t 中的字符
            if right-left < ansRight-ansLeft { // 如果找到更短的满足条件的子串
                ansLeft, ansRight = left, right // 更新最优解的左右端点
            }
            cntS[s[left]]-- // 将窗口左端点的字符移出并更新频率
            left++ // 左端点右移，尝试缩小窗口
        }
    }

    if ansLeft < 0 { // 如果最优解未更新，说明无解
        return "" // 返回空字符串
    }
    return s[ansLeft : ansRight+1] // 返回最优解子串
}
