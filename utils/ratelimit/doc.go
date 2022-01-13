package ratelimit

/*
限流算法常见的包括 leakyBucket 和 tokenBucket,以下几个实现用的比较多:
leakyBucket: go.uber.org/ratelimit
tokenBucket: golang.org/x/time/rate
当前比较推荐的是 自适应限流算法。
参考：https://www.jianshu.com/p/60fa376b9849
源码：https://github.com/go-kratos/aegis/tree/main/ratelimit/bbr
*/
