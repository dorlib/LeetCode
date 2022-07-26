import (
    "fmt"
)

func lengthOfLastWord(s string) int {
    count := 0 
    count2 := 0
    n := len(s) - 1 
    
    if n == 0 && string(s[0]) != " " {
        return 1
    }
    
    // lets find where the last word finishes
    for count = 0; string(s[n - count]) == " "; count++ {}
    
    
    // lets get the length of the last word
    for count2 = 0; string(s[n - count]) != " " ; n = n  {
        count2 ++
        n--
        if n - count < 0 {
            break
        }
    }
    
    return count2
}
