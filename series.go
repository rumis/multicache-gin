package multicache_gin

import (
	"github.com/pkg/errors"
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
	nameMap := make(map[string]struct{})
	for _, v := range labelSet {
		name, ok := v["__name__"]
		if !ok {
			continue
		}
		nameMap[string(name)] = struct{}{}
	}
	// 转数组
	names := make([]string, 0, len(nameMap))
	for v := range nameMap {
		names = append(names, v)
	}
	return names, nil
}

// SeriesLabelSet 获取指定序列的标签和标签值
func SeriesLabelSet(ts QueryRange, series string) (map[string][]string, error) {

	query := prometheus.NewQuery(prometheus.WithPromHttpApiQueryHost(metricOptions.QueryHost))

	set, err := query.Series(ts.StartTime, ts.EndTime, series)
	if err != nil {
		return nil, errors.WithMessage(err, "获取序列信息失败")
	}
	labelSet := make(map[string]map[string]struct{})
	for _, labels := range set {
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
