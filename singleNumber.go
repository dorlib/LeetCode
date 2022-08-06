func singleNumber(nums []int) int {
    m := make(map[int]int)
    
    for i := 0; i < len(nums); i++ {
        if m[nums[i]] == 0.0 {
            m[nums[i]] = 1
        } else {
            m[nums[i]]++
        }
    }
    
    for j := range m{
        if m[j] == 1 {
            return j
        }
    }
    
    // according to the description we will not get here 
    return 0
}
