/**
 * 9. Palindrome Number
 * https://leetcode.com/problems/palindrome-number/submissions/
 */

import java.util.Arrays;

public class Solution {


    public static void main(String[] args) {
        int x = 1661;
        System.out.println(isPalindrome1(x));
        System.out.println(isPalindrome2(x));
    }


    public static boolean isPalindrome1(int x) {
        String str = String.valueOf(x);
        int start = 0, end = str.length() - 1;

        while (start < end) {
            if (str.charAt(start) != str.charAt(end)) {
                return false;
            }
            start++;
            end--;
        }
        return true;
    }


    public static boolean isPalindrome2(int x) {
        if (x < 0 || (x != 0 && x % 10 == 0)) {
            return false;
        }

        // 将x拆成两半之后再对比，后半部分要逆序
        int reverse = 0;
        while (x > reverse) {
            reverse = reverse * 10 + x % 10;
            x /= 10;
        }

        // 若数字长度为奇数，则要将翻转的数字去掉一位比较
        return x == reverse || x == reverse / 10;
    }


}
