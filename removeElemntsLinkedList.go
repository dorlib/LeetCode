/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    var h *ListNode
    var h2 *ListNode
    
    if head == nil {
        return head
    }
        
    for head != nil && h == nil {
        if head.Val != val {
            h = &ListNode{head.Val, nil}
        }
        head = head.Next
    }
    
    h2 = h
    
    for head != nil {
        if head.Val != val {
            h2.Next = &ListNode{head.Val, nil}
            h2 = h2.Next
        }
        head = head.Next
    }
    return h
}
