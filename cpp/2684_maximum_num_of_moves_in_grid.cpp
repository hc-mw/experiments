#include<bits/stdc++.h>
using namespace std;

int maxMoves(vector<vector<int>>&, vector<vector<int>>&, int, int);

int main() {


	vector<vector<int>> grid;
	int m, n;
	cin >> m >> n;

	for (int i = 0; i < m; i++) {
		vector<int> t;
		for (int j = 0; j < n; j++) {
			int x;
			cin >> x;
			t.push_back(x);
		}
		grid.push_back(t);
	}

	int maxMove = -1;

	vector<vector<int>> dp(m, vector<int>(n, -1));

	for (int i = 0; i < m; ++i)
		maxMove = max(maxMove, maxMoves(grid, dp, i, 0));

	cout << maxMove << endl;

	return 0;
}


int maxMoves(vector<vector<int>>& grid, vector<vector<int>>& dp, int m, int n) {
	if (m < 0 || m >= grid.size() || n < 0 || n >= grid[0].size())
		return 0;

	if (dp[m][n] != -1)
		return dp[m][n];

	int r = grid.size() , c = grid[0].size();

	int max_move = 0;

	vector<int> moves{ -1, 0, 1};
	for (int& move : moves) {
		int new_m = m + move;
		int new_n = n + 1;

		if (new_m >= 0 && new_m < r && new_m >= 0 && new_n < c && grid[new_m][new_n] > grid[m][n]) {
			max_move = max(max_move, 1 + maxMoves(grid, dp, new_m, new_n));
		}
	}

	return dp[m][n] = max_move;
}