import java.util.*;

/**
 * This problems are written for interview by xitongsys.
 */

public class TaskB {
    static void outputArray(int[] as) {
        for (int a : as) {
            System.out.print(a + " ");
        }
        System.out.println("");
    }

    public static void main(String[] args) throws Exception {
        SolutionB sol = new SolutionB();

        int[] ans = sol.splitTreeToTwo(3, new int[]{-1, 0, 0}, new int[]{2,3,5});
        outputArray(ans);

        int[][] ansK = sol.splitTreeToK(3, new int[]{-1, 0, 0}, new int[]{2,3,5}, 2);
        for(int[] vs : ansK){
            outputArray(vs);
        }
    }
}

class SolutionB {
    /**
     * Give you a tree with N nodes labeled  0 to N-1. Root is 0. Every node i has a value V[i], which V[i] > 0. You delete an edge to make the sum of values of the two seprated trees be same. If there is such an edge, return the edge [from, to] (oreder doesn't matter). Or return [].
     * 
     * Example:
     * 
     * N = 3, parents = [-1, 0, 0], values = [2, 3, 5]
     * 
     * n0v2 means node[0], value = 2
     * n1v3 means node[1], value = 3
     * n2v5 means node[2], value = 5
     * 
     *   n0v2
     *  /    \
     * n1v3   n2v5
     * 
     * You can delete the edge between 2 and 5, and the seperated trees are
     * 
     *   n0v2       n2v5
     *  /
     * n1v3   
     * 
     *            
     */

    ArrayList<Integer>[] g;
    int[] sumValues;

    int dfs(int u){
        for(int v : g[u]){
            sumValues[u] += dfs(v);
        }
        return sumValues[u];
    }

    public int[] splitTreeToTwo(int N, int[] parents,  int[] values) {
        g = new ArrayList[N];
        for(int i=0; i<N; i++){
            g[i] = new ArrayList();
        }

        for(int u=0; u<N; u++){
            int p = parents[u];
            if(p >= 0){
                g[p].add(u);
            }
        }

        sumValues = new int[N];
        for(int i=0; i<N; i++){
            sumValues[i] = values[i];
        }

        dfs(0);

        if(sumValues[0] % 2 == 0){
            int t = sumValues[0] / 2;
            for(int u=0; u<N; u++){
                if(sumValues[u] == t){
                    return new int[]{parents[u], u};
                }
            }
        }

        return new int[]{};
    }

    /**
     * Give you a tree with N nodes labeled  0 to N-1. Root is 0. Every node i has a value V[i], which V[i] > 0. You delete some edges to split the tree into K subtrees. You should make the sum of values of the subtrees be same. If there is solution, return the edges [[from_i, to_i]...] (oreder doesn't matter). Or return [].
     *
     * Constraints:
     * N < 1e5, V[i] > 0, H(Height of the tree) < 100, 1 < K < N
     *
     * 
     * Example:
     * 
     * N = 3, parents = [-1, 0, 0], values = [2, 3, 5], K = 2
     * 
     * n0v2 means node[0], value = 2
     * n1v3 means node[1], value = 3
     * n2v5 means node[2], value = 5
     * 
     *   n0v2
     *  /    \
     * n1v3   n2v5
     * 
     * You can delete the edge between 2 and 5, and the seperated trees are
     * 
     *   n0v2       n2v5
     *  /
     * n1v3   
     * 
     *            
     */
    public int[][] splitTreeToK(int N, int[] parents,  int[] values, int K) {
        g = new ArrayList[N];
        for(int i=0; i<N; i++){
            g[i] = new ArrayList();
        }

        for(int u=0; u<N; u++){
            int p = parents[u];
            if(p >= 0){
                g[p].add(u);
            }
        }

        sumValues = new int[N];
        for(int i=0; i<N; i++){
            sumValues[i] = values[i];
        }

        dfs(0);

        ArrayList<int[]> ans = new ArrayList<>();

        if(sumValues[0] % K == 0){
            Map<Integer, Set<Integer>> mp = new HashMap();
            for(int u=0; u<N; u++){
                int s = sumValues[u];
                if(!mp.containsKey(s)){
                    mp.put(s, new HashSet<>());
                }
                mp.get(s).add(u);
            }

            int t = sumValues[0] / K;
            int c = 0;
            while(c < K){
                if(!mp.containsKey(t) || mp.get(t).size()==0){
                    break;
                }

                int u = mp.get(t).iterator().next();
                int p = u;
                while(p >= 0){
                    int sp = sumValues[p];
                    sumValues[p] -= t;
                    int nsp = sp - t;
                    mp.get(sp).remove(p);
                    if(!mp.containsKey(nsp)){
                        mp.put(nsp, new HashSet<>());
                    }
                    mp.get(nsp).add(p);
                    p = parents[p];
                }
                c++;
            }

            if(c < K){
                ans.clear();
            }
        }

        int ne = ans.size();
        int[][] ret = new int[ne][2];
        for(int i=0; i<ne; i++){
            ret[i] = ans.get(i);
        }
        return ret;
    }

}
