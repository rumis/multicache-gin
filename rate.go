package multicache_gin

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"github.com/rumis/multicache/metrics/prometheus"
)

// TopKey 热key排行榜
// topk(%d, sum by(key) (rate(%s{job="%s", event="Hit"}[1m])))
func TopKey(ctx context.Context, ts QueryRange, series string, job string, topk int) error {
	qStr := fmt.Sprintf(`topk(%d, sum by(key) (rate(%s{job="%s", event="Hit"}[1m])))`, topk, series, job)
	query := prometheus.NewQuery(prometheus.WithPromHttpApiQueryHost(metricOptions.QueryHost))
	vals, err := query.Range(qStr, v1.Range{
		Start: ts.StartTime,
		End:   ts.EndTime,
		Step:  time.Second * 60,
	})
	if err != nil {
		return errors.WithMessage(err, "获取序列信息失败")
	}

	switch vals.Type() {
	case model.ValMatrix:
		// 解析matrix，按照时间
		matrix := vals.(model.Matrix)
		for _, v := range matrix {
			k := v.Metric["key"]
			for _, s := range v.Values {
				fmt.Println(k, s.Timestamp.Time().Format("2006-01-02 15:04:05"), s.Value)
			}
		}
	default:
		fmt.Println(vals)
	}
	// fmt.Println(vals)
	return nil
}
