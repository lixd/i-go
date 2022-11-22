#!/usr/bin/env bash
# 在 netfilter 所有位置都开启日志记录，
# 需要手动把 dport 修改成需要记录的端口
# 开启后使用 iptables -L -n|grep 8080 查看是否添加成功，然后使用 dmesg -Tx 查看
# 使用后记得关闭，否则会有大量内核日志。

dport=8080

iptables -t raw -I PREROUTING -p tcp --dport $dport -j LOG --log-prefix "target in raw.prerouting>"
iptables -t mangle -I PREROUTING -p tcp --dport $dport -j LOG --log-prefix "target in mangle.prerouting>"
iptables -t nat -I PREROUTING -p tcp --dport $dport -j LOG --log-prefix "target in nat.prerouting>"
iptables -t mangle -I INPUT -p tcp --dport $dport -j LOG --log-prefix "target in mangle.input>"
iptables -t filter -I INPUT -p tcp --dport $dport -j LOG --log-prefix "target in filter.input>"
iptables -t raw -I OUTPUT -p tcp --dport $dport -j LOG --log-prefix "target in raw.output>"
iptables -t mangle -I OUTPUT -p tcp --dport $dport -j LOG --log-prefix "target in mangle.output>"
iptables -t nat -I OUTPUT -p tcp --dport $dport -j LOG --log-prefix "target in nat.output>"
iptables -t filter -I OUTPUT -p tcp --dport $dport -j LOG --log-prefix "target in filter.output>"
iptables -t mangle -I FORWARD -p tcp --dport $dport -j LOG --log-prefix "target in mangle.forward>"
iptables -t filter -I FORWARD -p tcp --dport $dport -j LOG --log-prefix "target in filter.forward>"
iptables -t mangle -I POSTROUTING -p tcp --dport $dport -j LOG --log-prefix "target in mangle.postrouting>"
iptables -t nat -I POSTROUTING -p tcp --dport $dport -j LOG --log-prefix "target in nat.postrouting>"



