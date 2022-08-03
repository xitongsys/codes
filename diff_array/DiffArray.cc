template <class T>
class DiffArray {
public:
    int n;
    vector<T> ds;
    DiffArray(int n, int v = 0)
    {
        this->n = n;
        ds = vector<T>(n, v);
    }

    void add(int b, int e, T v)
    {
        if (b > e) {
            return;
        }
        ds[b] += v;
        if (e + 1 < n) {
            ds[e + 1] -= v;
        }
    }

    vector<T> sum()
    {
        vector<T> res = ds;
        for (int i = 1; i < n; i++) {
            res[i] += res[i - 1];
        }
        return res;
    }
};
