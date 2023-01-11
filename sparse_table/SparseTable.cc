template <class T>
class SparseTable {
public:
    int n, m;
    vector<vector<T>> sparse_table;
    vector<T> as;
    T(*cmp)
    (const T& a, const T& b);

    SparseTable(const vector<T>& vs, T (*cmp)(const T& a, const T& b))
    {
        this->cmp = cmp;
        as = vs;
        n = vs.size();
        m = int(log2(n * 2));

        sparse_table = vector<vector<T>>(n, vector<T>(m, 0));

        for (int i = 0; i < n; i++) {
            sparse_table[i][0] = as[i];
        }
        for (int j = 1; j < m; j++) {
            for (int i = 0; i < n; i++) {
                sparse_table[i][j] = sparse_table[i][j - 1];
                int ii = i + (1 << (j - 1));
                if (ii < n) {
                    sparse_table[i][j] = cmp(sparse_table[i][j - 1], sparse_table[ii][j - 1]);
                }
            }
        }
    }

    T get(int b, int e)
    {
        T res = sparse_table[b][0];
        while (b <= e) {
            int j = 0;
            while (b + (1 << j) - 1 <= e) {
                j++;
            }
            j--;
            res = cmp(res, sparse_table[b][j]);
            b += 1 << j;
        }
        return res;
    }
};
