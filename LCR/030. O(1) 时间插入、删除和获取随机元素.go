type RandomizedSet struct {
    values map[int]int // 值到索引的映射
    nums   []int       // 存储值的切片
}

/** 初始化数据结构 */
func Constructor() RandomizedSet {
    rand.Seed(time.Now().UnixNano()) // 为随机数生成器设置种子
    return RandomizedSet{
        values: make(map[int]int),
        nums:   []int{},
    }
}

/** 向集合中插入一个值。如果集合中没有该元素，则返回 true。 */
func (this *RandomizedSet) Insert(val int) bool {
    if _, exists := this.values[val]; exists {
        return false // 值已存在
    }
    // 将值添加到切片和映射中
    this.values[val] = len(this.nums)
    this.nums = append(this.nums, val)
    return true
}

/** 从集合中移除一个值。如果集合中包含该元素，则返回 true。 */
func (this *RandomizedSet) Remove(val int) bool {
    index, exists := this.values[val]
    if !exists {
        return false // 值不存在
    }
    // 用最后一个元素替代被移除元素的索引
    lastElement := this.nums[len(this.nums)-1]
    this.nums[index] = lastElement
    this.values[lastElement] = index // 更新最后一个元素的索引
    // 从切片中删除最后一个元素
    this.nums = this.nums[:len(this.nums)-1]
    delete(this.values, val) // 从映射中移除该值
    return true
}

/**
删除元素 Remove(val int)

    检查是否存在: 首先，检查 values 中是否存在要删除的值。如果不存在，返回 false，表示删除失败。

    删除过程:
        获取要删除值的索引。
        使用 nums 切片中的最后一个元素替换要删除元素的位置。这一步是为了保持切片的连续性（因为删除元素会导致切片中间出现空缺）。
        更新最后一个元素在 values 中的索引，以反映它的新位置。
        从 nums 切片中删除最后一个元素（即缩短切片）。
        从 values 映射中删除该值。

    返回结果: 返回 true，表示删除成功。
*/

/** 从集合中获取一个随机元素。 */
func (this *RandomizedSet) GetRandom() int {
    return this.nums[rand.Intn(len(this.nums))] // 返回一个随机元素
}

/**
 * 你的 RandomizedSet 对象将被实例化并按如下方式调用：
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
