// 定义一个存储键值对的结构体 entry
type entry struct {
    key, value int
}

// 定义 LRUCache 结构体
type LRUCache struct {
    capacity  int               // 缓存容量
    list      *list.List        // 用于存储缓存内容的双向链表，按照访问顺序存储
    keyToNode map[int]*list.Element // 哈希表，存储键对应链表节点的映射
}

// 构造函数，初始化缓存
func Constructor(capacity int) LRUCache {
    return LRUCache{
        capacity:  capacity,             // 初始化缓存容量
        list:      list.New(),           // 创建一个新的双向链表
        keyToNode: map[int]*list.Element{}, // 初始化键到链表节点的映射
    }
}

// Get 方法，获取缓存中的值
func (c *LRUCache) Get(key int) int {
    node := c.keyToNode[key] // 从哈希表中找到 key 对应的链表节点
    if node == nil {         // 如果没有找到节点，表示缓存中没有这个 key
        return -1            // 返回 -1 表示未命中
    }
    c.list.MoveToFront(node) // 将该节点移到链表的最前面，表示它最近被访问过
    return node.Value.(entry).value // 返回节点存储的值
}

// Put 方法，往缓存中存储键值对
func (c *LRUCache) Put(key, value int) {
    if node := c.keyToNode[key]; node != nil { // 如果缓存中已经存在这个 key
        node.Value = entry{key, value} // 更新该节点的值
        c.list.MoveToFront(node)       // 将该节点移到链表的最前面，表示它最近被访问
        return                         // 直接返回
    }
    
    // 如果缓存中没有这个 key，则需要新增
    c.keyToNode[key] = c.list.PushFront(entry{key, value}) // 将新键值对插入链表最前面，并更新哈希表
    if len(c.keyToNode) > c.capacity { // 如果缓存已满
        // 移除链表最后面的节点（即最久未被访问的节点）
        lastNode := c.list.Back() // 获取链表最后一个节点
        lastEntry := c.list.Remove(lastNode).(entry) // 删除该节点并返回其存储的键值对
        delete(c.keyToNode, lastEntry.key) // 从哈希表中删除该键
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
