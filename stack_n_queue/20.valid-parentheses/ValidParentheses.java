import java.util.HashMap;
import java.util.Map;
import java.util.Stack;

// 有效的括号
// leetcode: https://leetcode.com/problems/valid-parentheses/

public class ValidParentheses {

    /**
     * 想象成一个类似俄罗斯方块不停向下丢括号的栈，正括号会一直累积，
     * 反括号降落时若能配对上正括号，那么就可以消除，否则直接返回false
     * 
     * @param s
     * @return boolean
     */
    public static boolean isValid(String s) {
        Map<Character, Character> pMap = new HashMap<>();
        pMap.put(')', '(');
        pMap.put('}', '{');
        pMap.put(']', '[');

        Stack<Character> stack = new Stack<>();

        for (int i = 0; i < s.length(); i++) {
            if (pMap.get(s.charAt(i)) != null && !stack.isEmpty()) {
                // 反括号的情况下，尝试消栈
                if (stack.peek() == pMap.get(s.charAt(i))) {
                    stack.pop();
                } else {
                    return false;
                }
            } else {
                // 正括号的情况下，直接入栈累积
                stack.add(s.charAt(i));
            }
        }

        return stack.isEmpty();
    }

    public static void main(String... args) {
        if (args.length == 0) {
            return;
        }

        System.out.println(args[0]);
        if (isValid(args[0])) {
            System.out.println("valid");
        } else {
            System.out.println("not valid");
        }
    }

}