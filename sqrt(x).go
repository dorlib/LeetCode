func strStr(haystack string, needle string) int {
    
    needleLength := len(needle)
    
    if needleLength > len(haystack) {
        return -1
    }
    
    for i := 0 ; i < len(haystack); i++ {
        if haystack[i] == needle[0] {
            n := i + needleLength
            if len(haystack[i:]) >= needleLength {
                if haystack[i:n] == needle {
                    return i
                }
            }
        }
    }
    return -1
}
