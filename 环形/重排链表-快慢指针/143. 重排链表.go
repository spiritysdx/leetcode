// https://leetcode.cn/problems/reorder-list/description/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
    // 找到链表的中间节点
    mid := findMiddle(head)
    // 反转链表的后半部分
    head2 := reverse(mid)
    // 交替合并前半部分和反转后的后半部分
    for head2.Next != nil {
        nxt := head.Next         // 保存前半部分链表的下一个节点
        nxt2 := head2.Next       // 保存后半部分链表的下一个节点
        head.Next = head2        // 将后半部分的当前节点接到前半部分后
        head2.Next = nxt         // 将前半部分的下一个节点接到后半部分后
        head = nxt               // 移动前半部分链表的指针
        head2 = nxt2             // 移动后半部分链表的指针
    }
}

// 找到链表的中间节点
func findMiddle(head *ListNode) *ListNode {
    slow := head                // 慢指针
    fast := head                // 快指针
    // 快指针每次移动两步，慢指针每次移动一步
    for fast != nil && fast.Next != nil {
        slow = slow.Next        // 慢指针移动一步
        fast = fast.Next.Next   // 快指针移动两步
    }
    return slow                 // 返回慢指针，指向中间节点
}

// 反转链表
func reverse(head *ListNode) *ListNode {
    cur := head                 // 当前节点
    var pre *ListNode           // 前一个节点，初始为 nil
    // 遍历链表并反转指针方向
    for cur != nil {
        nxt := cur.Next         // 保存当前节点的下一个节点
        cur.Next = pre          // 当前节点的指针指向前一个节点
        pre = cur               // 前一个节点更新为当前节点
        cur = nxt               // 当前节点更新为下一个节点
    }
    return pre                  // 返回反转后的链表头节点
}
