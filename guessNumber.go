/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

import "math"

func guessNumber(n int) int {
    min := 0
    max := n
    
    for min <= max {
        mid := int(math.Floor(float64(min + max) / 2))
        res := guess(mid)
        if res == 0 {
            return mid
        }
        if res == -1 {
            max = mid - 1
        } else {
            min = mid + 1
        }
    }
    return int(math.Floor(float64(n / 2)))
}
