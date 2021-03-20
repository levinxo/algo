#include <vector>
#include <queue>
#include <utility>
#include <fmt/ranges.h>
#include <fmt/color.h>


using std::vector;
using std::queue;
using std::pair;
using std::make_pair;


/**
 * https://leetcode-cn.com/problems/walls-and-gates/
 * use multi source bfs.
 */

namespace bfs {

class Solution {

private:

    const int Door = 0;
    const int Empty = ((unsigned) 1 << 31) - 1;

    const int Positions[4][2] =  {
        {1, 0},
        {-1, 0},
        {0, 1},
        {0, -1},
    };

public:

    void wallsAndGates(vector<vector<int> >& rooms) {
        if (rooms.empty() || rooms[0].empty()) {
            return;
        }

        int row_len = rooms.size();
        int col_len = rooms[0].size();
        pair<int, int> node;
        int row, col, new_row, new_col;

        // find all doors, and push to queue
        queue<pair<int, int> > q;
        for (int i = 0; i < row_len; i++) {
            for (int j = 0; j < col_len; j++) {
                if (rooms[i][j] == Door) {
                    q.emplace(make_pair(i, j));
                }
            }
        }

        while (!q.empty()) {
            node = q.front();
            q.pop();

            row = node.first;
            col = node.second;

            for (auto pos: Positions) {
                new_row = row + pos[0];
                new_col = col + pos[1];

                if (new_row < 0 || new_col < 0
                    || new_row >= row_len || new_col >= col_len
                    || rooms[new_row][new_col] != Empty) {
                    continue;
                }

                rooms[new_row][new_col] = rooms[row][col] + 1;
                q.emplace(make_pair(new_row, new_col));
            }
        }
    }

};


void test_walls_and_gates() {
    fmt::print(fg(fmt::color::yellow), "Walls and Gates\n");
    vector<vector<int> > rooms = {
        {2147483647,-1,0,2147483647},
        {2147483647,2147483647,2147483647,-1},
        {2147483647,-1,2147483647,-1},
        {0,-1,2147483647,2147483647},
    };

    fmt::print("Origin rooms: {}\n", rooms);
    Solution solution;
    solution.wallsAndGates(rooms);
    fmt::print("Finished rooms: {}\n", rooms);
}


}       // namespace bfs


