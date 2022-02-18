package ratelimit

/*
限流算法常见的包括 leakyBucket 和 tokenBucket,以下几个实现用的比较多:
leakyBucket: go.uber.org/ratelimit
tokenBucket: golang.org/x/time/rate
当前比较推荐的是 自适应限流算法。
参考：https://www.jianshu.com/p/60fa376b9849
源码：https://github.com/go-kratos/aegis/tree/main/ratelimit/bbr
kratos文档：https://github.com/go-kratos/kratos/blob/v1.0.x/docs/ratelimit.md
Sentinel文档：https://github.com/alibaba/Sentinel/wiki/%E7%B3%BB%E7%BB%9F%E8%87%AA%E9%80%82%E5%BA%94%E9%99%90%E6%B5%81
*/
