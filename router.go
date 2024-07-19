package multicache_gin

import (
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter(group *gin.RouterGroup, fns ...MetricsOptionHandler) error {
	// 默认配置
	metricOptions = DefaultMetricsOption()
	for _, fn := range fns {
		fn(&metricOptions)
	}

	// 注册路由-按照配置的解决方案名称获取所有序列名称
	group.POST("/seriesname", func(c *gin.Context) {

		c.JSON(200, gin.H{})
	})

	// 缓存命中率 指定series, job , adaptor

	// 热key排行榜

	// 指定key查询次数

	return nil
}
