构建镜像
$ docker build -f file -t name:tag context_path
> docker build . -f API.Dockerfile -t test:0.0.1

查看 docker 的磁盘占用情况
docker system df
一键清理
docker system prune -a
删除所有已停止的容器
docker rm $(docker ps -aq -f status=exited)

删除所有虚悬镜像
$ docker rmi $(docker images -q -f dangling=true)

删除所有镜像

$ docker rmi $(docker images -q)

.dockerignore 忽略构建镜像时传输到 server 端的文件

清理 k8s 中的 pod
namespace=test-pods
pods=$(kubectl -n $namespace get pod|grep Completed|awk '{print $1}')
echo $pods
for i in $pods;do kubectl -n $namespace delete pod $i;done


pods=$(kubectl -n $namespace get pod|grep Completed|awk '{print $1}')
echo $pods
for i in $pods;do kubectl -n $namespace delete pod $i;done
