#!/bin/bash
# 移除已经停止的容器
docker system prune -f
# 删除当前没有使用(即没有启动为container的)的镜像
docker image prune -a
