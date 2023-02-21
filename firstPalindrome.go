func firstPalindrome(words []string) string {
    for i := 0; i < len(words); i++ {
        if isPalindrom(words[i]) {
            return words[i]
        }
    }

    return ""
}

func isPalindrom(word string) bool {
    first := 0
    last := len(word) - 1

    for first < last {
        if word[first] != word[last] {
            return false
        }

        first++
        last--
    }

    return true
}
