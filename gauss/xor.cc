using Arr = array<ll, 33>;
void gauss(Arr &vs, ll a) {
  for (ll &v : vs) {
    if ((v ^ a) < a) {
      a ^= v;
    }
  }
  if (a) {
    for (ll &v : vs) {
      if ((v ^ a) < v) {
        v ^= a;
      }
    }

    for(ll &v : vs){
      if(v == 0){
        v = a;
        break;
      }
    }
  }
}
