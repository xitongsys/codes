struct Node {
	ll v;
	ll lazy;
	ll lv, rv;
	Node *left ,*right;
	Node(ll l, ll r){
		v=0; lazy=0;
		lv=l; rv=r;
		left=NULL; right=NULL;
		if(lv>=rv) return;
		ll m=lv + (rv-lv)/2;
		left=new Node(l, m);
		right=new Node(m+1, r);
	}

	void output(){
		cout<<lv<<" "<<rv<<" "<<v<<endl;
		if(left!=NULL) left->output();
		if(right!=NULL) right->output();
	}

	void pushdown(){
		ll mv = lv + (rv - lv)/2;
		if(lazy!=0){
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
	}

	void add(ll l, ll r, ll vv){
		//cout<<l<<" "<<r<<" "<<lv<<" "<<rv<<endl;
		if(lv==l && rv==r){
			v += vv * (rv - lv + 1);
			lazy += vv;
			return;
		}
		pushdown();
		ll m=lv + (rv - lv)/2;
		if(r<=m){
			left->add(l, r, vv);
		}else if (l>m){
			right->add(l, r, vv);
		}else{
			left->add(l, m, vv);
			right->add(m+1, r, vv);
		}
		v=left->v + right->v;
	}

	ll query(ll l, ll r){
		if(l==lv && r==rv){
			return v;
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
