package multicache_gin

import (
	"fmt"
	"testing"
	"time"
)

const qHost = "http://localhost:8086"
const sName = "test"

func TestSeriesNames(t *testing.T) {

	// 基础配置
	metricOptions = DefaultMetricsOption()
	WithQueryHost(qHost)(&metricOptions)
	WithSolutionName(sName)(&metricOptions)

	ts := QueryRange{
		StartTime: time.Now().Add(-time.Hour * 24),
		EndTime:   time.Now(),
	}

	names, err := SeriesNames(ts)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(names)

}

func TestSeriesLabelSet(t *testing.T) {

	// 基础配置
	metricOptions = DefaultMetricsOption()
	WithQueryHost(qHost)(&metricOptions)
	WithSolutionName(sName)(&metricOptions)

	ts := QueryRange{
		StartTime: time.Now().Add(-time.Hour * 24),
		EndTime:   time.Now(),
	}
	names, err := SeriesLabelSet(ts, "jiaoyan_multicache_event")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(names)
}
