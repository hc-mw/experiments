#include <bits/stdc++.h>
using namespace std;

// Function to find the maximum subarray sum
int maxSubarraySum(vector<int> &arr) {
    int sum = arr[0];
    int maxSumSoFar = arr[0];

    for (int i = 1; i < arr.size(); i++) {
      // if prevSum is negative, start new sum
      if (sum < 0) {
        sum = arr[i];
      // else extend previous sum
      } else {
        sum += arr[i];
      }
      maxSumSoFar = max(maxSumSoFar, sum);
    }
    return maxSumSoFar;
}

int main() {
    vector<int> arr = {2, 3, -8, 7, -1, 2, 3};
    cout << maxSubarraySum(arr);
    return 0;
}
