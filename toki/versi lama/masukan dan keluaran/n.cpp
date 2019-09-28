#include <bits/stdc++.h>

using namespace std;

int main(){
    int x=0,y=0,z=0;
    string hasil;

      scanf("%d",&x);
      scanf("%d",&y);
      scanf("%d",&z);
  
    if((x > 0 and x <= 1000000) or (y > 0 and z <= 1000000) or (z > 0 and z <= 1000000))
    {
        
        if((x*x+y*y)!=(z*z)){
            printf("FALSE\n");
        }else{
            printf("TRUE\n");
        }
    }else{
            printf("FALSE\n");
        }
    return 0;
}