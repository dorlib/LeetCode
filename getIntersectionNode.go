/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "fmt"

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    
    // lets make a list of pointers to every linked list
    var lstA []*ListNode
    var lstB []*ListNode
    head := headA
    head1 := headB
    
    // insert the pointers to every node in the initialized listA
    for head != nil {
        lstA = append(lstA, head)
        head = head.Next
    }
    
    // insert the pointers to every node in the initialized listB
    for head1 != nil {
        lstB = append(lstB, head1)
        head1 = head1.Next
    }
    
    
    lenA := len(lstA)
    lenB := len(lstB)
    
    // check the conditions nil
    // 1. the last one pointer in every lists are not the same
    // 2. one of the lists is empty
    if lstA[lenA-1] != lstB[lenB-1] || lstA[0] == nil || lstB[0] == nil {
        return nil
    }
    
    
    condition := true
    indexA := lenA - 1
    indexB := lenB - 1 
    
    // lets check from the finish to start where the pointers stops being the same,
    // when we find two pointers from finish that are not the same we can return the previous one.
    for condition && indexA >= 0 && indexB >= 0 {
        if lstA[indexA] != lstB[indexB] {
            condition = false
            return lstA[indexA+1]
        }
        indexA--
        indexB--
    }
    return lstA[indexA+1]
}
