#include<bits/stdc++.h>
using namespace std;

class Solution {
public:
    int solve(vector<int>& nums, int target, int currSum, int i, int n, int totalSum, vector<vector<int> >& dp) {
        if (i == n)
            return currSum == target ? 1 : 0;

        if (dp[i][currSum + totalSum] != INT_MIN)
            return dp[i][currSum + totalSum];
        
        int plus = solve(nums, target, currSum + nums[i], i + 1, n, totalSum, dp);
        int minus = solve(nums, target, currSum - nums[i], i + 1, n, totalSum, dp);

        return dp[i][currSum + totalSum] = plus + minus;
    }

    int findTargetSumWays(vector<int>& nums, int target) {
        int n = nums.size();
        int sum = accumulate(nums.begin(), nums.end(), 0);
        vector<vector<int> > dp(n+1, vector<int>(2*sum + 1, 0));

        dp[0][0] = 1;

        for (int i = 0; i < n; ++i) {
          for (int j = 0; j < sum + 1; ++j) {
           dp[i + 1][j - nums[i]] += 1;
           dp[i + 1][j + nums[i]] += 1; 
          }
        }

        return dp[n][target];

        return solve(nums, target, 0, 0, n, sum, dp);
    }
};

int main() {
  Solution obj;

  vector<int> nums{1,1,1,1,1};
  int target = 3;

  int count = obj.findTargetSumWays(nums, target);
  cout << count << endl;
}
