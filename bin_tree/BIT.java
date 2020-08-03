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
            if(idx > N){
                return;
            }

            C[idx] += a;
            //io.out(idx + " " + lowbit(idx) + "\n");
            add(idx + lowbit(idx), a);
        }

        long sum(int idx){
            if(idx <= 0) return 0;
            return sum(idx - lowbit(idx)) + C[idx];
        }

        long query(int l, int r){
            long ans = sum(r);
            if(l > 1){
                ans -= sum(l - 1);
            }
            return ans;
        }
    }
