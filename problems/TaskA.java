import java.util.Arrays;
/**
 * This problems are written for interview by xitongsys.
 */

public class TaskA {
    static void outputArray(int[] as){
        for(int a : as){
            System.out.print(a + " ");
        }
        System.out.println("");
    
    }
    public static void main(String[] args) throws Exception {
        Solution sol = new Solution();

        int[][] cases = new int[][]{
            new int[]{2,2,2,2,3,3,3},
            new int[]{2,2,2,3,3,3,3},
            new int[]{1,1,2,2,2,3},
            new int[]{2,2,2,2,2,3,3,3},
            new int[]{2,2,3,3,3,5,6}
        };

        for(int[] c : cases){
            int[] ans = sol.makeZigzag(c);
            System.out.println(sol.canZigzag(c));
            outputArray(ans);
            System.out.println("");
        }
    }
}

class Solution {
    /**
     * Give you an array with distinct numbers, please arrange them to zigzag shape. If there are multiple answers, you can give any one of them.
     * Constraints:
     * 3 < nums.length < 1e5
     * 
     * Example: 
     * [1,2,3,4,5] -> [1,3,2,5,4]
     */
    public int[] makeZigzagDistinct(int[] nums) {
        int n = nums.length;
        Arrays.sort(nums);
        int m = n / 2;
        int[] ans = new int[n];
        int i = 0, j = n - m, k= 0;
        int f = 0;
        while(k < n){
            if(f == 0){
                ans[k] = nums[i];
                i++;
            }else{
                ans[k] = nums[j];
                j++;
            }
            k++;
            f ^= 1;
        }
        return ans;
    }

    /**
     * Give you an array with numbers(may have duplicated numbers), please arrange them to zigzag shape. If there are multiple answers, you can give any one of them. If there is no answer, return an empty array.
     * 
     * Constraints:
     * 3 < nums.length < 1e5
     * 
     * Example: 
     * [1,2,3,3,5] -> [1,3,2,5,3]
     */
    public int[] makeZigzag(int[] nums){
        int n = nums.length;
        int[] ans = new int[n];
        int m = (n + 1) / 2;
        int i = 0, j = n - m, k = 0;
        int f = 1;
        while(k < n){
            if(f == 1){
                ans[k] = nums[j];
                j++;
            }else{
                ans[k] = nums[i];
                i++;
            }
            k++;
            f ^= 1;
        }

        boolean ok = true;
        for(i=0; i<n && ok; i+=2){
            if(i-1 >= 0 && ans[i] <= ans[i-1]){
                ok = false;
            }
            if(i + 1 < n && ans[i] <= ans[i+1]){
                ok = false;
            }
        }

        if(ok) return ans;

        m = n / 2;
        f = 1;
        i = 0; j = n - m; k = 0;
        while(k < n){
            if(f == 1){
                ans[k] = nums[i];
                i++;
            }else{
                ans[k] = nums[j];
                j++;
            }
            k++;
            f ^= 1;
        }

        ok = true;
        for(i=1; i<n && ok; i+=2){
            if(i-1 >= 0 && ans[i] <= ans[i-1]){
                ok = false;
            }
            if(i + 1 < n && ans[i] <= ans[i+1]){
                ok = false;
            }
        }

        if(ok) return ans;
        return new int[]{};
    }
   
    /**
     * Give you an ordered array with numbers(may have duplicated numbers), please check if it can be arranged to zigzag shape.
     * 
     * Constraints:
     * 3 < nums.length < INF 
     * (upper bound has no constraints, it may be very large and you can't iterate them in an accepted time. But you can get the length of the array and any value by index.)
     * 
     * Example: 
     * [1,2,3,3,5] -> true
     * [1,2,2,2,2] -> false
     */
    public boolean canZigzag(int[] nums){
        int n = nums.length;
        //Arrays.sort(nums);

        int m = (n + 1) / 2;
        int bi = n - m;
        if(nums[bi-1] < nums[bi]) return true;
        m = n / 2;
        bi = n - m;
        if(nums[bi-1] < nums[bi]) return true;

        m = n / 2;
        bi = n - m;
        int t = nums[m];
        int[] bs = binSearch(nums, t);
        int b = bs[0], e = bs[1];
        
        if(e - bi < b) return true;
        return false;
    }

    public int[] binSearch(int[] nums, int t){
        int n = nums.length;
        int l = 0, r = n - 1;
        int[] ans = new int[]{-1, -1};
        while(l <= r){
            int m = l + (r-l)/2;
            if(nums[m] >= t) r = m - 1;
            else l = m + 1;
        }
        ans[0] = l;

        l = 0; r = n - 1;
        while(l <= r){
            int m = l + (r-l) / 2;
            if(nums[m] <= t) l = m + 1;
            else r = m - 1;
        }
        ans[1] = r;
        return ans;
    }


}
