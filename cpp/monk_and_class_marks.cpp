#include<bits/stdc++.h>
using namespace std;

int main(int argc, char const *argv[]) {
	int t;
	cin >> t;
	vector<pair<int, string>> v;

	while (t--) { // O(t*c1)
		// get name and marks
		string name;
		int mark;

		cin >> name >> mark;
		v.push_back({mark, name});
	}

	sort(v.begin(), v.end()); // O(t*log(t))

	for (auto &p : v) { // O(t*c2)
		cout << p.second << " " << p.first << "\n";
	}

	return 0;
}