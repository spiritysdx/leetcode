type Checkout struct {
	maxPrice int       // 当前队列的最大值
	size     int       // 队列的大小
	p        []int     // 用于存储队列的值
	maxDeque []int     // 双端队列，维护最大值
}

func Constructor() Checkout {
	return Checkout{
		maxPrice: -1,       // 初始化最大值为-1（表示空队列）
		size:     0,        // 初始化队列大小为0
		p:        []int{},  // 初始化空队列
		maxDeque: []int{},  // 初始化空的双端队列
	}
}

func (this *Checkout) Get_max() int {
	if this.size == 0 {
		return -1 // 如果队列为空，返回-1
	}
	return this.maxPrice // 返回当前的最大值
}

func (this *Checkout) Add(value int) {
	this.p = append(this.p, value) // 将新值添加到队列末尾
	this.size++                    // 队列大小增加
	// 更新最大值队列
	// 在向队列中添加新元素时，我们需要保证双端队列中的元素是单调递减的，以确保队列头部始终是最大值。
	// 当新元素比双端队列尾部的元素更大时：
	// 尾部元素已经不可能成为当前或未来的最大值。
	// 移除这些尾部元素可以避免冗余，保持双端队列的有效性和简洁性。
	for len(this.maxDeque) > 0 && this.maxDeque[len(this.maxDeque)-1] < value {
		this.maxDeque = this.maxDeque[:len(this.maxDeque)-1] // 移除双端队列尾部的较小元素
	}
	this.maxDeque = append(this.maxDeque, value) // 将新值添加到双端队列尾部
	this.maxPrice = this.maxDeque[0]            // 更新当前最大值
}

func (this *Checkout) Remove() int {
	if this.size == 0 {
		return -1 // 如果队列为空，返回-1
	}
	// 移除队列首部元素
	removed := this.p[0] // 获取队列首部元素
	this.p = this.p[1:]  // 更新队列，移除首部元素
	this.size--          // 队列大小减少
	// 更新最大值队列
	if len(this.maxDeque) > 0 && this.maxDeque[0] == removed {
		this.maxDeque = this.maxDeque[1:] // 如果移除的元素是当前最大值，更新双端队列
	}
	if this.size == 0 {
		this.maxPrice = -1 // 队列为空，最大值重置为-1
	} else {
		this.maxPrice = this.maxDeque[0] // 更新当前最大值
	}
	return removed // 返回被移除的元素
}

/**
 * Your Checkout object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get_max();
 * obj.Add(value);
 * param_3 := obj.Remove();
 */
