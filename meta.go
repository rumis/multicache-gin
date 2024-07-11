package multicache_gin

import "time"

// 配置
var metricOptions MetricsOption

// QueryRange 时间范围
type QueryRange struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}
