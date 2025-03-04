// https://leetcode.cn/problems/reverse-linked-list/description/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    cur := head
    // 迭代反转链表 - 双指针
    // 一个变量存当前指针，一个变量存实际未反转前的下一个指针
    for cur != nil {
        next := cur.Next       // 保存当前节点的下一个节点，防止断链
        cur.Next = prev        // 将当前节点的 Next 指向前一个节点，反转指针方向

        prev = cur             // 将 prev 移动到当前节点
        cur = next             // cur 移动到下一个节点，继续迭代
    }
    // 返回反转后的链表头
    return prev
}


// package main
// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )
// // 定义单链表节点
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }
// // 反转单链表
// func reverseList(head *ListNode) *ListNode {
// 	var pre *ListNode = nil
// 	cur := head
// 	for cur != nil {
// 		next := cur.Next  // 保存下一个节点
// 		cur.Next = pre    // 反转当前节点指向
// 		pre = cur         // 移动 pre 到当前节点
// 		cur = next        // 移动 cur 到下一个节点
// 	}
// 	return pre // 返回新的头结点
// }
// // 从数组构造链表
// func createList(arr []int) *ListNode {
// 	if len(arr) == 0 {
// 		return nil
// 	}
// 	dummy := &ListNode{} // 虚拟头节点
// 	cur := dummy
// 	for _, v := range arr {
// 		cur.Next = &ListNode{Val: v}
// 		cur = cur.Next
// 	}
// 	return dummy.Next
// }
// // 打印链表
// func printList(head *ListNode) {
// 	cur := head
// 	for cur != nil {
// 		fmt.Printf("%d ", cur.Val)
// 		cur = cur.Next
// 	}
// 	fmt.Println()
// }

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	line, _ := reader.ReadString('\n')
// 	line = strings.TrimSpace(line)

// 	// 处理输入
// 	nums := []int{}
// 	if line != "" {
// 		for _, s := range strings.Split(line, " ") {
// 			n, _ := strconv.Atoi(s)
// 			nums = append(nums, n)
// 		}
// 	}

// 	// 构造链表
// 	head := createList(nums)

// 	// 反转链表
// 	newHead := reverseList(head)

// 	// 输出反转后的链表
// 	printList(newHead)
// }
