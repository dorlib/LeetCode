/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    m := make(map[*ListNode]float64)
    
    for head != nil {
            if m[head] != 0.0 {
                return true
            } else {
                m[head] = float64(head.Val)
            }
            head = head.Next
        }
    return false
}
