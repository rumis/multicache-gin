package multicache_gin

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/rumis/liutils/set"
	"github.com/rumis/multicache/metrics/prometheus"
)

// SeriesNames 获取所有序列名称-基于初始化路由时提供的解决方案名称
func SeriesNames(ts QueryRange) ([]string, error) {
	query := prometheus.NewQuery(prometheus.WithPromHttpApiQueryHost(metricOptions.QueryHost))
	labelSet, err := query.Series(ts.StartTime, ts.EndTime, `{__name__=~"^`+metricOptions.SolutionName+`.+"}`)
	if err != nil {
		return nil, errors.WithMessage(err, "获取序列名称列表失败")
	}
	// 去重
	names := set.NewSet[string]()
	for _, v := range labelSet {
		name, ok := v["__name__"]
		if !ok {
			continue
		}
		if strings.HasSuffix(string(name), "_bucket") || strings.HasSuffix(string(name), "_sum") || strings.HasSuffix(string(name), "_count") {
			// 直方图的几个序列
			continue
		}
		names.Add(string(name))
	}
	return names.ToSlide(), nil
}

// SeriesJobs 获取指定序列的job名称
func SeriesJobs(ts QueryRange, series string) ([]string, error) {
	query := prometheus.NewQuery(prometheus.WithPromHttpApiQueryHost(metricOptions.QueryHost))
	labelSet, err := query.Series(ts.StartTime, ts.EndTime, series)
	if err != nil {
		return nil, errors.WithMessage(err, "获取序列信息失败")
	}
	jobs := set.NewSet[string]()
	for _, labels := range labelSet {
		for k, v := range labels {
			if k == "job" {
				jobs.Add(string(v))
			}
		}
	}
	return jobs.ToSlide(), nil
}

// SeriesLabelSet 获取指定序列的标签和标签值
func SeriesLabelSet(ts QueryRange, series string, job string) (map[string][]string, error) {
	query := prometheus.NewQuery(prometheus.WithPromHttpApiQueryHost(metricOptions.QueryHost))
	set, err := query.Series(ts.StartTime, ts.EndTime, series)
	if err != nil {
		return nil, errors.WithMessage(err, "获取序列信息失败")
	}
	labelSet := make(map[string]map[string]struct{})
	for _, labels := range set {
		labelJob, ok := labels["job"]
		if !ok || string(labelJob) != job {
			continue
		}
		for k, v := range labels {
			if k == "__name__" || k == "job" {
				continue
			}
			col, ok := labelSet[string(k)]
			if !ok {
				labelSet[string(k)] = map[string]struct{}{string(v): {}}
				continue
			}
			col[string(v)] = struct{}{}
			labelSet[string(k)] = col
		}
	}
	labelCol := make(map[string][]string)
	for k, v := range labelSet {
		col := make([]string, 0, len(v))
		for k := range v {
			col = append(col, k)
		}
		labelCol[k] = col
	}
	return labelCol, nil
}
