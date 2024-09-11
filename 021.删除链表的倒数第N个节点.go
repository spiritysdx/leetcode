/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // 创建一个哑节点，它的下一个节点是头节点
    dummy := &ListNode{0, head}
    // 定义两个指针，first指向头节点，second指向哑节点
    first, second := head, dummy
    // 让first指针先走n步
    for i := 0; i < n; i++ {
        first = first.Next
    }
    // 让first和second指针同时移动，直到first指针到达链表的末尾
    for ; first != nil; {
        first = first.Next
        second = second.Next
    }
    // 此时second指针指向的是待删除节点的前一个节点
    // 将待删除节点从链表中移除
    second.Next = second.Next.Next
    // 返回哑节点的下一个节点，即新的头节点
    return dummy.Next
}
