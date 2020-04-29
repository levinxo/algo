import java.lang.Math;

/**
 * https://leetcode.com/problems/longest-valid-parentheses/
 * 32. Longest Valid Parentheses
 * 使用left和right两个变量分别代表"("和")"
 * 从左边开始遍历，如果")"个数大于"("个数，说明是非法的，直接让left/right=0跳过
 * 从右边开始遍历，如果"("个数大于")"个数，说明是非法的，直接让left/right=0跳过
 */

class LongestValidParentheses {

    public static void main(String... args) {
        if (args.length == 0) {
            return;
        }
        System.out.println(args[0]);
        System.out.println("可配对的括号数量：");
        System.out.println(longestValidParentheses(args[0]));
    }

    public static int longestValidParentheses(String s) {
        int left = 0, right = 0, max = 0;
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == '(') {
                left++;
            } else {
                right++;
            }
            if (left == right) {
                // 别忘记乘以2
                max = Math.max(max, right * 2);
            } else if (right > left) {
                left = right = 0;
            }
			// 除了==和>，还有小于的情况（例如"((())"这种），所以下面需要从右遍历一次
        }
        
        left = right = 0;

        for (int i = s.length() - 1; i >= 0; i--) {
            if (s.charAt(i) == ')') {
                right++;
            } else {
                left++;
            }
            if (right == left) {
                // 别忘记乘以2
                max = Math.max(max, left * 2);
            } else if (left > right) {
                left = right = 0;
            }
        }
        return max;
    }

}
