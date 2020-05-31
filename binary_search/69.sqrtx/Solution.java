/**
 * 69. Sqrt(x)
 * https://leetcode.com/problems/sqrtx/
 */

class Solution {


    /**
     * Time Complexity O(log n)
     * Space Complexity O(n)
     */
    public static int mySqrt(int x) {
        if (x <= 1) {
            return x;
        }

        int l = 1, r = x;
        while (l <= r) {
            int m = l + (r - l) / 2;
            int sqrt = x / m;
            if (m == sqrt) {
                return m;
            } else if (m > sqrt) {
                r = m - 1;
            } else {
                l = m + 1;
            }
        }
        return r;
    }


}
