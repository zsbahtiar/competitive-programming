#include <bits/stdc++.h>

using namespace std;

int main(){
    int bilBul;

        scanf("%d",&bilBul);
        if(bilBul > 0 and bilBul <=9){
            printf("satuan\n");
        }else if(bilBul>=10 and bilBul <=99){
            printf("puluhan\n");
        }else if(bilBul>=100 and bilBul <=999){
            printf("ratusan\n");
        }else if(bilBul>=1000 and bilBul <=9999){
            printf("ribuan\n");
        }else if(bilBul>=10000 and bilBul <=99999){
            printf("puluhribuan\n");
        }

    return 0;
}