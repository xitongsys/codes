#include <bits/stdc++.h>
using namespace std;
typedef long long ll;

ll mpow(ll a, ll n, ll mod){
	if(n == 0) return 1;
	ll r = mpow(a, n/2, mod);
	r *= r; r %= mod;
	if(n % 2) r *= a;
	return r % mod;
}

ll inv(ll a, ll mod){
	return mpow(a, mod - 2, mod);
}

ll gcd(ll a, ll b){
    if(b==0) return a;
    return gcd(b, a%b);
}


ll extend_gcd(ll a, ll b, ll &x, ll &y) {  
    if (b == 0) {  
	x = 1, y = 0;  
	return a;  
    }  
    else {  
	ll r = extend_gcd(b, a % b, y, x);  
	y -= x * (a / b);  
	return r;  
    }  
}  
ll inv(ll a, ll mod) {  
    ll x, y;  
    extend_gcd(a, mod, x, y);  
    x = (x % mod + mod) % mod;  
    return x;  
}  


int main(){
    cout<<gcd(-2,4)<<endl;
    ll a=2, b=4, x=0, y=0;
    extend_gcd(a, b, x, y);
    cout<<x<<" "<<y<<endl;
    return 0;
}
