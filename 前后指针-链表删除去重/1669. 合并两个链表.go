// https://leetcode.cn/problems/merge-in-between-linked-lists/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    // 找到 a 的前一个节点和 b 的后一个节点
    cur := list1
    var prevA *ListNode
    var afterB *ListNode
    // 找到 a 的前一个节点
    for i := 0; i < a; i++ {
        prevA = cur
        cur = cur.Next
    }
    // 找到 b 的后一个节点
    for i := 0; i <= b-a; i++ {
        cur = cur.Next
    }
    afterB = cur
    // 连接 list2
    prevA.Next = list2
    // 找到 list2 的尾节点
    cur = list2
    for cur.Next != nil {
        cur = cur.Next
    }
    // 将 list2 的尾节点与 afterB 连接
    cur.Next = afterB
    return list1
}
