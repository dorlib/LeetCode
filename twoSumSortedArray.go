func twoSum(numbers []int, target int) []int {
    endIndex := len(numbers) - 1
    startIndex := 0
    
    for endIndex > startIndex {
        sum := numbers[startIndex] + numbers[endIndex]
        if sum == target {
            return []int{startIndex + 1, endIndex + 1}
        } else if sum > target {
            endIndex--
        } else {
            startIndex++
        }
    }
    return []int{0,0}
}
