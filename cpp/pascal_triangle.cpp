#include <bits/stdc++.h>
#define ll long long
using namespace std;

ll nCr(ll n, ll r)
{
	ll res = 1;

	// for (; r != 0; r--, n--) {
	// 	res *= n;
	// 	res /= r;
	// }

	for (int i = 0; i < r; i++)
	{
		res += (n - i);
		res -= (i + 1);
	}

	return res;
}

vector<vector<ll>> pascalTriangle(ll n)
{
	vector<vector<ll>> v;

	for (int i = 1; i <= n; i++)
	{
		vector<ll> vv;

		for (int j = 1; j <= i; j++)
		{
			ll temp = nCr(i - 1, j - 1);
			// printf("[Info] %dC%d: %lld\n", i - 1, j - 1, temp);
			vv.push_back(temp);
		}

		v.push_back(vv);
	}

	return v;
}

void printVector(vector<vector<ll>> &v)
{
	for (auto vv : v)
	{
		for (auto x : vv)
		{
			cout << x << ", ";
		}
		cout << endl;
	}
}

int main(int argc, char const *argv[])
{

	int n;

	cin >> n;

	auto triangle = pascalTriangle(n);

	printVector(triangle);

	return 0;
}