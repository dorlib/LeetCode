func containsDuplicate(nums []int) bool {
    m := make(map[int]int)
    
    for i := 0; i < len(nums); i++ {
        if m[nums[i]] == 0 {
            m[nums[i]] = i + 1
        } else {
            return true
        }
    }
    return false
}
