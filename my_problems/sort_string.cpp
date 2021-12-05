/**
 * Problem
 *
 * 给定两个长度为 N (N < 1e6) 的字符串S，T (仅由小写字母组成)
 * 重新排列S，使得S为S >= T  所有排列中字母序最小的
 *
 * 例子：
 * S = "abcd", T = "dcaz"
 * Result = "dcba"
 */

#include <bits/stdc++.h>
using namespace std;
using ll = long long;
using pii = pair<int, int>;
using pll = pair<ll, ll>;

int N;
string S, T;
string R;

bool dfs(int i, multiset<char>& st)
{
    if (i >= N) {
        return true;
    }

    char a = T[i];
    auto it = st.lower_bound(a);
    if (it == st.end()) {
        return false;
    }
    char b = *it;
    if (b == a) {
        R.push_back(b);
        st.erase(it);
        bool f = dfs(i + 1, st);
        if (f) {
            return true;
        } else {
            R.pop_back();
            st.insert(b);
            it = st.upper_bound(a);
            if (it == st.end()) {
                return false;
            }
            b = *it;
            R.push_back(b);
            st.erase(it);
            for (char c : st) {
                R.push_back(c);
            }
            return true;
        }

    } else {
        R.push_back(b);
        st.erase(it);
        for (char c : st) {
            R.push_back(c);
        }
        return true;
    }
    return false;
}

string solve()
{
    N = S.size();
    multiset<char> st(S.begin(), S.end());
    bool f = dfs(0, st);
    if (f) {
        return R;
    } else {
        return "-1";
    }
}

int main()
{
    S = "abcd";
    T = "dcaz";
    cin >> S >> T;
    cout << solve() << endl;
}