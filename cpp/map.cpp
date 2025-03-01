#include<bits/stdc++.h>
using namespace std;

int main(int argc, char const *argv[]) {
	string word;

	cin >> word;

	string res = "";
	int i = 0, n = word.length();

	while (i < n) {
		int count = 0;
		char ch = word[i];

		while (ch == word[i] && count < 9 && i < n) {
			i++;
			count++;
		}

		res += to_string(count) + ch;
	}


	cout << res << endl;

	return 0;
}