/****
Problem: HackEarth 2020 June Circuit 20 P1
***/
#include <bits/stdc++.h>
using namespace std;

typedef long long ll;
typedef pair<int,int> pii;
typedef pair<ll,ll> pll;

const ll MX = 200005;

ll N, M;
//tree
vector<ll> g[MX];

//parents sparse table
ll ps[MX][30];

//depth of every point
ll depths[MX];

//children numbers
ll cns[MX];

//chain top node
ll top[MX];

//weight son of every node
ll wsons[MX];

//index of the node in its chain
ll idxInChain[MX];

//chains
vector<ll> chains[MX];


set<ll> tops;


struct Node {
	ll v, v2;
	ll lazy, lazy2;
	ll lv, rv;
	ll s20;
	ll ct;
	Node *left ,*right;
	Node(ll l, ll r, ll ct){
		v=0; lazy=0;
		this->ct = ct;
		v2 = 0; lazy2 = 0;
		s20 = 0;
		lv=l; rv=r;
		left=NULL; right=NULL;
		if(l == r){
			s20 = depths[chains[ct][l]];
		}

		if(lv>=rv) return;
		ll m=lv + (rv-lv)/2;
		left=new Node(l, m, ct);
		right=new Node(m+1, r, ct);
		s20 = left->s20 + right->s20;
	}

	void output(){
		cout<<lv<<" "<<rv<<" "<<v<<endl;
		if(left!=NULL) left->output();
		if(right!=NULL) right->output();
	}

	void pushdown(){
		if(lazy!=0){
			ll mv = lv + (rv-lv)/2;
			if(left!=NULL) {
				left->v += lazy * (mv - lv + 1);
				left->lazy += lazy;
			}
			if(right!=NULL) {
				right->v += lazy * (rv - mv);
				right->lazy += lazy;
			}
			lazy=0;
		}
		if(lazy2!=0){
			ll mv = lv + (rv-lv)/2;
			if(left!=NULL) {
				left->v2 += lazy2 * left->s20;
				left->lazy2 += lazy2;
			}
			if(right!=NULL) {
				right->v2 += lazy2 * right->s20;
				right->lazy2 += lazy2;
			}
			lazy2=0;
		}

	}

	void add(ll l, ll r, ll vv, ll vv2){
		//cout<<l<<" "<<r<<" "<<lv<<" "<<rv<<endl;
		if(lv==l && rv==r){
			v += vv * ( rv- lv + 1);
			v2 += vv2 * s20;
			lazy += vv;
			lazy2 += vv2;
			return;
		}
		pushdown();
		ll m=lv + (rv - lv)/2;
		if(r<=m){
			left->add(l, r, vv, vv2);
		}else if (l>m){
			right->add(l, r, vv, vv2);
		}else{
			left->add(l, m, vv, vv2);
			right->add(m+1, r, vv, vv2);
		}
		v=left->v + right->v;
		v2 = left->v2 + right->v2;
	}

	ll query(ll l, ll r){
		if(l==lv && r==rv){
			return v + v2;
		}
		pushdown();
		ll m=lv + (rv - lv)/2;
		if(r<=m){
			return left->query(l, r);
		}else if(l>m){
			return right->query(l, r);
		}else{
			return left->query(l,m) + right->query(m+1, r);
		}
		return 0;
	}
};

Node* trees[MX];

void init(){
	memset(ps, 0, sizeof(ps));
	memset(depths, 0, sizeof(depths));
	memset(cns, 0, sizeof(cns));
	memset(top, 0, sizeof(top));
	memset(wsons, 0, sizeof(wsons));
	memset(trees, 0, sizeof(trees));
}

//calculate some initial parameters.
ll dfs(ll u, ll p, ll h){
	depths[u] = h;
	ps[u][0] = p;
	ll cn = 0;
	wsons[u] = -1;
	ll ccn = 0;

	for(ll v : g[u]){
		if(v == p) continue;
		ll ccn1 = dfs(v, u, h+1);
		if(ccn < ccn1){
			ccn = ccn1;
			wsons[u] = v;
		}
		cn += ccn1;
	}

	cns[u] = cn + 1;
	return cns[u];
}

//decomposite the tree
void dfs2(ll u, ll p, ll t){
	tops.insert(t);
	top[u] = t;
	ll idx = chains[t].size();
	chains[t].push_back(u);
	idxInChain[u] = idx;

	if(wsons[u] >= 0){
		dfs2(wsons[u], u, t);
	}

	for(ll v : g[u]){
		if(v == p) continue;
		if(v == wsons[u]) continue;
		dfs2(v, u, v);
	}
}

//cal lca
ll lca(ll u, ll v){
	ll tu = top[u], tv = top[v];
	if(tu == tv){
		return depths[u] > depths[v] ? v : u;

	}

	if(depths[tu] >= depths[tv]){
		return lca(ps[tu][0], v); 
	}else{
		return lca(u, ps[tv][0]); 
	}
}

//get parent with d0 distance.
ll queryParent(ll u, ll d0){
	assert(d0 >= 0);
	if(d0 == 0) return u;
	if(d0 == 1) return ps[u][0];

	ll d = d0;
	ll s = 0;
	while(d > 0){
		d = d >> 1;
		s++;
	}
	s--;
	return queryParent(ps[u][s], d0 ^ (1LL<<s));
}

//add values to node between u and v
void add(ll u, ll v, ll val, ll val2){
	assert(v > 1 && u > 1);
	ll iu = idxInChain[u], iv = idxInChain[v];
	ll tu = top[u], tv = top[v];
	if(tu == tv){
		trees[tu]->add(min(iv, iu), max(iv, iu), val, val2);

	}else{
		if(depths[tu] < depths[tv]){
			add(v, tv, val, val2);
			ll ptv = ps[tv][0];
			add(u, ptv, val, val2);
		}else{
			add(u, tu, val, val2);
			ll ptu = ps[tu][0];
			add(ptu, v, val, val2);
		}
	}
}

//query sum of the values in path between u and v
ll querySum(ll u, ll v){
	ll tu = top[u], tv = top[v];
	ll iu = idxInChain[u], iv = idxInChain[v];
	
	if(tu == tv){
		return trees[tu]->query(min(iu, iv), max(iu, iv));
	}

	if(depths[tu] < depths[tv]){
		ll sumv = querySum(v, tv);
		ll ptv = ps[tv][0];
		return sumv + querySum(u, ptv);
	}else{
		ll sumu = querySum(u, tu);
		ll ptu = ps[tu][0];
		return sumu + querySum(ptu, v);
	}
}


int main(){
	init();
	cin>>N;
	ll u, v, w;
	for(ll i=0; i<N-1; i++){
		cin>>u>>v;
		g[u].push_back(v);
		g[v].push_back(u);
	}


	dfs(1, 0, 0);

	cerr<<"dfs done"<<endl;

	dfs2(1, 0, 1);

	cerr<<"dfs2 done"<<endl;



	for(ll p=1; p<20; p++){
		for(ll i=1; i<=N; i++){
			ps[i][p] = ps[ps[i][p-1]][p-1];
		}
	}

	for(ll t : tops){
		trees[t] = new Node(0, ((ll)chains[t].size()-1), t);
	}

	cerr<<"create trees"<<endl;



//	for(int u=1; u<=N; u++){
//		cout<<u<<" "<<top[u]<<" "<<wsons[u]<<endl;
//	}

	cin>>M;

	for(ll i=0; i<M; i++){
		cin>>u>>v>>w;
		if(u == v || w <= 0) continue;
		ll a = lca(u, v);

		if(a == u){
			ll n = depths[v] - depths[a];
			ll vp = queryParent(v, n-1);
			if(n * (n+1)/2 <= w){
				add(v, vp, 1 - depths[vp] , 1);

			}else{
				ll x = (sqrt(1 + 8 * w) - 1)/2;
				ll vp2 = queryParent(v, n - x);
				add(vp, vp2, 1 - depths[vp], 1);
				w = w - x * (1+x)/2;
				ll vp3 = queryParent(v, n - x - 1);
				add(vp3, vp3, w, 0);
			}
			continue;

		}else if(a == v){
			ll n = depths[u] - depths[a];
			ll up = queryParent(u, n - 1);

			if(n * (n+1)/2 <= w){
				add(u, up, 1 + depths[u], -1);
			}else{
				ll x = (sqrt(1 + 8 * w) - 1)/2;
				up = queryParent(u, x - 1);
				add(u, up, 1 + depths[u], -1);
				w  = w - x*(1+x)/2;
				up = ps[up][0];
				if(up > 1 && w > 0) add(up, up, w, 0); 
			}

			continue;
		}

		ll nu = depths[u] - depths[a];
		ll up = queryParent(u, nu - 1);
		ll nv = depths[v] - depths[a];
		ll vp = queryParent(v, nv - 1);

		if(nu * (nu + 1)/2 <= w){
			ll co = nu * (nu + 1) / 2;
			add(u, up, 1 + depths[u], -1);
			ll w2 = w - co;

			co = nv*nu + (1 + nv)*nv/2;
			if(co <= w2){
				add(v, vp, nu-depths[a], 1);

			}else{
				ll x = (sqrt((2*nu+1)*(2*nu+1) + 8 * w2) - (2*nu + 1))/2;
				if(x > 0){
					ll vp2 = queryParent(v, nv - x);
					add(vp, vp2, nu - depths[a], 1);
				}
				w2 -= x*nu + x*(x+1)/2;
				ll vp3 = queryParent(v, nv - x - 1);
				if(vp3 > 1 && w2 > 0) add(vp3, vp3, w2, 0);
			}


		}else{
			ll x = (sqrt(1 + 8 * w) - 1)/2;
			up = queryParent(u, x - 1);
			add(u, up, 1 + depths[u], -1);
			w  = w - x*(1+x)/2;
			up = ps[up][0];
			if(up > 1 && w > 0) add(up, up, w, 0); 
		}

	}

	ll res = 0;
	for(int u=1; u<=N; u++){
		res = max(res, querySum(1, u));
	}

	cout<<res<<endl;

	return 0;
}
