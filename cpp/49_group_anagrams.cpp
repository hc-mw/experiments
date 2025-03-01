#include<bits/stdc++.h>
using namespace std;

string getKey(string str) {
    int count[26];
    memset(count, 0, sizeof(int) * 26);

    for (char ch : str)
        count[ch - 'a']++;

    string key = "";
    for (int i = 0; i < 26; i++)
        key.append(char(i + 'a') + to_string(count[i]));

    return key;
}

int main() {
    auto strs = vector<string> {"bdddddddddd", "bbbbbbbbbbc"};
    int x = 26, y = strs.size();
    unordered_map<string, vector<string>> map;


    for (auto &str : strs)
        map[getKey(str)].push_back(str);


    vector<vector<string>> res;

    for (auto &pair : map) {
        res.push_back(pair.second);
    }

    for (auto &i : res) {
        cout << "[" << " ";
        for (auto &j : i) cout << j << " ";
        cout << "]\n";
    }

    return 0;
}