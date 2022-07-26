public class Solution 
{
    public int RomanToInt(string s) 
    {
              
        IDictionary<char, int> dict = new Dictionary<char, int>();
        dict.Add('I', 1);
        dict.Add('V', 5);
        dict.Add('X', 10);
        dict.Add('L', 50);
        dict.Add('C', 100);
        dict.Add('D', 500);
        dict.Add('M', 1000);
        
        int result = dict[Convert.ToChar(s[s.Length - 1])];  
        
        for (int i = s.Length - 2; i >= 0; i--)
        {
            char current = s[i];
            char prev = s[i + 1];
            int currentVal = dict[Convert.ToChar(current)];
            int prevVal = dict[Convert.ToChar(prev)];
            
            if (currentVal < prevVal)
            {
                result -= (currentVal);
            } else if (currentVal >= prevVal)
            {
                result += (currentVal);
            }            
        }
        
        return result;
    }
}
