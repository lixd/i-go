package availability

type sreBreaker struct {
}

// Allow Google SRE 中推荐的断路器算法 max(0, (requests - K*accepts) / (requests + 1))
/*
requests：请求数
K：系数，激进程度，K 越大越激进。
accepts：请求成功数
*/
/*func (b *sreBreaker) Allow() error {
	success, total := b.stat.Value()
	k := b.k * float64(success)
	if total < b.request || float64(total) < k {
		return nil
	}
	dr := math.Max(0, (float64(total)-k)/float64(total+1))
	rr := b.r.Float64()
	if dr <= rr {
		return nil
	}
	return ecode.ServiceUnavailable
}
type BackoffConfig struct {

}*/
// backoff 退避算法
/*
客户端需要限制请求频次，retry backoff 做一定的请求退让。
可以通过接口级别的error_details，挂载到每个 API 返回的响应里。
*/
/*func (bc BackoffConfig) backoff(retries int) time.Duration {
	if retries == 0 {
		return bc.baseDelay
	}
	backoff, max := float64(bc.baseDelay), float64(bc.MaxDelay)
	for backoff < max && retries > 0 {
		backoff *= bc.factor
		retries--
	}
	if backoff > max {
		backoff = max
	}
	// Randomize backoff delays so that if a cluster of requests start at
	// the same time, they won't operate in lockstep.
	backoff *= 1 + bc.jitter*(rand.Float64()*2-1)
	if backoff < 0 {
		return 0
	}
	return time.Duration(backoff)
}
*/
