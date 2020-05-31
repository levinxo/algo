/**
 * 338. Counting Bits
 * https://leetcode.com/problems/counting-bits/
 */

import java.util.Arrays;


class Solution {


    public static void main(String[] args) {
        int x = 100;
        System.out.println(Arrays.toString(countBits(x)));
    }


    /**
     * Time Complexity O(n)
     * Space Complexity O(n)
     */
    public static int[] countBits(int num) {
        int[] ret = new int[num+1];

        for (int i = 1; i <= num; i++) {
            ret[i] = ret[i&(i-1)] + 1;
        }
        return ret;
    }


}
