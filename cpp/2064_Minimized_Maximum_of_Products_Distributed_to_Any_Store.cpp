#include<bits/stdc++.h>
using namespace std;

int minimizedMaximum(int n, vector<int>& q) {
	int res = INT_MAX;
	int qmax = 0;
	for (int i = 0; i < q.size(); ++i)
		qmax = max(qmax, q[i]);

	for (int i = 1; i <= qmax; ++i) {
		int store = 0;
		for (int &qe : q) {
			store += (int)ceil((double)qe / i);
		}

		if (store <= n) {
			res = min(res, i);
		}
	}

	return res;
}

bool canAllProductsDistributed(vector<int>& q, int i , int n) {
	int store = 0;

	for (int &qe : q) {
		cout << "qe: " << qe << ", i: " << i << endl;
		store += (int)ceil((double)qe / i);
	}

	return (store <= n);
}

// O(N∗Log(Max(Q)))
int solveOptimised(int n, vector<int>& q) {
	int r = *max_element(q.begin(), q.end());

	int l = 1;

	int res = INT_MAX;

	while (l <= r) {
		int m = l + (r - l) / 2;

		if (canAllProductsDistributed(q, m, n)) {
			res = min(res, m);
			r = m - 1;
		} else {
			l = m + 1;
		}
	}

	return res;
}

int main() {
	// int t;
	// cin >> t;

	// while (t--) {
	int n, m;
	cin >> n >> m;

	vector<int> q(m);
	for (int i = 0; i < m; ++i)
		cin >> q[i];

	cout << minimizedMaximum(n, q) << endl;
	cout << solveOptimised(n, q) << endl;
	// }
	return 0;
}