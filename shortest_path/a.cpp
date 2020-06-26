//leetcode 743
#include <bits/stdc++.h>
using namespace std;
#define LL long long


int inf=100000000;
//dijkstra
vector<int> dijkstra(vector<vector<int>>& gra, int N, int K){
	vector<int> dis(N+1, inf);
	vector<int> flags(N+1, 0);
	dis[K]=0;
	for(int i=0; i<N; i++){
		int minv=inf, minidx=-1;
		for(int j=1; j<=N; j++){
			if(flags[j]==1) continue;
			if(dis[j]<=minv){
				minv=dis[j]; minidx=j;
			}
		}
		flags[minidx]=1;
		dis[minidx]=minv;
		for(int j=1; j<=N; j++){
			if(gra[minidx][j]>=0){
				dis[j] = min(dis[j], minv + gra[minidx][j]);
			}
		}
	}
	return dis;
}


//dijkstra with priority_queue
struct CMP {
	bool operator()(pair<int,int>& a, pair<int,int>& b){
		return a.second > b.second;
	}
};
vector<int> dijkstra2(vector<vector<int>>& gra, int N, int K){
	vector<int> dis(N+1, inf);
	priority_queue<pair<int,int>, vector<pair<int,int>>, CMP> pq;
	pq.push(pair<int,int>(K, 0));
	while(pq.size()>0){
		auto p=pq.top(); pq.pop();
		int u=p.first, d=p.second;
		if(dis[u]!=inf) continue;
		dis[u]=d;
		for(int v=1; v<=N; v++){
			if(gra[u][v]>=0 && dis[v]==inf){
				pq.push(pair<int,int>(v, dis[u]+gra[u][v]));
			}
		}
	}
	return dis;
}

//bellman_ford
vector<int> bellman_ford(vector<vector<int>>& edges, int N, int K){
	vector<int> dis(N+1, inf);
	dis[K]=0;
	for(int k=0; k<N; k++){
		for(auto e : edges){
			int u=e[0], v=e[1], w=e[2];
			dis[v] = min(dis[v], dis[u] + w);
		}
	}
	return dis;
}

//floyd
vector<int> floyd(vector<vector<int>>& gra, int N, int K){
	vector<vector<int>> dis(N+1, vector<int>(N+1, inf));
	for(int i=1; i<=N; i++){
		dis[i][i]=0;
		for(int j=1; j<=N; j++){
			if(gra[i][j]>=0){
				dis[i][j]=gra[i][j];
			}
		}
	}

	for(int k=1; k<=N; k++){
		for(int i=1; i<=N; i++){
			for(int j=1; j<=N; j++){
				dis[i][j] = min(dis[i][j], dis[i][k] + dis[k][j]);
			}
		}
	}

	vector<int> res(1,0);
	for(int i=1; i<=N; i++){
		res.push_back(dis[K][i]);
	}
	return res;
}


class Solution {
public:
    int networkDelayTime(vector<vector<int>>& times, int N, int K) {
		vector<vector<int>> gra(N+1, vector<int>(N+1, -1));
		for(auto a : times){
			int u=a[0], v=a[1], w=a[2];
			gra[u][v]=w;
		}
		vector<int> dis=floyd(gra, N, K);
		int ans=0;
		for(int i=1; i<=N; i++){
			ans=max(ans, dis[i]);
		}
		if(ans==inf) ans=-1;
		return ans;

    }
};

int main(){
	Solution s;
	vector<vector<int>> times={{1,2,1}, {2,3,2}, {1,3,4}};
	cout<<s.networkDelayTime(times, 3, 1)<<endl;
	return 0;
}
