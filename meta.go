package multicache_gin

import "time"

// 配置
var metricOptions MetricsOption

// QueryRange 时间范围
type QueryRange struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

// Label 标签
type Label struct {
	Key string
	Val string
}

// RateOptions 比率查询参数，一般用于统计缓存命中率
type RateOptions struct {
	QueryRange
	Series      string
	Job         string
	Numerator   []Label
	Denominator []Label
}
