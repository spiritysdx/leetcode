type KthLargest struct {
    size int
    maxk int
    q []int
}

func Constructor(k int, nums []int) KthLargest {
    sort.Ints(nums)
    return KthLargest{maxk: k, q: nums, size: len(nums)}
}

func (this *KthLargest) Add(val int) int {
    this.q = append(this.q, val)
    this.size += 1
    sort.Ints(this.q)
    return this.q[this.size-this.maxk]
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
