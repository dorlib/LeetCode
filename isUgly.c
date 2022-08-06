bool isUgly(int n){

    if (n < 1) {
        return false;
    }
    
    while (true) {
        if (n == 1) return true;
        else if (!(n%2)) {
            n = n / 2;
        } else if (!(n%3)) {
            n = n / 3;
        } else if (!(n%5)) {
            n = n / 5;
        } else {
            return false;
        }
    }
}
