func plusOne(digits []int) []int {
    n := len(digits)
    
    for i := n - 1; i >= 0 ; i-- {
        if digits[i] < 9 {
            digits[i]++
            return digits
        }
        // the index points to 9
        digits[i] = 0
    }
    
    // the array is from the format: [9,9,9,9,....]
    res := make([]int, n + 1)
    res[0] = 1
    return res
}
