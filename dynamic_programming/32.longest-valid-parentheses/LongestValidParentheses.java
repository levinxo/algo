
/**
 * https://leetcode.com/problems/longest-valid-parentheses/
 * 32. Longest Valid Parentheses
 * 这是一道"最值型"题目，考虑使用动态规划
 *
 * 1. "最后一步"分析：对于有效的括号来说，最后一个元素一定是')'
 * 定义一个dp数组，其中第i个元素表示到当前元素为止最长有效子字符串的长度
 * 当dp[i] == '('时，无法构成有效括号对，因此dp[i] = 0
 * 当dp[i] == ')'时，分两种情况：
 *   * dp[i-1] == '('时，即dp[i-1]和dp[i]组成一对有效括号，再考虑dp[i-2]这个位置的括号对，此时dp[i] = dp[i-2] + 2
 *   * dp[i-1] == ')'时，类似"..(())"这样的括号对，那么需要判断位置i-dp[i-1]-1是否是'('
 *     且还得同时加上在位置i-dp[i-1]-2的括号对，类似"()(())"这样的情况
 * 2. 得到子问题和转移方程：
 * 类似'..()'情况：dp[i] = dp[i-2] + 2    同时要保证i-2>=0
 * 类似'..(())'情况：dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2    同时要保证i-dp[i-1]-2>=0
 *
 * Time complexity O(n)
 * Space complexity O(n)
 */

import java.lang.Math;

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
        int max = 0;
        if (s.length() == 0) {
            return max;
        }

        int dp[] = new int[s.length()];
        for (int i = 1; i < s.length(); i++) {
            if (s.charAt(i) == ')') {
                if (s.charAt(i - 1) == '(') {
                    dp[i] = (i - 2 >= 0 ? dp[i - 2] : 0) + 2;
                } else if (i - dp[i - 1] - 1 >= 0 && s.charAt(i - dp[i - 1] - 1) == '(') {
                    dp[i] = dp[i - 1] + 2 + (i - dp[i - 1] - 2 >= 0 ? dp[i - dp[i - 1] - 2] : 0);
                }
                max = Math.max(max, dp[i]);
            }
        }

        return max;
    }

}
