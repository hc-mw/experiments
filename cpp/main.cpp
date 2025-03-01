#include<bits/stdc++.h>
using namespace std;

int main(int argc, char const *argv[]) {
	cout << argc << endl;

	int n = sizeof(argv) / sizeof(char);

	for (int i = 0; i < n; i++) cout << argv[i] << " ";

	return 0;
}
