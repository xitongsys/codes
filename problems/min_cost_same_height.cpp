/*
 * Give an array H[h0, h1, h2, ... hn] where hi is the number of cube at position i.
 * Move a cube from i to i-1 or i+1 has a cost 1.
 * You need to make all the positions have the same height by minimal cost.
 * Please give the cost.
 *
 * constraints:
 *
 * 1 < n < 1e6
 * 0 <= hi < 1e9
 * sum(hi) % n == 0
*/

#include <vector>
#include <cstdio>
#include <cstdlib>
#include <string>
#include <cstring>
#include <iostream>
#include <map>
#include <set>
#include <stack>
#include <queue>
#include <bitset>
#include <algorithm>
#include <random>
#include <chrono>
#include <functional>
using namespace std;


long long minCostOfSameHeight(vector<int> H){
	int n = H.size();
	long long s = 0;
	for(int i=0; i<n; i++){
		s += H[i];
	}
	long long t = s / n;
	long long cost = 0;
	long long w = 0;
	for(int i=0; i<n; i++){
		cost += abs(w);
		w += H[i];
		w -= t;
	}
	return cost;
}

int main(){
	vector<int> H = {1,2,3,4,5};
	cout<<minCostOfSameHeight(H)<<endl;
	return 0;
}
