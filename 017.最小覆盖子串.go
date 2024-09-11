// 这个不行，只能做序列的，题目要求做组合的
// func minWindow(s string, t string) string {
// 	s1Byte, s2Byte := []byte(t), []byte(s)
// 	s1Size, s2Size := len(s1Byte), len(s2Byte)
// 	if s1Size > s2Size {
// 		return ""
// 	}
// 	ans1, ans2 := []byte{}, []byte{}
// 	var status bool
// 	var i, j int
// 	for j < s2Size {
// 		for ; i < s1Size; {
// 			for ; j < s2Size; {
// 				if s1Byte[i] == s2Byte[j] {
// 					status = true
// 					ans2 = append(ans2, s2Byte[j])
// 					break
// 				} else if status {
// 					ans2 = append(ans2, s2Byte[j])
// 				}
//                 // 1.需要处理i没匹配全但j全部用完的情况
//                 if j == s2Size-1 && i != s1Size-1 && s1Byte[i] != s2Byte[j] {
//                     j = s2Size
//                     ans2 = []byte{}
//                 }
//                 j++
// 			}
//             // 2.需要处理i刚好匹配完，但j剩下的部分不足以继续匹配新的一轮i的情况
//             if s2Size-j < s1Size {
//                 j = s2Size
//             }
//             i++
// 		}
//         if len(ans1) == 0 && len(ans2) > 0 && len(ans2) < len(ans1) {
//             ans1 = ans2
//         }
//         ans2 = []byte{}
// 	}
// 	if len(ans1) > 0 {
//         return string(ans1)
// 	}
//     return ""
// }

func minWindow(s string, t string) string {
    ori, cnt := map[byte]int{}, map[byte]int{}  // 定义两个map，ori记录t中各字符出现的次数，cnt记录当前窗口中各字符出现的次数
    for i := 0; i < len(t); i++ {  // 初始化ori
        ori[t[i]]++
    }

    sLen := len(s)  // 字符串s的长度
    len := math.MaxInt32  // 最小窗口长度，初始值设为一个很大的数
    ansL, ansR := -1, -1  // 记录最小窗口的左右边界

    check := func() bool {  // 检查当前窗口是否包含t中的所有字符
        for k, v := range ori {  // 遍历t中所有字符及其出现次数
            if cnt[k] < v {  // 如果当前窗口中某字符出现的次数小于在t中出现的次数
                return false  // 返回false，表示当前窗口不符合要求
            }
        }
        return true  // 返回true，表示当前窗口符合要求
    }
    for l, r := 0, 0; r < sLen; r++ {  // 使用双指针法，r为右指针，l为左指针
        if r < sLen && ori[s[r]] > 0 {  // 如果s[r]在t中出现过
            cnt[s[r]]++  // 记录s[r]在当前窗口中出现的次数
        }
        for check() && l <= r {  // 如果当前窗口符合要求，且左指针不超过右指针
            if (r - l + 1 < len) {  // 如果当前窗口长度小于已记录的最小窗口长度
                len = r - l + 1  // 更新最小窗口长度
                ansL, ansR = l, l + len  // 更新最小窗口的左右边界
            }
            if _, ok := ori[s[l]]; ok {  // 如果s[l]在t中出现过
                cnt[s[l]] -= 1  // 减少当前窗口中s[l]的出现次数
            }
            l++  // 移动左指针，尝试缩小窗口
        }
    }
    if ansL == -1 {  // 如果没有找到符合条件的窗口
        return ""  // 返回空字符串
    }
    return s[ansL:ansR]  // 返回最小窗口子字符串
}
