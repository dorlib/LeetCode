func fib(n int) int {
    
    if n == 0 {
        return 0
    }
    
    if n == 1 {
        return 1
    }
    
    fibSlice := make([]int, n)
    
    fibSlice[0] = 1
    fibSlice[1] = 1
    
    for i := 2; i < n; i++ {
        fibSlice[i] = fibSlice[i-1] + fibSlice[i - 2]
    }
    
    return fibSlice[n-1]
}
