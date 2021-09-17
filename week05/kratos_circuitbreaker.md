
#### 前言
之前了解了自适应限流，它属于过载保护：统计cpu 的使用率，当cpu 到达设定的上限值后，再利用利特尔法则实现自适应限流，这只是一种自我保护措施，如果是调用其他服务，也需要一种保护机制，避免把下游服务打死，这就需要一种熔断机制。

#### kratos 熔断器：
先大概说下kratos 熔断器的原理：就是发现错误率达到一定的程度是，自动拒绝请求，但是并非完全熔断，而是有概率丢弃一些请求，随着失败率的提高，拒绝的概率就越高，慢慢的，随着成功率的提升，就拒绝请求的概率就越小。这就到达另一种自适应限流的效果。

源码路径：https://github.com/go-kratos/aegis/blob/main/circuitbreaker/sre/sre.go

为了更容易阅读源码，先去看下 google sre 算法的原理：[Handling Overload – Google SRE books](https://sre.google/sre-book/handling-overload/)

go-zero 貌似也是用google sre 算法实现熔断器。

##### google sre 核心的算法：
rejectProba = max(0,(requests−K∗accepts)/(requests+1))

从上面的公式可以看到，
1. 拒绝请求的概率rejectProba 的取值范围是[0,1.0), requests 大于K∗accepts 后，才可能拒绝请求或者说限流。
2. K应该是一个大于1 的值，不然在所有请求都成功的情况下，也可能拒绝请求
3. K 值越大，因请求失败而限流的敏感度就越低，如果k=1,只要有一个失败，requests - K∗accepts 就会大于0，就可能拒绝请求，k=2时，失败的次数大于成功的次数后，才可能决绝请求。可以看出K是用来调整熔断器敏感度的。
3. 当拒绝请求，服务得到恢复后，成功的次数accepts 就会增加，虽然 requests 也增加，但是由于k的存在，k*accepts 增加得更快，requests−K∗accepts 就会变成一个负数，rejectProba 变成0，就不会再拒绝请求。

##### 总结下，K只越小，拒绝请求的行为就越激进，K只越大，拒绝请求的行为就越温和或迟钝，


##### 从上面Google SRE 算法大概理解，理解kratos 熔断器就容易很多了：
```go
// Breaker is a sre CircuitBreaker pattern.
type Breaker struct {
	stat window.RollingCounter  // 滑动窗口统计一段时间内请求数量和成功请求数量
	r    *rand.Rand
	// rand.New(...) returns a non thread safe object
	randLock sync.Mutex

	// Reducing the k will make adaptive throttling behave more aggressively,
	// Increasing the k will make adaptive throttling behave less aggressively.
	k       float64 // 就公式里的K，大于1，
	request int64 //请求数量，即请求达到一定数量后，才开始判断熔断器是否打开，避免一开始有失败的请求，就打开熔断器

	state int32 //熔断器的状态，StateOpen 打开， 或者StateClosed 关闭
}

// NewBreaker return a sreBresker with options
func NewBreaker(opts ...Option) circuitbreaker.CircuitBreaker {
	opt := options{
		success: 0.6, // 成功率低于 0.6, 就进入熔断逻辑，就会有一定概率拒绝请求。
		request: 100, // 总的请求数量至少100个后，才去判断是否开启熔断，即请求数量达到100个后，算法才启动。
		bucket:  10,
		window:  3 * time.Second, //滑动窗口
	}
	for _, o := range opts {
		o(&opt)
	}
	counterOpts := window.RollingCounterOpts{
		Size:           opt.bucket,
		BucketDuration: time.Duration(int64(opt.window) / int64(opt.bucket)),
	}
	stat := window.NewRollingCounter(counterOpts)
	return &Breaker{
		stat:    stat,
		r:       rand.New(rand.NewSource(time.Now().UnixNano())),
		request: opt.request,
		k:       1 / opt.success, //根据成功率可以计算出K值
		state:   StateClosed, //初始的熔断器的状态是关闭状态
	}
}
```


业务层是通过调用Breaker  Allow() 方法来决定是否拒绝请求。
```go
// Allow request if error returns nil.
func (b *Breaker) Allow() error {
	// The number of requests accepted by the backend
	accepts, total := b.summary() // 获取成功的次数，和 总的请求数
	// The number of requests attempted by the application layer(at the client, on top of the adaptive throttling system)
	requests := b.k * float64(accepts) // 算法里 k*accepts
	// check overflow requests = K * accepts
	//如果总请求数低于100，直接放过； 如果 total <  k*accepts , 也直接放过
	if total < b.request || float64(total) < requests {
		if atomic.LoadInt32(&b.state) == StateOpen { //判断熔断器是否是打开的，如果是，就关闭熔断器
			atomic.CompareAndSwapInt32(&b.state, StateOpen, StateClosed)
		}
		return nil
	}
	// 到这里就说明，total > k*accepts , 即失败次数过多了，成功的概率已经低于0.6 了，就进入熔断机制。
	if atomic.LoadInt32(&b.state) == StateClosed {
		atomic.CompareAndSwapInt32(&b.state, StateClosed, StateOpen) // 打开熔断器
	}
	dr := math.Max(0, (float64(total)-requests)/float64(total+1)) // google sre 算法，计算出拒绝请求的概率
	drop := b.trueOnProba(dr) // 根据概率判断此次请求是否放过，
	if drop {
		return circuitbreaker.ErrNotAllowed //如果拒绝此次请求，就返回错误，上层就丢弃
	}
	return nil //根据概率, 接受此次请求，返回ni
}

func (b *Breaker) trueOnProba(proba float64) (truth bool) {
	b.randLock.Lock()
	truth = b.r.Float64() < proba //判断是否低于拒绝请求的概率proba，
	b.randLock.Unlock()
	return
}

```
Allow() 方法 主要实现SRE 算法。

##### 业务层调用Allow() 后，放过请求，
1. 如果请求的最后结果是成功的，需要更新计数MarkSuccess()
```go
// MarkSuccess mark requeest is success.
func (b *Breaker) MarkSuccess() {
	b.stat.Add(1)
}
```
2. 请求失败或者本地拒绝请求，也需要计数, MarkFailed()：
```go
// MarkFailed mark request is failed.
func (b *Breaker) MarkFailed() {
	// NOTE: when client reject requets locally, continue add counter let the
	// drop ratio higher.
	b.stat.Add(0)
}
```
因为 google sre 算法是是基于失败的严重情况来实现自适应的丢弃请求。

##### 示例：
在kratos 中间件展示了熔断器的用法：
https://github.com/go-kratos/kratos/blob/main/middleware/circuitbreaker/circuitbreaker.go:
```go
func Client(opts ...Option) middleware.Middleware {
	opt := &options{
		group: group.NewGroup(func() interface{} {
			return sre.NewBreaker()
		}),
	}
	for _, o := range opts {
		o(opt)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			info, _ := transport.FromServerContext(ctx)
			breaker := opt.group.Get(info.Operation()).(circuitbreaker.CircuitBreaker)
			if err := breaker.Allow(); err != nil { // 使用熔断器判断是否放通此次请求
				// rejected
				// NOTE: when client reject requets locally,
				// continue add counter let the drop ratio higher.
				breaker.MarkFailed()  //本地拒绝请求后，也计数
				return nil, ErrNotAllowed
			}
			// allowed
			reply, err := handler(ctx, req)
			if err != nil && (errors.IsInternalServer(err) || errors.IsServiceUnavailable(err) || errors.IsGatewayTimeout(err)) {
				breaker.MarkFailed() //如果此次请求失败，计数
			} else {
				breaker.MarkSuccess() //如果此次请求成功，计数
			}
			return reply, err
		}
	}
}
```
