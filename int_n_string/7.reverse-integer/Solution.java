/**
 * 7. 整数反转
 * https://leetcode-cn.com/problems/reverse-integer/
 */
class Solution {


    public static void main(String[] args) {
        int x = 15566;
        System.out.println(reverse(x));
    }


    public static int reverse(int x) {
        int rev = 0;
        while (x != 0) {
            int pop = x % 10;
            if (rev > Integer.MAX_VALUE / 10
                    || (rev == Integer.MAX_VALUE / 10 && pop > Integer.MAX_VALUE % 10)) {
                return 0;
            } else if (rev < Integer.MIN_VALUE / 10
                    || (rev == Integer.MIN_VALUE / 10 && pop < Integer.MIN_VALUE % 10)) {
                return 0;
            }
            rev = rev * 10 + pop;
            x = x / 10;
        }
        return rev;
    }


}
