#!/usr/bin/env bash

# qq bot 操作脚本，根据用户输入执行对应指令。

function op() {
  svc=$1
        while [ 1 ]
        do
                    echo "请输入您的操作:"
                    echo "    1) 查看状态"
                    echo "    2) 查看实时日志"
                    echo "    3) 停止"
                    echo "    4) 启动"
                    echo "    5) 重启"
                    echo "    6) 退出脚本"
                    read  op
                    case $op in
                            1)
                                  systemctl status $svc
                                  ;;
                            2)
                                  journalctl -xu $svc -f
                                  ;;
                            3)
                                  systemctl stop $svc
                                  ;;
                            4)
                                  systemctl start $svc
                                  ;;
                            5)
                                  systemctl restart $svc
                                  ;;
                            6)
                                  break 2 # 跳出 case + while 循环
                                  ;;
                    esac
        done
}


{
  echo "请输入您需要操作的服务:"
  echo "    1) gocqhttp"
  echo "    2) zbp"
  echo "    3) 退出脚本"
  read  svc

  case $svc in
          1)
              op "bot-gocqhttp"
              ;;
          2)
              op "bot-zbp"
              ;;
          3)
              exit
              ;;
  esac
}
