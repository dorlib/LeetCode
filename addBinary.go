func addBinary(a string, b string) string {
    longer, shorter := a, b
    if len(b) > len(a) {
        longer, shorter = b, a
    }
    result := make([]byte, len(longer))
    c := byte(0)
   
    for i, j := len(shorter)-1, len(longer)-1 ; j >=0 ; i, j = i-1, j-1 {
        s := byte(0)
        if i >= 0 {
            s = shorter[i] - '0'
        }
        l := longer[j] - '0'
        sum := s + l + c
        switch sum {
        case 2, 3:
            result[j] = sum % 2 + '0'
            c = 1
        default:
            result[j] = sum + '0'
            c = 0
        }
    }
    if c == 1 {
        result = append([]byte{'1'}, result...)
    }
    return string(result)
}
