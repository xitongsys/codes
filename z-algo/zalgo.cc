/**
 * https://www.scaler.com/topics/data-structures/z-algorithm/
 */

#include <bits/stdc++.h>
using namespace std;

template <class T>
vector<int> z_algo(const T& arr)
{
    int n = arr.size();
    vector<int> zarr(n, 0);

    int left, right, k;
    left = right = 0;

    for (int i = 1; i < n; i++) {
        if (i > right) {
            left = right = i;
            while (right < n && arr[right - left] == arr[right]) {
                right++;
            }
            zarr[i] = right - left;
            right--;

        } else {
            k = i - left;
            if (zarr[k] < right - i + 1) {
                zarr[i] = zarr[k];
            } else {
                left = i;
                while (right < n && arr[right - left] == arr[right]) {
                    right++;
                }
                zarr[i] = right - left;
                right--;
            }
        }
    }
    return zarr;
}

int main()
{
    string text = "faabbcdeffghiaaabbcdfgaabf";
    string pattern = "aabb";

    string cs = pattern + "$" + text;

    vector<int> zarr = z_algo(cs);
    for (int i = pattern.size() + 1; i < cs.size(); i++) {
        if (zarr[i] == pattern.size()) {
            cout << i - pattern.size() - 1 << endl;
        }
    }

    return 0;
}