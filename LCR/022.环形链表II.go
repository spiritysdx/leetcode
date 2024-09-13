/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    seen := map[*ListNode]struct{}{} // 创建一个map用于存储已访问过的节点，指针做键，指使用struct{}做0空间大小的值
    for head != nil { // 遍历链表
        if _, ok := seen[head]; ok { // 如果当前节点已经在map中，说明有环，返回该节点
            return head
        }
        seen[head] = struct{}{} // 将当前节点添加到map中
        head = head.Next // 移动到下一个节点
    }
    return nil // 如果遍历完整个链表没有发现环，返回nil
}
