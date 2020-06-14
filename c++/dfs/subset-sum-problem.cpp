#include <bits/stdc++.h>
using namespace std;

static const int MAX_N = 20;
int cnt = 0;
int k,n;
int a[MAX_N];

bool dfs(int i,int s)
{
    if(i==n)return s==k;
    //a[i]を使わない場合
    if(dfs(i+1,s))return true;
    //a[i]を使う場合
    if(dfs(i+1,s+a[i]))return true;
    return false;
}

int main()
{
    cin>>n>>k;
    for(int i=0;i<n;i++)cin>>a[i];
    cout<<(dfs(0,0)?"yes":"no")<<endl;
    return 0;
}