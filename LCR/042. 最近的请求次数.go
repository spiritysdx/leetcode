type RecentCounter struct {
    size int
    q []int
}


func Constructor() RecentCounter {
    return RecentCounter{size: 0}
}


func (this *RecentCounter) Ping(t int) int {
    this.size += 1
    this.q = append(this.q, t)
    if this.size == 1 {
        return 1
    }
    target := t-3000
    // 二分查找比target大于等于的值开始的地方
    left := 0
    right := this.size
    for left < right {
        mid := left + (right-left)/2
        if this.q[mid] < target {
            left = mid+1
        } else {
            right = mid
        }
    }
    return this.size-left
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
