func findKthPositive(arr []int, k int) int {
    low := 0 
    high := len(arr)
    
    if k > arr[high - 1] - high {
        return k + high
    }
        
    for low <= high {
        mid := int((float64(low + high) / 2))
        
        if (arr[mid] - mid - 1) < k {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return low + k 
}
