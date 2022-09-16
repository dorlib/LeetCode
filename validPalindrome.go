func validPalindrome(s string) bool {
    left, right := 0, len(s) - 1
    for left < right {
        if s[left] != s[right] {
            return validSub(s, left + 1, right) || validSub(s, left, right - 1)
        }
        left++
        right--
    }
    return true
}

func validSub(s string, left int, right int) bool {
    for left < right {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}
