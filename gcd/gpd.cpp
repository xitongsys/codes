//greast prime divisor algorithm
long long gpd[nax];

void init()
{
    memset(gpd,-1,sizeof(gpd));
    for(int i=2;i<nax;i++)
    {
        if(gpd[i]!=-1)
        continue;
        for(int j=i;j<nax;j+=i)
        {
            //good point
            gpd[j]=i;
        }
    }
}

void decompose(){
    vector<long long> v;//all prime divisor
    while(x!=1)
    {
        int  y= gpd[x] ;
        v.push_back(gpd[x]);
        while(x%y==0) x/=y;
    }
}

