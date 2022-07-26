import "math"

func searchInsert(nums []int, target int) int {
    n := len(nums) - 1
    max := n 
    min := 0 
    var mid int
    
    for ok := true; ok; ok = min <= max {
        mid = int(math.Floor(float64((max+min) / 2)))
        if nums[mid] == target {
            return mid
        } else if nums[mid] > target {
            max = mid - 1
        } else {
            min = mid + 1
        }
    }
    
    // checks....
    if nums[mid] < target {
        return mid + 1
    }
    
    if nums[mid] >= target {
        return mid
    }
    
    return mid
}
