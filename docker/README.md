构建镜像
$ docker build -f file -t name:tag context_path
删除所有虚悬镜像
$ docker rmi $(docker images -q -f dangling=true)
删除所有已停止的容器
docker rm $(docker ps -aq -f status=exited)

.dockerignore 忽略构建镜像时传输到 server 端的文件
