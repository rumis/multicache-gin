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

	// 注册路由
	group.POST("/seriesname", func(c *gin.Context) {

		c.JSON(200, gin.H{})
	})

	return nil
}
