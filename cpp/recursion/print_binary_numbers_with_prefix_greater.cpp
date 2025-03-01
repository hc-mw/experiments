#include<bits/stdc++.h>
using namespace std;

void solve(int nZeros, int nOnes, int n, string op) {
    if (n == 0) {
        cout << op << endl;
        return;
    }

    solve(nZeros, nOnes + 1, n - 1, op + "1");

    if (nOnes > nZeros)
        solve(nZeros + 1, nOnes, n - 1, op + "0");
}

int main() {
    int n;
    cin >> n;
    solve(0, 0, n, "");
    return 0;
}