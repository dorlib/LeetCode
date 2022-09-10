func reversePrefix(word string, ch byte) string {
    var charIndex int
    for i := 0; i < len(word); i++ {
        if word[i] == ch {
            charIndex = i
            break
        }
    }
    
    if charIndex == 0 {
        return word
    } else {
        return reverseString(word[:charIndex + 1]) + word[charIndex + 1:]
    }
}

func reverseString(str string) string{
   byte_str := []rune(str)
   for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
      byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
   }
   return string(byte_str)
}
