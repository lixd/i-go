 #include<unistd.h>
 #include<stdio.h>
 #include<stdlib.h>
 #include<sys/types.h>
 int main(void){
 	pid_t pid;
 	int n=0,m=30;
// 	主进程一直fork子进程，直到子进程达到1W个
 	while(1){
    pid=fork();
    if(pid==0){
      break;
    }
    else if(pid>0){
      printf(" %d\n",n++);
      if (n>10000){ // 限制最多运行10000个进程
        break;
      }
    }
    else{
      exit(1);
    }
  }
  // 子进程休眠30秒后退出
   while(m--){
    printf("sleep %d\n",m);
    sleep(1);
   }
   return 0;
 }
