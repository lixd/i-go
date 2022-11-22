构建镜像
$ docker build -f file -t name:tag context_path
> docker build . -f API.Dockerfile -t test:0.0.1

删除所有已停止的容器
docker rm $(docker ps -aq -f status=exited)

删除所有虚悬镜像
$ docker rmi $(docker images -q -f dangling=true)

删除所有镜像

$ docker rmi $(docker images -q)

.dockerignore 忽略构建镜像时传输到 server 端的文件
