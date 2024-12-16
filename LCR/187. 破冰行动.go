func iceBreakingGame(num int, target int) int {
    // (当前index + target) % 上一轮剩余数字的个数
    // 约瑟夫问题
    idx := 0
    for i:=2; i<=num; i++ {
        idx = (idx+target)%i
    }
    return idx
}
