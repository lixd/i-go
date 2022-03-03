 #include<unistd.h>
 #include<stdio.h>
// 测试fork调用
 int main(void){
 	pid_t pid;
 	pid=fork();
 	if(pid==0)	{
 		printf("child process\n");
 		printf("child  pid is %d\n",getpid());
 		printf("child  ppid is %d\n",getppid());
 	}	else if(pid>0)	{
 		printf("parent process\n");
 		printf("parent pid is %d\n",getpid());
 		printf("parent ppid is %d\n",getppid());
 	}else	{
 		printf("fork error\n");
 	}
 	return 0;
 }
// gcc -o fork fork.c
// man fork
