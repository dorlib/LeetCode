/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubPath(head *ListNode, root *TreeNode) bool {
    if head == nil {
        return true
    }
    
    originalHead := head
    return isSubPathReq(head, root, originalHead)
}

func isSubPathReq(head *ListNode, root *TreeNode, realHead *ListNode) bool {
    val := head.Val
    
    if root == nil {
        return false
    }
    
    if root.Val == val && head.Next == nil {
        return true
    }
    
    if root.Val != val && root.Left == nil && root.Right == nil {
        return false
    }
    
    if root.Val != val && (root.Left != nil || root.Right != nil) {
        var left bool
        var right bool
        
        if root.Val == realHead.Val {
            for root.Val == realHead.Val {
                realHead = realHead.Next
            }
            
            if root.Left != nil {
                left = isSubPathReq(realHead, root.Left, realHead)
            }
        
            if root.Right != nil {
                right = isSubPathReq(realHead, root.Right, realHead)
            }
        } else {
            if root.Left != nil {
                left = isSubPathReq(realHead, root.Left, realHead)
            }
        
            if root.Right != nil {
                right = isSubPathReq(realHead, root.Right, realHead)
            }
        }
        return right || left
    }
    
    if root.Val == val && head.Next != nil {
        node := head.Next
        var left bool
        var right bool
        
        if root.Left != nil {
            left = isSubPathReq(node, root.Left, realHead)
        }
        
        if root.Right != nil {
            right = isSubPathReq(node, root.Right, realHead)
        }
        
        return right || left
    }
    
    return true
}



// not mine, but a lot more simple 
func isSubPath(h *ListNode, r *TreeNode) bool {
    var isContiguous func(h *ListNode, r *TreeNode) bool
    isContiguous = func(h *ListNode, r *TreeNode) bool {
        if h == nil { return true }
        if r == nil { return false }
        if h.Val == r.Val {
            return isContiguous(h.Next, r.Left) ||
                    isContiguous(h.Next, r.Right)
        }
        return false
    }
    
    if h == nil { return true }
    if r == nil { return false }
    if h.Val == r.Val && isContiguous(h,r) { return true }
    return isSubPath(h, r.Left) || isSubPath(h, r.Right)
}
