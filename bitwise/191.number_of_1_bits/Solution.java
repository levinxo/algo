/**
 * 191. Number of 1 Bits
 * https://leetcode.com/problems/number-of-1-bits/
 */
public class Solution {


    public static void main(String[] args) {
        int n = 100;
        System.out.println(hammingWeight(n));
    }


    public static int hammingWeight(int n) {
        int cnt = 0;
        while (n != 0) {
            cnt++;
            n = n & (n - 1);
        }
        return cnt;
    }


}
