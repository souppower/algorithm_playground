#include <iostream>
#include <cstdio>
#include <algorithm>
#include <cmath>
#include <vector>
using namespace std;

long long cnt;
int l;
int A[1000000];
int n;
vector<int> G;

void insertionSort(int A[], int n, int g) {
    for (int i = 0; i < n; i++) {
        int v = A[i]; // 基準値
        int j = i - g; // 基準のindex - ギャップのindex
        while(j >= 0 && A[j] > v) {
            A[j+g] = A[j];
            j -= g;
            cnt++;
        }
        A[j+g] = v;
    }
}

void shellSort(int A[], int n) {
    // gapをvecotrに詰めていく
    for (int h = 1; ; ) {
        if (h > n) break;
        G.push_back(h);
        h = 3*h + 1;
    }

    // gapの大きい方から順番にinsertionSortに渡していく
    for (int i = G.size()-1; i >= 0; i--) {
        insertionSort(A, n, G[i]);
    }
}

int main() {
    cin >> n;
    for(int i = 0; i < n; i++) cin >> A[i];
    shellSort(A, n);
    for (int i = 0; i < n; i++) printf("%d", A[i]);
}