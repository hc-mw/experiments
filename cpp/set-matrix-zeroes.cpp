#include<bits/stdc++.h>
using namespace std;

void setZeroes(vector<vector<int>>& mat) {
	int m = mat.size();
	int n = mat[0].size();
	int col0 = 1;

	// #1: iterate over array and mark accordinlgy
	for (int i = 0; i < m; i++) {
		for (int j = 0; j < n; j++) {
			if (mat[i][j] == 0) {
				// mark row
				mat[i][0] = 0;
				// mark column
				if (j == 0)
					col0 = 0;
				else
					mat[0][j] = 0;
			}
		}
	}

	// #2: set 0s in array from 1:1 to M-1:N-1
	for (int i = 1; i < m; i++) {
		for (int j = 1; j < n; j++) {
			if (mat[i][j] != 0) {
				if (mat[0][j] == 0 || mat[i][0] == 0)
					mat[i][j] = 0;
			}
		}
	}

	// #3: set 1st row and 1st col accordingly
	// rows
	if (mat[0][0] == 0)
		for (int i = 0; i < n; i++)
			mat[0][i] = 0;
	// columns
	if (col0 == 0)
		for (int i = 0; i < n; i++)
			mat[i][0] = 0;
}

vector<vector<int>> getInputVector() {
	int m;
	int n;

	cin >> m;
	cin >> n;
	vector<vector<int>> v(m, vector<int>(n));

	for (int i = 0; i < m; i++)
		for (int j = 0; j < n; j++)
			cin >> v[i][j];

	return v;
}

void printVectorMatrix(vector<vector<int>>& mat) {
	for (int i = 0; i < mat.size(); i++) {
		for (int j = 0; j < mat[0].size(); j++) {
			cout << mat[i][j] << " ";
		}
		cout << "\n";
	}
}

int main() {
	auto v = getInputVector();

	setZeroes(v);

	printVectorMatrix(v);

	return 0;
}