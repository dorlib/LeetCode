class Solution {
    public List<String> fizzBuzz(int n) {
        List<String> result = new ArrayList<String>();
        for (int i = 1; i < n+1; i++) {
            if (i % 3 == 0) {
                if (i % 5 == 0) {
                    result.add("FizzBuzz");
                    continue;
                }
                result.add("Fizz");
                continue;
            }
            if (i % 5 == 0) {
                result.add("Buzz");
                continue;
            }

            result.add(String.valueOf(i));
        }

        return result;
    }
}
