class Solution {
    public int minimumSum(int num) {
        int arr[] = new int[4];
        for (int i = 0; i < 4; i++) {
            arr[i] = num % 10;
            num = num / 10;
        }
        
        Arrays.sort(arr);
        
        int num1 = 0;
        int num2 = 0;
        
        for (int i = 0; i < 4; i++) {
            if ((i % 2) == 0)
                num1 = (num1 * 10) + arr[i];
            else
                num2 = (num2 * 10) + arr[i];
        }
        
        return num1 + num2;
    }
}
