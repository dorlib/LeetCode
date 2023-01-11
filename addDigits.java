class Solution {
    public int addDigits(int num) {
        int res = add(num);

        if (String.valueOf(res).length() == 1) {
            return res;
        }

        return addDigits(res);
    }

    private int add(int num) {
        int res = 0;
        String numAsString = String.valueOf(num);

        for (int i = 0; i < numAsString.length(); i++) {
            res += Integer.parseInt(String.valueOf(numAsString.charAt(i)));
        }

        return res;
    }
}
