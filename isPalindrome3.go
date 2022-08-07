/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "math"

func isPalindrome(head *ListNode) bool {
    var values []int
    
    for head != nil {
        values = append(values, head.Val)
        head = head.Next
    }
        
    for i := 0; i < int(math.Floor(float64(len(values) / 2))); i++ {
        if values[i] != values[len(values) - i - 1] {
            return false
        }
    }
    return true
}
