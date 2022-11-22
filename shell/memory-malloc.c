/*
usage: cc mem.c -o mem.out 后 使用./mem.out 100 消耗对应数字MB单位的内存，释放时杀掉对应进程或者等待超时即可
没有 gcc 的话使用 yum install gcc-c++ -y 安装
*/
#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>

#define UNIT (1024*1024)

int main(int argc, char *argv[])
{
        long long i = 0;
        int size = 0;
        int timeout = 0;

        if (argc != 3) {
                printf(" === argc must 3\n");
                return 1;
        }
        size = strtoull(argv[1], NULL, 10);
        if (size == 0) {
                printf(" argv[1]=%s not good\n", argv[1]);
                return 1;
        }
        timeout = strtoull(argv[2], NULL, 10);
        if (timeout == 0) {
                printf(" argv[2]=%s not good\n", argv[2]);
                return 1;
        }


        char *buff = (char *) malloc(size * UNIT);
        if (buff)
                printf(" we malloced %d Mb\n", size);
        buff[0] = 1;

        for (i = 1; i < (size * UNIT); i++) {
                if (i%1024 == 0)
                        buff[i] = buff[i-1]/8;
                else
                        buff[i] = i/2;
        }
        printf("sleep %d s\n", timeout);
        sleep(timeout);
        return 0;
}
