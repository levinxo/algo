
/**
 * https://leetcode.com/problems/minimum-path-sum/
 * 64. Minimum Path Sum
 * 新建一个dp数组，和原先矩阵保持一致，dp[i][j]记录从坐标i,j到右下角的最小路径值
 * 初始化右下角的dp值为对应的矩阵值，然后去填整个矩阵，每个元素考虑移动到右边或下面
 * 可得转移方程：dp[i,j] = grid[i,j] + min(dp[i+1,j], dp[i,j+1])
 * 同时要注意边界情况，最右列和最下行
 */

import java.lang.Math;

public class Solution {

    public static void main(String... args) {
        int[][] grid = new int[][] { { 1, 3, 1 }, { 1, 5, 1 }, { 4, 2, 1 } };
        System.out.println("第1种方法：");
        System.out.println(minPathSum1(grid));
        System.out.println("第2种方法：");
        System.out.println(minPathSum2(grid));
    }

    public static int minPathSum1(int[][] grid) {
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

    public static int minPathSum2(int[][] grid) {
        if (grid == null || grid.length == 0 || grid[0].length == 0) {
            return 0;
        }
        int row = grid.length;
        int column = grid[0].length;
        int[][] dp = new int[row][column];

        for (int i = 0; i < row; i++) {
            for (int j = 0; j < column; j++) {
                if (i == 0 && j == 0) {
                    dp[i][j] = grid[i][j];
                } else if (i == 0 && j != 0) {
                    // 第一行，只能从左侧位置过来
                    dp[i][j] = grid[i][j] + dp[i][j - 1];
                } else if (j == 0 && i != 0) {
                    // 第一列，只能从上方位置过来
                    dp[i][j] = grid[i][j] + dp[i - 1][j];
                } else {
                    dp[i][j] = grid[i][j] + Math.min(dp[i - 1][j], dp[i][j - 1]);
                }
            }
        }

        return dp[row - 1][column - 1];
    }

}