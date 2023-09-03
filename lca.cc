#include <bits/stdc++.h>
using namespace std;

using ll = long long;
using pll = pair<ll, ll>;
using pii = pair<int, int>;

class Solution
{
public:
	vector<vector<int>> ps;
	vector<vector<array<int, 3>>> g;
	vector<array<int, 27>> cnts;
	vector<int> hs;

	void dfs(int u, int p, int h)
	{
		ps[u][0] = p;
		hs[u] = h;
               
		for (auto &e : g[u])
		{
			int v = e[1], w = e[2];
			if (v == p)
			{
				continue;
			}
            
            cnts[v] = cnts[u];
			cnts[v][w]++;
			dfs(v, u, h + 1);
		}
	}

	int getp(int u, int h)
	{
		for (int i = 0; i < 21; i++)
		{
			if (h & (1 << i))
			{
				int p = ps[u][i];
				h -= (1 << i);
				if (h == 0)
				{
					return p;
				}
				return getp(p, h);
			}
		}
        return 0;
	}

	int lca(int u, int v)
	{
		int hu = hs[u], hv = hs[v];
		if (hu > hv)
		{
			u = getp(u, hu - hv);
			hu = hv;
		}
		else if (hv > hu)
		{
			v = getp(v, hv - hu);
			hv = hu;
		}

		int h = hu;
		if (u == v)
		{
			return u;
		}

		int l = 1, r = h;
		while (l <= r)
		{
			int m = l + (r - l) / 2;
			int pu = getp(u, m), pv = getp(v, m);
			if (pu == pv)
			{
				r = m - 1;
			}
			else
			{
				l = m + 1;
			}
		}

		int p = getp(u, l);
		return p;
	}

	vector<int> minOperationsQueries(int n, vector<vector<int>> &edges, vector<vector<int>> &queries)
	{
		g = vector<vector<array<int, 3>>>(n);
		cnts = vector<array<int, 27>>(n);
		ps = vector<vector<int>>(n, vector<int>(21, -1));
		hs = vector<int>(n, 0);
		for (auto &e : edges)
		{
			int u = e[0], v = e[1], w = e[2];
			g[u].push_back({u, v, w});
			g[v].push_back({v, u, w});
		}
        
        dfs(0,-1,0);

		for (int l = 1; l < 21; l++)
		{
			for (int u = 0; u < n; u++)
			{
				int p = ps[u][l - 1];
				if (p >= 0)
				{
					ps[u][l] = ps[p][l - 1];
				}
			}
		}

		vector<int> res;
		for (auto &q : queries)
		{
			int u = q[0], v = q[1];

			int p = lca(u, v);
            
			array<int, 27> cu = cnts[u], cv = cnts[v], cp = cnts[p];
			array<int, 27> cs;
			int mxc = 0, sumc = 0;
			for (int i = 0; i < 27; i++)
			{
				cs[i] = cu[i] + cv[i] - 2 * cp[i];
				mxc = max(mxc, cs[i]);
				sumc += cs[i];
			}

			res.push_back(sumc - mxc);
		}
		return res;
	}
};

//https://leetcode.cn/contest/weekly-contest-361/problems/minimum-edge-weight-equilibrium-queries-in-a-tree/