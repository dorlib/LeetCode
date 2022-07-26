func climbStairs(n int) int {
    mem := make(map[int]int)
    res := climbStairsReq(n, mem)
    return res
}
 
func climbStairsReq(n int, mem map[int]int ) int {
    if _, ok := mem[n]; ok {
		return (mem[n])
	}

	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	res1, res2 := climbStairsReq(n-1, mem), climbStairsReq(n-2, mem)

	mem[n] = res1 + res2

	return res1 + res2
}
