import java.lang.Math;

public class Solution {

    public static void main(String[] args) {
        int[][] arr = new int[][]{
            {-2, -3, 3},
            {-5, -10, 1},
            {10, 30, -5}
        };
        System.out.println(calculateMinimumHP(arr));
    }


    public static int calculateMinimumHP(int[][] arr) {
        /**
         * 1.最后一步和子问题
         * 骑士只能向下或向右，假定有一个二维数组存储了当前骑士进入房间之前的生命值
         * 如果房间是正数，那么进入的生命值为1即可
         * 如果房间是负数，那么进入的生命值为1-arr[i][j]即可
         * 因此这两个生命值取最大的即可
         * 2.状态转移公式
         * 3.初始化和边界条件
         * 注意最下面只能往右走，最右边只能往下走
         * 4.计算方向
         * 从尾部开始
         */
        if (arr == null || arr.length == 0 || arr[0].length == 0) {
            return 0;
        }
        int row = arr.length;
        int col = arr[0].length;

        int[][] dp = new int[row][col];

        for (int i = row - 1; i >= 0; i--) {
            for (int j = col - 1; j >= 0; j--) {
                if (i == row - 1 && j == col - 1) {
                    // 初始化
                    dp[i][j] = Math.max(1, 1 - arr[i][j]);
                } else if (i == row - 1) {
                    dp[i][j] = Math.max(1, dp[i][j + 1] - arr[i][j]);
                } else if (j == col - 1) {
                    dp[i][j] = Math.max(1, dp[i + 1][j] - arr[i][j]);
                } else {
                    dp[i][j] = Math.max(1, Math.min(dp[i + 1][j] - arr[i][j], dp[i][j + 1] - arr[i][j]));
                }
            }
        }
        return dp[0][0];
    }


}
