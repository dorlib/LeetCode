/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if (head == nil) {return head}
    var current *ListNode = head
    
    for current.Next != nil {
        if(current.Next.Val == current.Val){
            current.Next = current.Next.Next;
        }else{current = current.Next}
    }
    
    return head
}
