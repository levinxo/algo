
/**
 * https://leetcode-cn.com/problems/string-to-integer-atoi/
 * 8.字符串转整数(atoi)
 * 使用了自动机方法（Automaton）
 *
 * 或者也可使用正则表达式
 */

import java.util.Arrays;
import java.util.List;
import java.util.HashMap;

class Solution {


    private long n = 0;

    private int sign = 1;

    private HashMap<String, List<String>> table;


    public Solution() {
        // 结合之前的状态和现在的字符进行判断，当前状态应该是什么
        table = new HashMap<>();
        table.put("start", Arrays.asList("start", "signed", "in_number", "end"));
        table.put("signed", Arrays.asList("end", "end", "in_number", "end"));
        table.put("in_number", Arrays.asList("end", "end", "in_number", "end"));
        table.put("end", Arrays.asList("end", "end", "end", "end"));
    }


    public int getStatusIndex(char c) {
        if (c == ' ') return 0;
        if (c == '+' || c == '-') return 1;
        if (c >= '0' && c <= '9') return 2;
        return 3;
    }


    public int myAtoi(String str) {
        char[] cnum = str.toCharArray();
        String status = "start";
        for (char c : cnum) {
            status = table.get(status).get(getStatusIndex(c));
            if ("in_number".equals(status)) {
                n = n * 10 + (c - '0');
                n = sign > 0? Math.min(Integer.MAX_VALUE, n) : Math.min(-(long)Integer.MIN_VALUE, n);
            } else if ("signed".equals(status) && c == '-') {
                sign = -1;
            }
        }
        return (int) (sign * n);
    }


    public static void main(String[] args) {
        String n = args[0];
        Solution s = new Solution();
        System.out.println(s.myAtoi(n));
    }


}
