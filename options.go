package multicache_gin

type MetricsOptionHandler func(opt *MetricsOption)

// MetricsOption 用于配置Metrics
type MetricsOption struct {
	SolutionName string
	QueryHost    string
}

// DefaultMetricsOption 默认配置
func DefaultMetricsOption() MetricsOption {
	return MetricsOption{
		SolutionName: "default",
		QueryHost:    "http://localhost:8086",
	}
}

// WithSolutionName 设置解决方案名称
func WithSolutionName(name string) MetricsOptionHandler {
	return func(opt *MetricsOption) {
		opt.SolutionName = name
	}
}

// WithQueryHost 设置查询API地址
func WithQueryHost(host string) MetricsOptionHandler {
	return func(opt *MetricsOption) {
		opt.QueryHost = host
	}
}
