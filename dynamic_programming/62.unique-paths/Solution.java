/**
 * https://leetcode.com/problems/unique-paths/description/
 * 62. Unique Paths
 * https://www.bilibili.com/video/BV1xb411e7ww
 */
public class Solution {

    public static void main(String... args) {
        if (args.length != 2) {
            return;
        }
        int m = Integer.parseInt(args[0]);
        int n = Integer.parseInt(args[1]);

        System.out.println(uniquePaths(m, n));
    }

	/**
	 * Time complexity O(m*n)
	 * Space complexity O(m*n)
	 */
    public static int uniquePaths(int m, int n) {
		/**
		 * 1.确定状态：最后一步和子问题：
		 * 假定右下角坐标为(m-1,n-1)，有x种方式从左上角走到(m-2,n-1)，
		 * 有y种方式从左上角走到(m-1,n-2)，那么机器人有x+y种方式走到(m-1,n-1)
		 * 问题转化为子问题，有多少种方式从左上角走到(m-2,n-1)和(m-1,n-2)
		 * 2.转移方程
		 * dp[i][j] = dp[i][j-1] + dp[i-1][j]
		 * 3.初始条件和边界情况：
		 * dp[0][0] = 1
		 * 注意第一行、第一列，分别是没有上面和左侧的格子的，只能有一个方向过来，因此dp[i][j] = 1
		 * 4.计算顺序
		 * 先循环i，即从上到下，再循环j，即从左到右，这样可以满足转移方程可以取到值
		 */
        if (m <= 0 || n <= 0) {
            return 0;
        }
        int[][] dp = new int[n][m];

        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                if (i == 0 || j == 0) {
                    dp[i][j] = 1;
                } else {
                    dp[i][j] = dp[i][j - 1] + dp[i - 1][j];
                }
            }
        }

        return dp[n - 1][m - 1];
    }

}
