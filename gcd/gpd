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
