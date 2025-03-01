#include<bits/stdc++.h>
using namespace std;

stack<int> st;
int min;

void push(int val) {
	if (st.empty()) {
		st.push(val);
		min = val;
		return;
	}

	if (val < min) {
		long long temp = 2  * val - min;
		st.push(temp);
		min = val;
		return;
	}

	st.push(val);
}

void pop() {
	if (st.empty())
		return;

	int top = st.top();
	st.pop();

	if (top < min) min = 2 * min - top;

	st.pop();
}

int top() {
	if (st.empty()) return -1;

	if (st.top() >=  min) return st.top();

	return min;
}

int getMin() {
	if (st.empty()) return -1;
	return min;
}

int main(int argc, char const *argv[]) {
	string s;
	getline(cin, s);

	vector<string> ops = splitString(s);

	for (string str : ops) {
		cout << str << endl;
	}

	return 0;
}

vector<string> splitString(string str) {
	vector<string> res;

	for (int i = 0; i < str.size(); i++) {
		if (str[i] == "\"") {
			string t = "";
			while (str[i] != "\"") {
				i++;
				t.push_back(str[i]);
				continue;
			}

			res.push_back(t);
		}
	}

	return res;
}