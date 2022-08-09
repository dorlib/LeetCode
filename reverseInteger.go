func reverse(x int) int {
    result := 0
    res := x
    
    for res != 0 {
        num := res % 10
        result = result * 10 + num
        res = res / 10
    }
    
    if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	} else {
		return result
	}
}
