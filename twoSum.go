func twoSum(nums []int, target int) []int {
    cache :=make(map[int]int)
    var result [] int
    
    for i,item := range nums{
        j , ok := cache[item]
        if ok {
            result = append(result,j)
            result = append(result,i)
        }
        cache[target - item] = i
    }
    
    return result
}
