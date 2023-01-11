import java.util.Arrays;
import java.util.Comparator;
import java.util.Random;

public class KthQuery {
    int n, m, bl;
    long[] as;
    int[] mp2old, mp2new;
    int[][] buckets;

    public KthQuery(long[] as){
        n = as.length;
        this.as = as;

        long[][] ps = new long[n][2];
        for(int i=0; i<n; i++){
            ps[i][0] = i;
            ps[i][1] = as[i];
        }

        Arrays.sort(ps, new Comparator<long[]>() {
            @Override
            public int compare(long[] o1, long[] o2) {
                if(o1[1] > o2[1]){
                    return 1;
                }else if(o1[1] < o2[1]){
                    return -1;
                }else{
                    return (int)(o1[0] - o2[0]);
                }
            }
        });

        mp2old = new int[n];
        mp2new = new int[n];
        for(int i=0; i<n; i++){
            mp2old[i] = (int)ps[i][0];
            mp2new[(int)ps[i][0]] = i;
        }

        //System.out.println(mp2old);
        //System.out.println(mp2new);

        bl = (int)Math.sqrt(n);
        m = n / bl + 1;


        buckets = new int[m][n];
        for(int i=0; i<n; i++){
            int bi = i / bl;
            buckets[bi][mp2new[i]]++;
        }

        for(int i=0; i<m; i++){
            for(int j=1; j<n; j++){
                buckets[i][j] += buckets[i][j-1];
            }
        }
    }

    long query(int l, int r, int k){
        int lv = 0, rv = n - 1;
        while(lv <= rv) {
            int mv = lv + (rv - lv) / 2;
            int count = 0;
            for (int bi = 0; bi < m; bi++) {
                int lb = bi * bl, rb = bi * bl + bl - 1;
                if(lb >= l && rb <= r) {
                    count += buckets[bi][mv];

                }else if(rb < l || lb > r){
                    continue;

                }else{
                    for(int i=lb; i<=rb; i++){
                        if(i >= l && i <= r) {
                            int ni = mp2new[i];
                            if(ni <= mv) count++;
                        }
                    }
                }
            }

            if(count >= k) rv = mv - 1;
            else lv = mv + 1;
        }

        return as[mp2old[lv]];
    }

    long queryTest(int l, int r, int k){
        long[] cs = new long[r - l + 1];
        for(int i=l; i<=r; i++){
            cs[i - l] = as[i];
        }

        Arrays.sort(cs);
        return cs[k-1];
    }


    static public void main(String[] args) {
        Random rnd = new Random();
        int N = 1000;
        long[] as = new long[N+1];
        for(int i=1; i<=N; i++){
            as[i] = 1000;
            //as[i] = rnd.nextInt();
        }

        KthQuery kq = new KthQuery(as);

        while(true){
            int l = rnd.nextInt(N+1), r = rnd.nextInt(N+1);
            if(r < l){
                int t = l; l = r; r = t;
            }
            int k = rnd.nextInt(r - l + 1) + 1;
            long a = kq.query(l, r, k);
            long b = kq.queryTest(l, r, k);
            //System.out.println(a + " " + b);
            if(a != b)
                System.out.println(l + " " + r + " " + k);
        }
    }
}
