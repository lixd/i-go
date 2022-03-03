#define _CRT_SECURE_NO_WARNINGS 1
#include<stdio.h>
#include<stdlib.h>

//模2除2法，该方法不能判断负数，因为负数模2的结果总是0
int count_one_bits(int n)
{
    int count = 0;
    while (n)
    {
        if (n % 2 == 1)//模2可以得到最低位的数，如果模2的结果是1，则将count++
        {
            count++;
        }
        n = n / 2;//除2可以去掉最末位的数
    }
    return count;
}
int main()
{
    int i = 0;
    int num = 0;
    scanf("%d", &num);
    int ret = count_one_bits(num);//将实参num传递到函数中
    printf("%d\n", ret);
    system("pause");
    return 0;
}
//判断二进制数中1个个数 https://blog.csdn.net/windyJ809/article/details/79821148

int count1(unsigned int n) {
    int count = 0;
    while (n > 0) { // 所有1都换成0跳出循环
        n = n & (n - 1); // 1换0
        count++;
    }
    return count;
}
/*
1111 1110 1110
1110 1101 1100
1100 10111 1000
1000 0111 0000
*/
int count0(unsigned int n) {
    int count = 0;
    while (n != -1) { // 所有0都换成1跳出循环
        n = n | (n + 1); // 0换1
        count++;
    }
    return count;
}
/*
1001 1010 1010
1010 1011 1011
1011 1110 1111
*/

