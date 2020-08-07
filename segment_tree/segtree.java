   static class Node {
        public int l, r, v, lazy;
        public Node left, right;
        public Node(int l, int r){
            this.l = l; this.r = r;
            if(l == r) return;
            int m = l + (r - l) / 2;
            left = new Node(l, m);
            right = new Node(m+1, r);
        }

        public void push(){
            if(lazy == 0) return;
            v += lazy * (r - l + 1);
            if(left!=null) left.lazy += lazy;
            if(right!=null) right.lazy += lazy;
            lazy = 0;
        }

        public void add(int l1, int r1, int v1){
            if(l1 == l && r1 == r){
                lazy += v1;
                return;
            }
            v += (r1 - l1 + 1) * v1;

            int m = l + (r - l) / 2;
            if(r1 <= m){
                left.add(l1, r1, v1);
            }else if(l1 > m){
                right.add(l1, r1, v1);
            }else{
                left.add(l1, m, v1);
                right.add(m+1, r1, v1);
            }
        }

        public int query(int l1, int r1){
            push();
            if(l1 == l && r1 == r){
                return v;
            }

            int m = l + (r-l)/2;
            if(r1 <= m){
                return left.query(l1, r1);
            }else if(l1 > m){
                return right.query(l1, r1);
            }else{
                return left.query(l1, m) + right.query(m+1, r1);
            }
        }
    }
