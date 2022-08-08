/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

import (
    "math"
)

func firstBadVersion(n int) int {
    min := 0
    max := n 
    
    for min <= max {
        mid := int(math.Floor(float64((min + max) / 2)))
        if isBadVersion(mid) {
            if mid == 1 {
                return 1
            } else if !isBadVersion(mid - 1) {
                return mid 
            } else {
                max = mid - 1
            }
        } else {
            min = mid + 1
        }
    }
    // we will not get here anyway...
    return n
}
