# find


```bash
# 找出 7 天前的文件并删除
find ./ -maxdepth 1 -mtime +7 | xargs rm -rf
```
