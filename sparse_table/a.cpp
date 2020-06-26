#include <bits/stdc++.h>
using namespace std;

vector<int> nums={0,1,2,3,4,5};
vector<vector<int>> dp;


void init(){
	int n=nums.size();
	dp=vector<vector<int>>(n, vector<int>(20, 0));
	for(int i=0; i<n; i++){
		dp[i][0] = nums[i];
	}
	for(int j=1; j<20; j++){
		for(int i=0; i<n; i++){
			dp[i][j] = dp[i][j-1]; 
			if((i + 1<<(j-1)) < n){
				dp[i][j]=max(dp[i][j], dp[i + (1<<(j-1))][j-1]);
			}
		}
	}
}

int query(int b, int e){
	int k=0;
	while((1<<(k+1))<=e-b+1)k++;
	int res=max(dp[b][k], dp[e-(1<<k)+1][k]);
	return res;
}

int main(){
	init();
	int n=nums.size();
	for(int i=0; i<n; i++){
		for(int j=i; j<n; j++){
			cout<<i<<" "<<j<<" "<<query(i,j)<<endl;
		}
	}
	return 0;
}

