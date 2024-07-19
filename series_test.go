package multicache_gin

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var qHost = os.Getenv("PROM_QUERY_HOST")

const sName = "jiaoyan_multicache"

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

func TestSeriesJobs(t *testing.T) {

	// 基础配置
	metricOptions = DefaultMetricsOption()
	WithQueryHost(qHost)(&metricOptions)
	WithSolutionName(sName)(&metricOptions)

	ts := QueryRange{
		StartTime: time.Now().Add(-time.Hour * 24),
		EndTime:   time.Now(),
	}
	names, err := SeriesJobs(ts, "jiaoyan_multicache_event")
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
	names, err := SeriesLabelSet(ts, "jiaoyan_multicache_event", "courseware")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(names)
}
