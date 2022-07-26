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


// second solution 

func climbStairs(n int) int {
    if n < 3 { return n }
    
    dp := make([]int, n + 1)
    dp[0] = 0
    dp[1] = 1
    dp[2] = 2
    
    for i := 3; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    
    return dp[n]
}
