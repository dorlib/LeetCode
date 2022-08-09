/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


// this function returns the middle of a linked list
// with fast and slow pointers, in each iteration the fast pointer goes twise the length of the slow one
// so when the fast finishes, we know that the slow one points to the middle node of the linked list
func getMiddle(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    
    slow := head
    fast := head.Next
    
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}


func mergeSort(head1 *ListNode, head2 *ListNode) *ListNode {
    result := &ListNode{}
    tail := result
    
    for head1 != nil && head2 != nil {
        if head1.Val < head2.Val {
            tail.Next = head1
            head1 = head1.Next
        } else {
            tail.Next = head2
            head2 = head2.Next
        }
        
        tail = tail.Next
    }
    
    if head1 != nil {
        tail.Next = head1
    } else {
        tail.Next = head2
    }
    return result.Next
}

func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    middle := getMiddle(head)
    skip := middle.Next
    middle.Next = nil 
    return mergeSort(sortList(head), sortList(skip))   
}
