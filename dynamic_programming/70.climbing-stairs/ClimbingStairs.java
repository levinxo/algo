/**
 * https://leetcode.com/problems/climbing-stairs/description/
 * 70. Climbing Stairs
 * 思路考虑使用动态规划
 * 定义一个数组dp存储上台阶的方法数，dp[i]表示上台阶i的方法数
 * 最后一步：最后一个台阶i可能有两种方式到达：走1个台阶或者走2个台阶
 * 因此dp[i] = dp[i-1] + dp[i-2]
 * 考虑到dp[i]只与dp[i-1]和dp[i-2]有关，因此可以简化一下由两个变量来存储dp[i-1]和dp[i-2]
 * 使得空间复杂度由O(n)变为O(1)，且动态规划的写法变为了普通的迭代写法
 */

public class ClimbingStairs {

    public static void main(String args[]) {
        if (args.length == 0) {
            return;
        }
        System.out.println("上台阶" + args[0] + "的方法数：");
        System.out.println(climbStairs(Integer.parseInt(args[0])));
    }

    // 最好是用long类型，int很容易溢出
    public static int climbStairs(int n) {
        if (n <= 2) {
            return n;
        }
        int pre1 = 2, pre2 = 1;
        for (int i = 2; i < n; i++) {
            int curr = pre1 + pre2;
            pre2 = pre1;
            pre1 = curr;
        }
        return pre1;
    }

}