// 定义链表节点的结构体
type ListNode struct {
    Val  int
    Next *ListNode
}

// 主函数：重新排列链表
func reorderList(head *ListNode) {
    // 如果链表为空，则直接返回
    if head == nil {
        return
    }

    // 找到链表的中点
    mid := middleNode(head)
    
    // l1 是链表的前半部分，l2 是链表的后半部分
    l1 := head
    l2 := mid.Next
    
    // 将链表从中点断开
    mid.Next = nil
    
    // 将后半部分的链表反转
    l2 = reverse(l2)
    
    // 合并前半部分和反转后的后半部分
    mergeList(l1, l2)
}

// 反转链表的辅助函数
// 该函数接收一个链表的头节点，并返回反转后的链表头节点
func reverse(head *ListNode) *ListNode {
    var prev *ListNode  // 定义一个指针用于存储前一个节点
    cur := head         // 当前节点从头节点开始遍历
    // 循环遍历链表，直到当前节点为 nil
    for cur != nil {
        next := cur.Next  // 记录当前节点的下一个节点
        cur.Next = prev   // 反转当前节点的指针指向前一个节点
        prev = cur        // 将前一个节点更新为当前节点
        cur = next        // 当前节点前进到下一个节点
    }
    return prev  // 返回反转后的新头节点
}

// 寻找链表的中间节点的辅助函数
// 该函数通过快慢指针法找到链表的中点
func middleNode(head *ListNode) *ListNode {
    slow, fast := head, head  // 定义快指针和慢指针，初始都指向头节点
    // 当快指针没有到达链表尾部时，慢指针每次走一步，快指针走两步
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next      // 慢指针走一步
        fast = fast.Next.Next // 快指针走两步
    }
    return slow  // 返回慢指针，指向链表的中间节点
}

// 合并两个链表的辅助函数
// 该函数交替合并链表 l1 和 l2，形成重新排序后的链表
func mergeList(l1, l2 *ListNode) {
    var l1Tmp, l2Tmp *ListNode  // 临时存储 l1 和 l2 的下一个节点
    // 当 l1 和 l2 都不为空时，交替合并它们
    for l1 != nil && l2 != nil {
        l1Tmp = l1.Next  // 记录 l1 的下一个节点
        l2Tmp = l2.Next  // 记录 l2 的下一个节点

        l1.Next = l2     // l1 的下一个节点指向 l2
        l1 = l1Tmp       // l1 前进到下一个节点

        l2.Next = l1     // l2 的下一个节点指向 l1
        l2 = l2Tmp       // l2 前进到下一个节点
    }
}
