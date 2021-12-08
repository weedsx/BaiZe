package systemRouter

import (
	"baize/app/common/middlewares"
	"baize/app/system/controller/configController"
	"github.com/gin-gonic/gin"
)

func InitSysConfigRouter(router *gin.RouterGroup) {
	systemConfig := router.Group("/system/config")
	systemConfig.GET("/list", middlewares.HasPermission("system:config:list"), configController.ConfigList)
	systemConfig.GET("/export", middlewares.HasPermission("system:config:export"), configController.ConfigExport)
	systemConfig.GET("/:configId", middlewares.HasPermission("system:config:query"), configController.ConfigGetInfo)
	systemConfig.GET("/configKey/:configKey", middlewares.HasPermission("system:config:query"), configController.ConfigGetConfigKey)
	systemConfig.POST("", middlewares.HasPermission("system:config:add"), configController.ConfigAdd)
	systemConfig.PUT("", middlewares.HasPermission("system:config:edit"), configController.ConfigEdit)
	systemConfig.DELETE("/:configIds", middlewares.HasPermission("system:config:remove"), configController.ConfigRemove)
	systemConfig.POST("/clearCache", middlewares.HasPermission("system:config:remove"), configController.ConfigClearCache)

}
