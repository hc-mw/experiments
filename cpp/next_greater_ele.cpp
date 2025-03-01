#include<bits/stdc++.h>
using namespace std;

vector<int> NGE(vector<int> &v) {
	vector<int> res(v.size());
	stack<int> st;
	for (int i = 0; i < v.size(); ++i) {
		while (!st.empty() && v[st.top()] < v[i]) {
			res[st.top()] = i;
			st.pop();
		}
		st.push(i);
	}
	while (!st.empty()) {
		res[st.top()] = -1;
		st.pop();
	}
	return res;
}

// Input:
// 6
// 4 5 2 25 7 8
// Output:
// 5 25 25 -1 8 -1

int main(int argc, char const *argv[]) {
	ios::sync_with_stdio(0);
	cin.tie(0);

	int n;
	cin >> n;
	vector<int> v(n);

	for (int i = 0; i < n; i++)  // O(t*c1)
		cin >> v[i];

	auto res = NGE(v);


	for (auto i : res)
		cout << (i == -1 ? -1 : v[i]) <<  " ";

	cout << "\n";
	return 0;
}