#include <bits/stdc++.h>
using namespace std;
using ll = long long;
using pll = pair<ll, ll>;
using pii = pair<ll, ll>;

const int MX = 10000007;
int min_div[MX];

void sieve()
{
    for (int i = 0; i < MX; i++) {
        min_div[i] = i;
    }

    for (int i = 2; i * i < MX; i++) {
        if (min_div[i] == i) {
            for (int j = i * i; j < MX; j += i) {
                if (min_div[j] == j) {
                    min_div[j] = i;
                }
            }
        }
    }
}

vector<int> fs(int a)
{
    vector<int> res;
    while (a > 1) {
        int b = min_div[a];
        res.push_back(b);
        while (a % b == 0) {
            a /= b;
        }
    }
    return res;
}
