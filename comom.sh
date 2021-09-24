# 统计进程打开的FD
lsof -n| awk '{print $3}'|sort|uniq -c|sort -nr|head -10
# 统计用户打开的FD
lsof -n| awk '{print $2}'|sort|uniq -c|sort -nr|head -10
# 统计命令打开的FD
lsof -n| awk '{print $1}'|sort|uniq -c|sort -nr|head -10

# aliyun cdn log analyzer 被访问次数最多的url及访问用户ip
awk '{print $3$8}' {filename} |\
awk '{ips[$1]++;next} END {for (ip in ips) print ips[ip] " " ip}'|\
sort -rn|\
head -n 10


# 查看TCP连接数及其状态
netstat -n | awk '/^tcp/ {++S[$NF]} END {for(a in S) print a, S[a]}'

# 查看句柄占用情况 模糊值
lsof -n|awk '{print $2}'|sort|uniq -c|sort -nr|more

# 查看句柄占用情况 精确值
ls /proc/$pid/fd/ |wc -l

