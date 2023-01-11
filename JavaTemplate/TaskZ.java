import java.io.*;
import java.util.Arrays;
import java.util.Comparator;
import java.util.Random;
import java.util.StringTokenizer;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class TaskZ {
    static FastIO io = new FastIO("in.txt");

    public static void main(String[] args) throws Exception {
        io.flush();
    }
    
    static class Util {
        static <T> void shuffle(T[] ts) {
            int n = ts.length;
            Random random = new Random();
            for (int i = n - 1; i >= 0; i--) {
                int j = random.nextInt(i + 1);
                if (j != i) {
                    T tmp = ts[i];
                    ts[i] = ts[j];
                    ts[j] = tmp;
                }
            }
        }

        static void shuffle(int[] ts) {
            int n = ts.length;
            Random random = new Random();
            for (int i = n - 1; i >= 0; i--) {
                int j = random.nextInt(i + 1);
                if (j != i) {
                    int tmp = ts[i];
                    ts[i] = ts[j];
                    ts[j] = tmp;
                }
            }
        }

        static void shuffle(long[] ts) {
            int n = ts.length;
            Random random = new Random();
            for (int i = n - 1; i >= 0; i--) {
                int j = random.nextInt(i + 1);
                if (j != i) {
                    long tmp = ts[i];
                    ts[i] = ts[j];
                    ts[j] = tmp;
                }
            }
        }

        static void shuffle(double[] ts) {
            int n = ts.length;
            Random random = new Random();
            for (int i = n - 1; i >= 0; i--) {
                int j = random.nextInt(i + 1);
                if (j != i) {
                    double tmp = ts[i];
                    ts[i] = ts[j];
                    ts[j] = tmp;
                }
            }
        }

        static void shuffle(char[] ts) {
            int n = ts.length;
            Random random = new Random();
            for (int i = n - 1; i >= 0; i--) {
                int j = random.nextInt(i + 1);
                if (j != i) {
                    char tmp = ts[i];
                    ts[i] = ts[j];
                    ts[j] = tmp;
                }
            }
        }
    }
    static class FastIO {
        static StringBuilder sb = new StringBuilder();
        static BufferedReader br;
        static StringTokenizer st;
        static PrintStream ps = new PrintStream(System.out);

        public FastIO(String fname) {
            try {
                File input = new File(fname);
                if (input.exists()) {
                    System.setIn(new FileInputStream(fname));
                }
            }catch (Exception e){
                e.printStackTrace();
            }

            br = new BufferedReader(new
                    InputStreamReader(System.in));
        }

        String next() {
            while (st == null || !st.hasMoreElements()) {
                try {
                    st = new StringTokenizer(br.readLine());
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
            return st.nextToken();
        }

        int nextInt() {
            return Integer.parseInt(next());
        }

        long nextLong() {
            return Long.parseLong(next());
        }

        double nextDouble() {
            return Double.parseDouble(next());
        }

        void flush(){
            System.out.print(sb.toString());
            sb.setLength(0);
        }

        void print(Object o){
            sb.append(o);
        }

        void println(Object o){
            sb.append(o);
            sb.append('\n');
        }

        String nextLine() {
            String str = "";
            try {
                str = br.readLine();
            } catch (IOException e) {
                e.printStackTrace();
            }
            return str;
        }
    }
}
