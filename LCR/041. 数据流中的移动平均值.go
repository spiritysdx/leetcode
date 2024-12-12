type MovingAverage struct {
    size, sum int
    q []int
}


/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
    return MovingAverage{size: size}
}


// 如果数量未达到size，那么直接添加进数组，并且count++， sum+=val
// 如果数量已达到size，那么添加到尾部的时候（覆盖掉头部的值），sum-=头部，count不变
// 最后返回sum/count。
func (this *MovingAverage) Next(val int) float64 {
    if len(this.q) == this.size {
        this.sum -= this.q[0]
        this.q = this.q[1:]
    }
    this.sum += val
    this.q = append(this.q, val)
    return float64(this.sum) / float64(len(this.q))
}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
