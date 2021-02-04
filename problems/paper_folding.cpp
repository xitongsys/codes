/* 从右向左不断折纸,折痕形状是V或者^,
 * 给定折纸次数n,以及m,
 * 输出折纸n此后,第m条(m从0开始)
 * 折痕的形状(V or ^)
 *
 * constraints：
 * 1 < n <= 63 
 * 0 <= m < 2**n - 1
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


char shape(int n, long long m){
	int a = 0;
	long long i = pow(2, n-1) - 1;
	while(true){
		if(i == m) return a == 0 ? 'V' : '^';
		if(i < m){
			a = 1;
			m -= i + 1;
		}else{
			a = 0;
		}
		i = (i-1)/2;
	}
	return 0;
}

int main(){
	for(int n=2; n<10; n++){
		for(int m=0; m<pow(2, n)-1; m++){
			cout<<shape(n, m);
		}
		cout<<endl;
	}
	return 0;
}
