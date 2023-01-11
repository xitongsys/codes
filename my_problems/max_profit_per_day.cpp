/*
 * Give you n days stock price [p0, p1, p2, ... pn], choose one day to buy and choose another day(a day after the buy day) to sell. Please get the max profit per day you can get(err < 1e-4).
 *
 * constraints:
 * 1 < n < 1e6
 * 0 < pi < 1e5
 */

#include <algorithm>
#include <assert.h>
#include <bitset>
#include <cmath>
#include <cstdio>
#include <cstdlib>
#include <iostream>
#include <list>
#include <map>
#include <queue>
#include <set>
#include <sstream>
#include <stack>
#include <string>
#include <tuple>
#include <vector>
#include <random>
using namespace std;

double maxProfitPerDay_NLogN(vector<double> ps){
	int n = ps.size();
	double l = -1e6, r = 1e6;
	while(r - l > 1e-4){
		double m = (l + r) / 2;
		double p = ps[0];
		bool f = false;
		for(int i=0; i<n && !f; i++){
			if(p < ps[i]){
				f = true;
			}else{
				p = ps[i];
			}
			p += m;
		}
		if(f){
			l = m;
		}else{
			r = m;
		}
	}
	return (l+r)/2;
}

double maxProfitPerDay_N(vector<double> ps){
	int n = ps.size();
	double ans = ps[1] - ps[0];
	for(int i=0; i+1<n; i++){
		ans = max(ans, ps[i+1] - ps[i]);
	}
	return ans;

}

double check(vector<double> ps){
	int n = ps.size();
	double ans = -1e6;
	for(int i=0; i<n; i++){
		for(int j=i+1; j<n; j++){
			ans = max(ans, (ps[j] - ps[i])/(j-i));
		}
	}
	return ans;
}

void test(){
	for(int t=0; t<1000; t++){
		int N = 1000;
		default_random_engine e;
		e.seed(t);
		vector<double> ps;
		for(int i=0; i<N; i++){
			double p = e() % 10000000;
			ps.push_back(p / 100);
		}
		cout<<check(ps)<<" "<<maxProfitPerDay_NLogN(ps)<<" "<<maxProfitPerDay_N(ps)<<endl;
	}
}

int main(){
	vector<double> ps = {6,5,3,2,1};
	cout<<check(ps)<<" "<<maxProfitPerDay_NLogN(ps)<<" "<<maxProfitPerDay_N(ps)<<endl;
	//test();
	return 0;
}
