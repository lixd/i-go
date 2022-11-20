#!/usr/bin/env bash
# 移除在 netfilter 中的日志规则
# 需要手动把 dport 修改成需要记录的端口
dport=8080

iptables -t raw -D PREROUTING -p tcp --dport $dport -j LOG --log-prefix "target in raw.prerouting>"
iptables -t mangle -D PREROUTING -p tcp --dport $dport -j LOG --log-prefix "target in mangle.prerouting>"
iptables -t nat -D PREROUTING -p tcp --dport $dport -j LOG --log-prefix "target in nat.prerouting>"
iptables -t mangle -D INPUT -p tcp --dport $dport -j LOG --log-prefix "target in mangle.input>"
iptables -t filter -D INPUT -p tcp --dport $dport -j LOG --log-prefix "target in filter.input>"
iptables -t raw -D OUTPUT -p tcp --dport $dport -j LOG --log-prefix "target in raw.output>"
iptables -t mangle -D OUTPUT -p tcp --dport $dport -j LOG --log-prefix "target in mangle.output>"
iptables -t nat -D OUTPUT -p tcp --dport $dport -j LOG --log-prefix "target in nat.output>"
iptables -t filter -D OUTPUT -p tcp --dport $dport -j LOG --log-prefix "target in filter.output>"
iptables -t mangle -D FORWARD -p tcp --dport $dport -j LOG --log-prefix "target in mangle.forward>"
iptables -t filter -D FORWARD -p tcp --dport $dport -j LOG --log-prefix "target in filter.forward>"
iptables -t mangle -D POSTROUTING -p tcp --dport $dport -j LOG --log-prefix "target in mangle.postrouting>"
iptables -t nat -D POSTROUTING -p tcp --dport $dport -j LOG --log-prefix "target in nat.postrouting>"

