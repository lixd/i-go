# 使用nsenter调试
# 入参： podName NS
function e() {
    set -eu
    ns=${2-"default"}
    pod=`kubectl -n $ns describe pod $1 | grep -A10 "^Containers:" | grep -Eo 'docker://.*$' | head -n 1 | sed 's/docker:\/\/\(.*\)$/\1/'`
    pid=`docker inspect -f {{.State.Pid}} $pod`
    echo "entering pod netns for $ns/$1"
    cmd="nsenter -n --target $pid"
    echo $cmd068
    $cmd
}

# 根据pod名称和命名空间查找pid
# pod2pid
# 入参： podName NS
function e(){
    set -eu
    ns=${2-"default"}
    podID=`kubectl get   pods -n  $ns  $1  -o json |jq  .status.containerStatuses[0].containerID | awk -F '/' '{print $NF}'| awk -F '"' '{print $1}'`
    pid=`crictl  inspect $podID | jq .info.pid`
    echo "entering pod netns for $ns/$1"
    cmd="nsenter -n --target $pid"
    echo $cmd
    $cmd
}

# 根据pid查找pod名称和命名空间
# 入参：pid
function pid2pod {
  local pid=$1
  if [ -f /proc/${pid}/cgroup ]; then
    local cid=$(cat /proc/${pid}/cgroup | grep ":memory:" | awk -F '/' '{print $NF}' | awk -F '-'  '{print $3}'| awk -F '.' '{print $1}')
    if [ "${cid}" != "" ]; then
      ctr -n k8s.io c info ${cid} 2>/dev/null | jq -r '.Labels["io.kubernetes.pod.namespace"]+" "+.Labels["io.kubernetes.pod.name"]' 2>/dev/null
    fi
  fi
}

# 解析 pod pid
podID=`kubectl get   pods -n kc-system  kc-server-0  -o json |jq  .status.containerStatuses[0].containerID | awk -F '/' '{print $NF}'| awk -F '"' '{print $1}'`
