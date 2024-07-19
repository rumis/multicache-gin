package multicache_gin

import (
	"context"
	"testing"
	"time"
)

func TestTopKey(t *testing.T) {
	// 基础配置
	metricOptions = DefaultMetricsOption()
	WithQueryHost(qHost)(&metricOptions)
	WithSolutionName(sName)(&metricOptions)

	ts := QueryRange{
		StartTime: time.Now().Add(-time.Hour * 24),
		EndTime:   time.Now(),
	}
	err := TopKey(context.Background(), ts, "multicache_cache_hitmiss", "multicache_test_solution", 5)
	if err != nil {
		t.Fatal(err)
	}
	// fmt.Println(names)

}
