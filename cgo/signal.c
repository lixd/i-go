#include <stdio.h>
#include <unistd.h>
#include <signal.h>
#include <errno.h>
/*
使用 sigaction 函数：
 signal 函数的使用方法简单，但并不属于 POSIX 标准，在各类 UNIX 平台上的实现不尽相同，因此其用途受

到了一定的限制。而 POSIX 标准定义的信号处理接口是 sigaction 函数，其接口头文件及原型如下：
 #include <signal.h>
 int sigaction(int signum, const struct sigaction *act, struct sigaction *oldact);

 ◆ signum：要操作的信号。
 ◆ act：要设置的对信号的新处理方式。
 ◆ oldact：原来对信号的处理方式。
 ◆ 返回值：0 表示成功，-1 表示有错误发生。

 struct sigaction 类型用来描述对信号的处理，定义如下：
 struct sigaction
 {
  void     (*sa_handler)(int);
  void     (*sa_sigaction)(int, siginfo_t *, void *);
  sigset_t  sa_mask;
  int       sa_flags;
  void     (*sa_restorer)(void);
 };

 在这个结构体中，成员 sa_handler 是一个函数指针，其含义与 signal 函数中的信号处理函数类似。成员

sa_sigaction 则是另一个信号处理函数，它有三个参数，可以获得关于信号的更详细的信息。当 sa_flags 成员的值

包含了 SA_SIGINFO 标志时，系统将使用 sa_sigaction 函数作为信号处理函数，否则使用 sa_handler 作为信号处理

函数。在某些系统中，成员 sa_handler 与 sa_sigaction 被放在联合体中，因此使用时不要同时设置。
 sa_mask 成员用来指定在信号处理函数执行期间需要被屏蔽的信号，特别是当某个信号被处理时，它自身会被

自动放入进程的信号掩码，因此在信号处理函数执行期间这个信号不会再度发生。
 sa_flags 成员用于指定信号处理的行为，它可以是一下值的“按位或”组合。

 ◆ SA_RESTART：使被信号打断的系统调用自动重新发起。
 ◆ SA_NOCLDSTOP：使父进程在它的子进程暂停或继续运行时不会收到 SIGCHLD 信号。
 ◆ SA_NOCLDWAIT：使父进程在它的子进程退出时不会收到 SIGCHLD 信号，这时子进程如果退出也不会成为僵

尸进程。
 ◆ SA_NODEFER：使对信号的屏蔽无效，即在信号处理函数执行期间仍能发出这个信号。
 ◆ SA_RESETHAND：信号处理之后重新设置为默认的处理方式。
 ◆ SA_SIGINFO：使用 sa_sigaction 成员而不是 sa_handler 作为信号处理函数。

 re_restorer 成员则是一个已经废弃的数据域，不要使用。


 在这个例程中使用 sigaction 函数为 SIGUSR1 和 SIGUSR2 信号注册了处理函数，然后从标准输入读入字符

 。程序运行后首先输出自己的 PID，如：
  My PID is 5904

  这时如果从另外一个终端向进程发送 SIGUSR1 或 SIGUSR2 信号，用类似如下的命令：
  kill -USR1 5904

  则程序将继续输出如下内容：
  SIGUSR1 received
  read is interrupted by signal

  这说明用 sigaction 注册信号处理函数时，不会自动重新发起被信号打断的系统调用。如果需要自动重新发

 起，则要设置 SA_RESTART 标志，比如在上述例程中可以进行类似一下的设置：
  sa_usr.sa_flags = SA_RESTART;
*/
static void sig_usr(int signum)
{
    if(signum == SIGUSR1)
    {
        printf("SIGUSR1 received\n");
    }
    else if(signum == SIGUSR2)
    {
        printf("SIGUSR2 received\n");
    }
    else
    {
        printf("signal %d received\n", signum);
    }
}

int main(void)
{
    char buf[512];
    int  n;
    struct sigaction sa_usr;
    sa_usr.sa_flags = 0; // sigaction再处理信号时，不会自动重新发起被信号打断的系统调用。如果需要自动重新发起，则要设置 SA_RESTART 标志
    sa_usr.sa_handler = sig_usr;   //信号处理函数

    sigaction(SIGUSR1, &sa_usr, NULL);
    sigaction(SIGUSR2, &sa_usr, NULL);

    printf("My PID is %d\n", getpid());

    while(1)
    {
        if((n = read(STDIN_FILENO, buf, 511)) == -1)
        {
            if(errno == EINTR)
            {
                printf("read is interrupted by signal\n");
            }
        }
        else
        {
            buf[n] = '\0';
            printf("%d bytes read: %s\n", n, buf);
        }
    }

    return 0;
}

