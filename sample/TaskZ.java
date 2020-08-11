import java.io.*;
import java.util.*;

public class TaskZ {
    static FastIO io = new FastIO("in.txt");

    public static void main(String[] args) throws Exception {



        io.close();
    }

    static class FastIO {
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
            ps.flush();
        }

        void close(){
            ps.close();
        }

        void print(Object o){
            ps.print(o);
        }

        void println(Object o){
            ps.println(o);
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
