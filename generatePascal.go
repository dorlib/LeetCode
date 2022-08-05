func generate(numRows int) [][]int {
    result := [][]int{{1}}
    
    if numRows == 1 {
        return result
    }
    
    if numRows >= 2 {
        result = append(result, []int{1,1})
    }
    
    res := []int{1}
    
    for i:= 3; i < numRows + 1; i++ {
        above := result[i -2]
        for j := 0; j < i - 2; j++ {
            res = append(res, above[j] + above[j+1])
        }
        res = append(res,1)
        result = append(result,res)
        res = []int{1}

    }
    
    return result
}
