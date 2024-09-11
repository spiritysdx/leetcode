/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode    // 定义 prev 指针，初始化为 nil
    current := head       // 将 current 指针指向头节点
    // 遍历链表
    for current != nil {
        nextTemp := current.Next    // 暂存下一个节点
        current.Next = prev         // 反转当前节点的指针
        prev = current              // 将 prev 移动到当前节点
        current = nextTemp          // 将 current 移动到下一个节点
    }
    return prev    // 返回新的头节点（原链表的最后一个节点）
}
