func isPrefixString(s string, words []string) bool {
    return isPrefixStringReq(s, words)
}

func isPrefixStringReq(s string, words []string) bool {
    if s == "" {
        return true
    }
    
    if len(words) == 0 && len(s) > 0 {
        return false
    }
    
    word := words[0]
    wordLength := len(word)
    
    if wordLength > len(s) {
        return false
    }
    
    if s[:wordLength] == word {
        return isPrefixStringReq(s[wordLength:], words[1:])
    } else {
        return false
    }
    
}
