/**
 * 407. Trapping Rain Water II
 * https://leetcode.com/problems/trapping-rain-water-ii/
 */
import java.util.PriorityQueue;

public class Solution {

    // 可盛水总数
    private int water = 0;

    // 已经访问过的点
    private boolean[][] visited;

    // 已访问的次数
    private int visitedCnt = 0;

    // 实现一个Comparable接口，用于存放Wall，Wall放入优先队列中
    private static class Wall implements Comparable<Wall> {
        int row;
        int col;
        int value;
        Wall(int row, int col, int value){
            this.row = row;
            this.col = col;
            this.value = value;
        }
        @Override
        public int compareTo(Wall wall) {
            return value - wall.value;
        }
    }


    public int trapRainWater(int[][] heightMap) {
        if (heightMap == null || heightMap.length < 3 || heightMap[0].length < 3) {
            return 0;
        }
        int rows = heightMap.length;
        int cols = heightMap[0].length;
        PriorityQueue<Wall> pq = new PriorityQueue<Wall>();
        visited = new boolean[rows][cols];

        // 下面两个循环将矩阵周围的Wall收集起来
        for (int i = 0; i < cols; i++) {
            Wall wall1 = new Wall(0, i, heightMap[0][i]);
            Wall wall2 = new Wall(rows - 1, i, heightMap[rows - 1][i]);
            pq.offer(wall1);
            pq.offer(wall2);
            visited[0][i] = true;
            visited[rows - 1][i] = true;
            visitedCnt += 2;
        }

        for (int i = 1; i < rows - 1; i++) {
            Wall wall1 = new Wall(i, 0, heightMap[i][0]);
            Wall wall2 = new Wall(i, cols - 1, heightMap[i][cols - 1]);
            pq.offer(wall1);
            pq.offer(wall2);
            visited[i][0] = true;
            visited[i][cols - 1] = true;
            visitedCnt += 2;
        }

        while (pq.size() > 0 && visitedCnt < rows * cols) {
            Wall wall = pq.poll();
            fill(wall.row - 1, wall.col, wall.value, pq, heightMap);
            fill(wall.row, wall.col - 1, wall.value, pq, heightMap);
            fill(wall.row + 1, wall.col, wall.value, pq, heightMap);
            fill(wall.row, wall.col + 1, wall.value, pq, heightMap);
        }
        return water;
    }


    public void fill(int row, int col, int min, PriorityQueue<Wall> pq, int[][] heightMap) {
        if (row < 0 || col < 0) {
            return;
        } else if (row >= heightMap.length || col >= heightMap[0].length) {
            return;
        } else if (visited[row][col]) {
            return;
        } else if (heightMap[row][col] >= min) {
            Wall wall1 = new Wall(row, col, heightMap[row][col]);
            pq.offer(wall1);
            visited[row][col] = true;
            visitedCnt++;
        } else {
            water += min - heightMap[row][col];
            visited[row][col] = true;
            visitedCnt++;
            fill(row - 1, col, min, pq, heightMap);
            fill(row, col - 1, min, pq, heightMap);
            fill(row + 1, col, min, pq, heightMap);
            fill(row, col + 1, min, pq, heightMap);
        }
    }


    public static void main(String[] args) {
        int[][] matrix = new int[][]{
            {1,4,3,1,3,2},
            {3,2,1,3,2,4},
            {2,3,3,2,3,1}
        };

        Solution solution = new Solution();
        System.out.println(solution.trapRainWater(matrix));
    }


}
