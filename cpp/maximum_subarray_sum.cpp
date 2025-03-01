#include<bits/stdc++.h>
using namespace std;

int maximumSubarraySum(int arr[], int n) {
	int max_sum_so_far = INT_MIN;
	int sum = 0;

	for (int i = 0; i < n; i++) {
		sum += arr[i];
		if (sum > max_sum_so_far) max_sum_so_far = sum;
		if (sum < 0) sum = 0;
	}

	return max_sum_so_far;
}

vector<int> maximumSubarraySumIdx(int arr[], int n) {
	int max_sum_so_far = INT_MIN;
	int sum = 0;
	int s = -1, e = -1;

	for (int i = 0; i < n; i++) {
		sum += arr[i];
		if (sum > max_sum_so_far) {
			max_sum_so_far = sum;
			e = i;
		}
		if (sum < 0) {
			sum = 0;
			s = i + 1;
		}
	}

	return {s, e};
}

int main() {
	int arr[] = {};
	int n = sizeof(arr) / sizeof(arr[0]);
	int sum = maximumSubarraySum(arr, n);
	vector<int> res = maximumSubarraySumIdx(arr, n);
	for (int x : res) {
		cout << x << ", ";
	}
	cout << endl;
	return 0;
}