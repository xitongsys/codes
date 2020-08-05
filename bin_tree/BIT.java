    //index from 1
    class BIT {
        long[] C;
        int N;
        public BIT(int n){
            C = new long[n + 1];
            N = n;
        }

        int lowbit(int a){
            return ((a ^ (a - 1)) + 1) >> 1;
        }

        void add(int idx, int a){
            while(idx <= N){
                C[idx] += a;
                idx += lowbit(idx);
            }
        }

        long sum(int idx){
            long ans = 0;
            while(idx > 0){
                ans += C[idx];
                idx -= lowbit(idx);
            }
            return ans;
        }

        long query(int l, int r){
            long ans = sum(r);
            if(l > 1){
                ans -= sum(l - 1);
            }
            return ans;
        }
    }
