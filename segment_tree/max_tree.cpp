struct Node {
    int v, d;
    int l,r;
    Node *left, *right;
    Node(int ll, int rr){
	l = ll; r=rr;
	left = NULL; right=NULL;
	v = 0;d=0;
	if(ll==rr) return;
	int m = l + (r-l)/2;
	left = new Node(l, m);
	right = new Node(m+1, r);
    }

    void pushdown(){
	if(d!=0){
	    if(left!=NULL){
		left->v += d;
		left->d += d;
	    }
	    if(right!=NULL){
		right->v += d;
		right->d += d;
	    }
	    d = 0;
	}
    }

    void add(int ll, int rr, int dd){
	if(ll==l && rr==r){
	    d += dd;
	    v += dd;
	    return;
	}
	
	pushdown();
	int m = l + (r-l)/2;
	if(rr<=m){
	    left->add(ll,rr,dd);
	}else if(ll>m){
	    right->add(ll,rr,dd);
	}else{
	    left->add(ll, m, dd);
	    right->add(m+1, rr, dd);
	}

	v = max(left->v, right->v);
    }

    int query(int ll, int rr){
	if(l==ll && r==rr){
	    return v;
	}

	pushdown();
	int m = l + (r-l)/2;
	if(rr<=m){
	    return left->query(ll, rr);
	}else if(ll>m){
	    return right->query(ll, rr);
	}else{
	    return max(left->query(ll, m), right->query(m+1,rr));
	}
	return 0;
    }
    
};
