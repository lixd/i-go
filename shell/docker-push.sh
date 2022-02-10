#!/bin/bash
# docker inspect --format 语法：  https://blog.csdn.net/youmatterhsp/article/details/105735439
# 用法：sh docker-push.sh [imageID]
# 会自动根据ID提取镜像名称和Tag，然后推送到阿里云镜像仓库
imageID=$1
repoTag=$(docker image inspect --format '{{index .RepoTags 0}}' $imageID)
echo 'imageID: '$imageID ' tag: '$repoTag

echo '推送到北京'
docker tag $imageID registry.cn-beijing.aliyuncs.com/wlinno/$repoTag
docker push registry.cn-beijing.aliyuncs.com/wlinno/$repoTag
echo '推送到香港'
docker tag $imageID registry.cn-hongkong.aliyuncs.com/wlinno/$repoTag
docker push registry.cn-hongkong.aliyuncs.com/wlinno/$repoTag
echo '推送到硅谷'
docker tag $imageID registry.us-west-1.aliyuncs.com/wlinno/$repoTag
docker push registry.us-west-1.aliyuncs.com/wlinno/$repoTag
