func containsNearbyDuplicate(nums []int, k int) bool {
    m := make(map[int]int)
    
    for i := 0; i < len(nums); i++ {
        if m[nums[i]] != 0.0 {
            if i - m[nums[i]] + 1 <= k {
                return true
            }
        }
        m[nums[i]] = i + 1
    }
    return false
}
