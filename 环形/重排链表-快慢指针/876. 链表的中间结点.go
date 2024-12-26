// https://leetcode.cn/problems/middle-of-the-linked-list/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}

// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */
// func middleNode(head *ListNode) *ListNode {
//     slow, fast := head, head
//     for fast.Next != nil && fast.Next.Next != nil {
//         slow = slow.Next
//         fast = fast.Next.Next
//     }
//     if fast.Next == nil {
//         return slow
//     }
//     if fast.Next.Next == nil {
//         return slow.Next
//     }
//     return nil
// }
