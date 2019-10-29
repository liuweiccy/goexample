package uberguide

import "time"

// 针对超过3个可选参数的情况下，使用参数选项设计模式处理
type options struct {
	timeout time.Duration
	caching bool
}

type Option interface {
	apply(*options)
}

type optionFunc func(o *options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithTimeout(t time.Duration) Option {
	return optionFunc(func(o *options) {
		o.timeout = t
	})
}

func WithCaching(cache bool) Option {
	return optionFunc(func(o *options) {
		o.caching = cache
	})
}

func Connect(addr string, opts ...Option) {
	opts1 := &options{
		timeout: 0,
		caching: false,
	}

	for _, o := range opts {
		o.apply(opts1)
	}
}

func use()  {
	addr := ""
	Connect(addr)
	Connect(addr, WithCaching(true))
	Connect(addr, WithTimeout(time.Second))
	Connect(addr, WithTimeout(time.Microsecond), WithCaching(false))
}
