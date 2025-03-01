#include<bits/stdc++.h>
using namespace std;


long long countFairPairs(vector<int>& nums, int lower, int upper) {
	sort(nums.begin(), nums.end());

	long long res = 0;

	for (int i = 0; i < nums.size(); ++i) {
		int x = lower_bound(nums.begin() + i + 1, nums.end(), lower - nums[i]) - nums.begin();
		int y = upper_bound(nums.begin() + i + 1, nums.end(), upper - nums[i]) - nums.begin();
		res += (y - x);
	}

	return res;
}

int main() {
	int t;
	cin >> t;

	while (t--) {
		int n, lower, upper;
		cin >> n;
		cin >> lower;
		cin >> upper;

		vector<int> nums(n);

		for (int i = 0; i < n; ++i)
			cin >> nums[i];

		cout << countFairPairs(nums, lower, upper) << endl;
	}
}
