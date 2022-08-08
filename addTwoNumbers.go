/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    res := &ListNode{}
    result, num := res, 0
    
    for l1 != nil || l2 != nil || num > 0 {
        sum := num
        
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }
        
        result.Next = &ListNode{Val : sum % 10, Next : nil}
        result = result.Next
        num = sum / 10
    }
    return res.Next
}
