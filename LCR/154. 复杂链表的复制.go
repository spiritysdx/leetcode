/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }
    // 创建哈希表，用于存储原节点和复制节点的对应关系
    nodeMap := make(map[*Node]*Node)
    cur := head
    // 第一次遍历：复制每个节点，并将原节点和复制节点存入哈希表
    for cur != nil {
        nodeMap[cur] = &Node{Val: cur.Val}
        cur = cur.Next
    }
    cur = head
    // 第二次遍历：分配 `Next` 和 `Random` 指针
    for cur != nil {
        if cur.Next != nil {
            nodeMap[cur].Next = nodeMap[cur.Next]
        }
        if cur.Random != nil {
            nodeMap[cur].Random = nodeMap[cur.Random]
        }
        cur = cur.Next
    }
    // 返回复制链表的头节点
    return nodeMap[head]
}
