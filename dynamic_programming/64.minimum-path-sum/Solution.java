/**
 * https://leetcode.com/problems/minimum-path-sum/
 * 64. Minimum Path Sum
 * 新建一个dp数组，和原先矩阵保持一致，dp[i][j]记录从坐标i,j到右下角的最小路径值
 * 初始化右下角的dp值为对应的矩阵值，然后去填整个矩阵，每个元素考虑移动到右边或下面
 * 可得转移方程：dp[i,j] = grid[i,j] + max(dp[i+1,j], dp[i,j+1])
 * 同时要注意边界情况，最右列和最下行
 */

import java.lang.Math;

public class Solution {

    public static void main(String... args) {
        int[][] grid = new int[][] { { 1, 3, 1 }, { 1, 5, 1 }, { 4, 2, 1 } };
        System.out.println(minPathSum(grid));
    }

    public static int minPathSum(int[][] grid) {
        if (grid == null || grid.length == 0 || grid[0].length == 0) {
            return 0;
        }
        int row = grid.length;
        int column = grid[0].length;
        int[][] dp = new int[row][column];

        for (int i = row - 1; i >= 0; i--) {
            for (int j = column - 1; j >= 0; j--) {
                if (i == row - 1 && j != column - 1) {
                    // 最下行，只能往右走
                    dp[i][j] = grid[i][j] + dp[i][j + 1];
                } else if (j == column - 1 && i != row - 1) {
                    // 最右列，只能向下走
                    dp[i][j] = grid[i][j] + dp[i + 1][j];
                } else if (i != row - 1 && j != column - 1) {
                    dp[i][j] = grid[i][j] + Math.min(dp[i + 1][j], dp[i][j + 1]);
                } else {
                    dp[i][j] = grid[i][j];
                }
            }
        }

        return dp[0][0];
    }

}