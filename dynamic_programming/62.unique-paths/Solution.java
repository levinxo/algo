/**
 * https://leetcode.com/problems/unique-paths/description/
 * 62. Unique Paths
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

    public static int uniquePaths(int m, int n) {
        if (m <= 0 || n <= 0) {
            return 0;
        }
        int[][] dp = new int[n][m];

        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                if (i == 0 && j == 0) {
                    dp[i][j] = 1;
                } else if (i == 0 && j != 0) {
                    // 第一行，只能从左侧来
                    dp[i][j] = dp[i][j - 1];
                } else if (j == 0 && i != 0) {
                    // 第一列，只能从上面来
                    dp[i][j] = dp[i - 1][j];
                } else {
                    dp[i][j] = dp[i][j - 1] + dp[i - 1][j];
                }
                //System.out.println("i: " + i + " j: " + j);
                //System.out.println(dp[i][j]);
            }
        }

        return dp[n - 1][m - 1];
    }

}