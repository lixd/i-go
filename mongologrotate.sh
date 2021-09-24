#!/bin/sh

# rotate mongodb log, https://docs.mongodb.com/manual/tutorial/rotate-log-files/
# logrotate is not used, as mongodb will not be able to know the log file is changed.
# using copytruncate could avoid it. but there is a window between copy and truncate,
# some mongod logs may get lost.
# write this simple script to send SIGUSR1 to mongod directly.
# 定时清理 mongod.log
logdir=/usr/local/docker/mongo/log
logfile=$logdir/mongod.log

# the max number of log files
maxfiles=10

# max file size: 64MB
maxsize=67108864

fsize=$(stat -c%s $logfile)
if [ $fsize -gt $maxsize ]; then
  # ask mongod to rotate the log file
  pkill --signal SIGUSR1 mongod

  # check and remove the old log files
  fnum=$(ls -1U $logdir | wc -l)
  if [ $fnum -gt $maxfiles ]; then
    rmnum=`expr $fnum - $maxfiles`
    ls -1rt $logdir | head -n $rmnum | awk '{ system("rm -f /usr/local/docker/mongo/log/" $0) }'
  fi

fi
