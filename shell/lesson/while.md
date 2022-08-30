# while 循环
[shell 脚本 while 循环](https://www.cnblogs.com/bandaoyu/p/16510252.html#while%E5%BE%AA%E7%8E%AF)
语法：while 条件；do...；done

利用while循环计算1到100的和
```shell
#!/bin/bash
i=1
sum=0
while [ $i -le 100 ]
do
  let sum=sum+$i
  let i++
done

echo $sum
```

使用read结合while循环读取文本文件
```shell
#!/bin/bash
file=$1                  #将位置参数1的文件名复制给file
if [ $# -lt 1 ];then      #判断用户是否输入了位置参数
  echo "Usage:$0 filepath"
  exit
fi
while read -r line   #从file文件中读取文件内容赋值给line（使用参数r会屏蔽文本中的特殊符号，只做输出不做转译）
do

  echo $line        #输出文件内容

done   <  $file
```

按列读取文件内容
```shell
#!/bin/bash
file=$1
if [[ $# -lt 1 ]]
then
  echo "Usage: $0 please enter you filepath"
  exit
fi
while read -r  f1 f2 f3    #将文件内容分为三列
do
  echo "file 1:$f1 ===> file 2:$f2 ===> file 3:$f3"   #按列输出文件内容

done < "$file"
```
