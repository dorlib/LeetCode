func isIsomorphic(s string, t string) bool {
    sMap := make(map[string]string)
    for i := 0; i < len(s); i++ {
        if sMap[string(s[i])] != "" {
            if sMap[string(s[i])] != string(t[i]) {
                return false
            }
        } else {
            if !checkForValue(string(t[i]), sMap) {
                sMap[string(s[i])] = string(t[i])
            } else {
                return false
            }
        }
    }
    return true
}

func checkForValue(val string, m map[string]string)bool{

  //traverse through the map
  for _,value:= range m{

    //check if present value is equals to val
    if(value == val){

      //if same return true
      return true
    }
  }

  //if value not found return false
  return false
}
