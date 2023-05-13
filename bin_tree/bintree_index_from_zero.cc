#include <bits/stdc++.h>
using namespace std;

using ll = long long;
using pll = pair<ll, ll>;
using pii = pair<int, int>;

class BinTree : vector<ll> {
public:
    ll c0 = 0;
    explicit BinTree(ll k = 0) // 默认初始化一个能保存k个元素的空树状数组
    {
        assign(k + 1, 0); // 有效下标从0开始
    }
    int lowbit(ll k)
    {
        return k & -k;
        // 也可写作x&(x^(x–1))
    }
    ll sum(ll k) // [0,..,k] 的和，k为下标，不是个数
    {
        ll res = _sum(k) + c0;
        return res;
    }
    ll query(ll b, ll e)
    {
        ll res = sum(e);
        if (b > 0) {
            res -= sum(b - 1);
        }
        return res;
    }

    ll _sum(ll k) // 求第1个元素到第n个元素的和
    {
        ll res = (k > 0 ? _sum(k - lowbit(k)) + (*this)[k] : 0);
        return res;
    }
    ll last() // 返回最后一个元素下标
    {
        return size() - 1;
    }
    void add(ll k, ll w) // k 为下标
    {
        if (k == 0) {
            c0 += w;
            return;
        }
        _add(k, w);
    }
    void _add(ll k, ll w) // 为节点k加上w, k为下标
    {
        if (k > last())
            return;
        (*this)[k] += w;
        _add(k + lowbit(k), w);
    }
};
