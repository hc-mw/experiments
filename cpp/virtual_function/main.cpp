#include <bits/stdc++.h>
#include "cpp_demo.h"
using namespace std;

void memset_demo() {
	int f[26][1001];

	memset(f, 0, sizeof(f));
	long long sum = 0;
	for (int i = 0; i < 26; ++i) 
		for (int j = 0; j < 1001; ++j) sum += f[i][j];
	cout << "[INFO] sum: " << sum << endl;
}

int main(int argc, char const *argv[])
{
  virtual_function_demo();
	memset_demo();
  return 0;
}


