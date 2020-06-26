#include<iostream>
#include<vector>
using namespace std;
class BinTree:vector<int>
{
    public:
        explicit BinTree(int k=0)//默认初始化一个能保存k个元素的空树状数组
        {
            assign(k+1,0);//有效下标从1开始，0仅作逻辑用处
        }
        int lowbit(int k)
        {
            return k&-k;
            //也可写作x&(x^(x–1))
        }
        int sum(int k)//求第1个元素到第n个元素的和
        {
            return k>0?sum(k-lowbit(k))+(*this)[k]:0;
        }
        int last()//返回最后一个元素下标
        {
            return size()-1;
        }
        void add(int k,int w)//为节点k加上w
        {
            if(k>last())return;
            (*this)[k]+=w;
            add(k+lowbit(k),w);
        }
};
int main()
{
    BinTree test(123);
    test.add(27,72);
    cout<<test.sum(26)<<' '<<test.sum(27)<<' '<<test.sum(123);
}
