func prefixCount(words []string, pref string) int {
    var counter int
    lenPref := len(pref)
    for i := 0; i < len(words); i++ {
        if checkPref(words[i], pref, lenPref) {
            counter++
        }
    }
    return counter
}

func checkPref(word string, pref string, prefLength int) bool {
    if prefLength > len(word) {
        return false
    }
    
    for i := 0; i < prefLength; i++ {
        if word[i] != pref[i] {
            return false
        }
    }
    return true
}
