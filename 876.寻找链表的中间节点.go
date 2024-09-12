// https://leetcode.cn/problems/middle-of-the-linked-list/description/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    // 给你单链表的头结点 head ，请你找出并返回链表的中间结点。
    // 如果有两个中间结点，则返回第二个中间结点
    // 快慢指针，快指针双倍步长，慢指针正常步长，快指针的Next或者Next的Next为nil时检索到最后一个节点
    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    if fast.Next == nil {
        return slow
    }
    if fast.Next.Next == nil {
        return slow.Next
    }
    return nil
}
