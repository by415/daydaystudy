/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main
import "fmt"
//Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    if (l1 == nil || l2 == nil) {
        return nil
    }
    var sign int = 0
    var sum int = 0
    var head *ListNode = nil
    var tail *ListNode = head
    for l1!= nil || l2!= nil {
        var cur *ListNode = &ListNode{0,nil}
        if (l1 == nil) {
            sum = l2 -> Val + sign
            l2 = (*l2).Next
        } else if (l2 == nil){
            sum = ((*l1).Val) +sign
            l1 = (*l1).Next
        } else {
            sum = ((*l1).Val + (*l2).Val) + sign
            l1 = (*l1).Next
            l2 = (*l2).Next
        }
        (*cur).Val = sum%10
        sign = sum/10
        if (head == tail && head == nil) {
            head = cur
            tail = cur
        } else {
            (*tail).Next = cur
            tail= cur
        }
    }
    if (sum >= 10){
        var cur *ListNode = &ListNode{0,nil}
        (*cur).Val = sum/10
        (*tail).Next = cur
        tail= cur
    }
    return head
}
func main() {
	var p1 ListNode = ListNode{9,nil}
	var p2 ListNode = ListNode{9,nil}
	p1.Next = &p2
	var l1 ListNode = ListNode{1,nil}

	var x1 ListNode = ListNode{2,nil}
	var x2 ListNode = ListNode{4,nil}
	var x3 ListNode = ListNode{3,nil}
	x1.Next = &x2
	x2.Next = &x3
	var y1 ListNode = ListNode{5,nil}
	var y2 ListNode = ListNode{6,nil}
	var y3 ListNode = ListNode{4,nil}
	y1.Next = &y2
	y2.Next = &y3


	var ret *ListNode = addTwoNumbers(&p1,&l1)
	var retxy *ListNode = addTwoNumbers(&x1,&y1)
	for ret != nil {
		fmt.Println(ret.Val)
		ret = (*ret).Next
	}
	fmt.Println("==========================")
	for retxy != nil {
		fmt.Println(retxy.Val)
		retxy = (*retxy).Next
	}
}
