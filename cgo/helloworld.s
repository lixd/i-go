.section .data # 数据段
msg:
  .ascii "Hello, World\n"
.section .text # 代码段
.global _start # 全局开始函数，必须是 _start

_start:
  movl $13, %edx # 系统调用 write 的第 3 个参数为 13
  movl $msg, %ecx # 第 2 个参数是 Hello World\n
  movl $1, %ebx # 第 1 个参数是 1，标准输出的编号
  movl $4, %eax # sys_write 系统调用的编号是 4
  int $0x80

  movl $0, %ebx # exit 系统调用的第一个参数，即 exit(0)
  movl $1, %eax # sys_exit 编号为 1
  int $0x80 # 调用 exit，退出程序

# 从 Hello World 来看看 Go 的运行流程 https://shimo.im/docs/HQHVVYhdGxycdVwX
# 系统调用实现：将要调用的系统调用编号存到exa寄存器，然后触发0x80指令即可，对应参数则到相应寄存器
#  编译链接
#  as -o helloworld.o helloworld.s
#  ld -o helloworld helloworld.o
