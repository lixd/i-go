#!/bin/bash
# 手动进行内存回收
# Usage: sh free-memory.sh

print() {
  prefix=$1
  value=$2
  if [ $value -gt 1048576 ] ;then # 10485760 = 1024 * 1024 * 10
    # 使用 bc 执行除法保留小数
    value=$(printf "%.5f" `echo "scale=5;$value/1048576"|bc`)
    unit='GB'
  elif [ $value -gt 10240 ];then
    value=$(printf "%.5f" `echo "scale=5;$value/1024"|bc`)
    unit='MB'
  else
      unit='KB'
  fi
  echo $prefix $value $unit
}

freeMemory(){
  sync;echo 1 > /proc/sys/vm/drop_caches # 1 表示清除pagecache。
  sync;echo 2 > /proc/sys/vm/drop_caches # 2 表示清除回收slab分配器中的对象（包括目录项缓存和inode缓存）。slab分配器是内核中管理内存的一种机制，其中很多缓存数据实现都是用的pagecache。
  sync;echo 3 > /proc/sys/vm/drop_caches # 3 表示清除pagecache和slab分配器中的缓存对象。
}

{
before=$(free|grep Mem|awk '{print $4}')
print 释放前空闲内存: $before

freeMemory

after=$(free|grep Mem|awk '{print $4}')
print 释放后空闲内存: $after

deal=`expr $after - $before`
print 本次释放内存: $deal
}
