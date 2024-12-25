// https://leetcode.cn/problems/reverse-linked-list-ii/description/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 本题需要画图便于理解
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    // 由于right和left从1开始计数，所以都要-1代表真实位置
    dummyNode := &ListNode{Next: head}
    p0 := dummyNode
    for i := 0; i < left-1; i++ { // p0需要到达需要反转的上一个节点的上一个节点
        p0 = p0.Next
    }
    // 正常反转right-left+1次
    var pre *ListNode
    current := p0.Next 
    for i := left-1; i<right; i++ { // 本来right要-1的，因为上面的需求，不需要-1了
        next := current.Next
        current.Next = pre
        pre = current
        current = next
    }
    // 逆转完毕后，p0现在还是在需要反转的上一个节点的上一个节点，current已经到了需要反转的最后一个节点的后一个节点上了
    // pre指向反转链表的末尾节点(反转链表本身的最后一个节点)， cur指向反转链表的下一个节点(不在反转链表中)，中断节点的第一个节点
    p0.Next.Next = current // p0.Next最后到未拼接前最后中断的节点的上一个节点
    p0.Next = pre 
    return dummyNode.Next
}

// // 定义链表节点结构
// // ListNode 表示链表节点，包含一个值和指向下一个节点的指针。
// type ListNode struct {
//     Val  int
//     Next *ListNode
// }

// // 反转链表中从第 left 到第 right 个节点的子链表（1-based 索引）。
// func reverseBetween(head *ListNode, left int, right int) *ListNode {
//     // 创建一个虚拟头节点，指向原链表的头节点，方便处理头节点的特殊情况。
//     dummyNode := &ListNode{Next: head}
    
//     // p0 用来定位到需要反转区间的前一个节点的上一个节点。
//     p0 := dummyNode
    
//     // 循环找到 p0，使其停在第 left-1 个节点位置。
//     for i := 0; i < left-1; i++ {
//         p0 = p0.Next
//     }
    
//     // pre 用于反转过程中记录当前节点的前驱节点（初始为 nil）。
//     // current 是当前需要处理的节点，初始化为反转区间的第一个节点。
//     var pre *ListNode
//     current := p0.Next
    
//     // 循环执行反转操作，次数为 right-left+1（反转区间长度）。
//     for i := left - 1; i < right; i++ {
//         next := current.Next // 临时保存下一个节点的引用。
//         current.Next = pre   // 让当前节点指向前驱节点，完成反转。
//         pre = current        // 更新 pre，指向当前节点（为下次迭代准备）。
//         current = next       // 更新 current，处理下一个节点。
//     }
    
//     // 反转完成后，p0.Next 仍指向反转区间前的第一个节点（此时是反转后的尾部节点）。
//     // p0.Next.Next（反转后尾部节点的 Next）需要连接到反转区间后第一个未反转的节点。
//     p0.Next.Next = current
    
//     // p0.Next 改为指向反转后的第一个节点（即 pre）。
//     p0.Next = pre
    
//     // 返回新的链表头节点（dummyNode.Next）。
//     return dummyNode.Next
// }
