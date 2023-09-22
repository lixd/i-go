# 使用nsenter调试
# 入参： podName NS
function e() {
    set -eu
    ns=${2-"default"}
    pod=`kubectl -n $ns describe pod $1 | grep -A10 "^Containers:" | grep -Eo 'docker://.*$' | head -n 1 | sed 's/docker:\/\/\(.*\)$/\1/'`
    pid=`docker inspect -f {{.State.Pid}} $pod`
    echo "entering pod netns for $ns/$1"
    cmd="nsenter -n --target $pid"
    echo $cmd
    $cmd
}

# 根据pod名称和命名空间查找pid
# pod2pid
# 入参： podName NS
function e(){
    set -eu
    ns=${2-"default"}
    podID=`kubectl get   pods -n  $ns  $1  -o json |jq  .status.containerStatuses[0].containerID | awk -F '/' '{print $NF}'| awk -F '"' '{print $1}'`
    pid=`crictl  inspect 7b8c272aeb003f09b1c0a9dec29968fc713e520d928362f2266937a790e2680a | jq .info.pid`
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


# 用法
# 1.根据 pod namespace + name 找到 podID
# 2.根据 podID 找到 pid
# 3.nsenter -n --target $pid 进入到对应 pid 的网络命名空间进行测试
function foo() {
  namespace=caas-system
  podName=caas-dns-crd-controller-manager-598fc4bddc-m77sq
  podID=`kubectl get  pods -n $namespace $podName  -o json |jq  .status.containerStatuses[0].containerID | awk -F '/' '{print $NF}'| awk -F '"' '{print $1}'`
  echo "podID:" $podID
  pid=`crictl  inspect $podID | jq .info.pid`
  echo "pid:" $pid
  echo "entering pod netns for $ns/$1"
  cmd="nsenter -n --target $pid"
#  cmd="nsenter --mount --uts --ipc --net --pid --target $pid"
  echo $cmd
  $cmd
#  ping dailybuild.cpcsng.top
}


kubectl describe secrets $(kubectl get secrets -n kube-system |grep admin |cut -f1 -d ' ') -n kube-system |grep -E '^token' |cut -f2 -d':'|tr -d '\t'|tr -d ' '

eyJhbGciOiJSUzI1NiIsImtpZCI6IiJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJhZG1pbi10b2tlbi0ycTI4ZiIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50Lm5hbWUiOiJhZG1pbiIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6IjkzMzE2ZmZhLTc1NDUtMTFlOS1iNjE3LTAwMTYzZTA2OTkyZCIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDprdWJlLXN5c3RlbTphZG1pbiJ9.EQzj2LsWn2k31m-ksn9GmB1bZTi1Xjw1fnmWFgRKlwhS2QAaVnDXfV_TgUovpq5oWKh7h0gTVaNaK4KKK76yAv6GfMehpOdIO5xHCfQAWVRhla1cwUDC64tz7vJ1zGcx_lz4hKfhdXN1T8FYS0B0hf3h2OloAMfCZTzDjRWz24GVwH-WRTEwY_5tav65GiZzBTsnz1vV7NOcx-Kl8AK2HbowtBYqK05x7oOmp84FiQMwpYU-7g0c03h61zev4lvf0e-HFtqKiByPi8gD-uiVRvE-xayOz5oIESWw2GfhzfNf_uyR7eLplCKUBecVMtwVsBauNaeqU-IIJW5VIHAOxw
TOKEN=$(kubectl describe secrets $(kubectl get secrets -n kube-system |grep admin |cut -f1 -d ' ') -n kube-system |grep -E '^token' |cut -f2 -d':'|tr -d '\t'|tr -d ' ')

# kubectl config view |grep server|cut -f 2- -d ":" | tr -d " "

APISERVER=https://192.168.10.229:6443
TOKEN=$(cat k8s.token)
curl -k -H "Authorization: Bearer $TOKEN" $APISERVER/api

