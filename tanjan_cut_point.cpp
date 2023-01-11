vector<int> dfsClock(MX, -1), minPar(MX, INT_MAX), vis(MX, 0);
void tanjan(int u, int clock, int p){
	dfsClock[u] = clock;
	minPar[u] = clock;
	vis[u] = 1;
	int cn = 0;
	for(int ei : g[u]){
		Edge &e = es[ei];
		int v = (e.u == u ? e.v : e.u);
		if(vis[v] && v != p){
			minPar[u] = min(minPar[u], dfsClock[v]);
		}else if(vis[v] == 0){
			cn++;
			tanjan(v, clock + 1, u);
			minPar[u] = min(minPar[u], minPar[v]);
			if(minPar[v] >= dfsClock[u] && u != 1) all.insert(u);
		}
	}
	if(cn > 0 && u == 1) all.insert(u);
}

