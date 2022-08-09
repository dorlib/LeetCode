/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
    m := make(map[*TreeNode]int) 
    m =  generateMap(root, m, 0)
    return calculate(m)
}

func generateMap(root *TreeNode, m map[*TreeNode]int, depth int) map[*TreeNode]int {
    if root.Left == nil && root.Right == nil {
        m[root] = depth
        return m 
    }
    
    if root.Left != nil {
        generateMap(root.Left, m, depth + 1)
    }
    
    if root.Right != nil {
        generateMap(root.Right, m, depth + 1)
    }
    return m 
}

func calculate(m map[*TreeNode]int) int {
    depth := 0
    sum := 0 
    
    for _, value := range(m) {
        if value > depth {
            depth = value
        }
    }
    
    for key, value := range(m) {
        if value == depth {
            sum += key.Val
        }
    }
    return sum
}
