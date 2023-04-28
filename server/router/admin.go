package router

import (
	"app/api/controller/admin"
	"app/api/middleware"
	"github.com/gin-gonic/gin"
)

// 后台路由
func adminRouter(r *gin.Engine) *gin.Engine {
	//路由组
	adminGroup := r.Group("/admin")
	adminGroup.POST("/user/login", admin.AdminUserLogin)
	adminGroup.Use(middleware.AdminUserAuth)
	{
		//用户管理
		userGroup := adminGroup.Group("/user")
		{
			userGroup.GET("/info", admin.AdminUserInfo)
			userGroup.GET("/menu", admin.AdminUserMenu)
			userGroup.GET("/list", admin.AdminUserList)
			userGroup.POST("/create", admin.AdminUserCreate)
			userGroup.PUT("/update", admin.AdminUserUpdate)
			userGroup.DELETE("/delete", admin.AdminUserDelete)
			userGroup.PUT("/roleUpdate", admin.AdminUserRoleUpdate)
			userGroup.GET("/role", admin.AdminUserRole)
			userGroup.PUT("/passwordUpdate", admin.AdminUserPasswordUpdate)
			userGroup.PUT("/info/update", admin.AdminUserInfoUpdate)
		}
		//角色管理
		roleGroup := adminGroup.Group("/role")
		{
			roleGroup.GET("/list", admin.AdminRoleList)
			roleGroup.POST("/create", admin.AdminRoleCreate)
			roleGroup.PUT("/update", admin.AdminRoleUpdate)
			roleGroup.DELETE("/delete", admin.AdminRoleDelete)
			roleGroup.GET("/auth", admin.AdminRoleAuth)
			roleGroup.PUT("/authUpdate", admin.AdminRoleAuthUpdate)
		}
		//api管理
		apiGroup := adminGroup.Group("/api")
		{
			apiGroup.GET("/list", admin.AdminApiList)
			apiGroup.POST("/create", admin.AdminApiCreate)
			apiGroup.PUT("/update", admin.AdminApiUpdate)
			apiGroup.DELETE("/delete", admin.AdminApiDelete)
			apiGroup.GET("/group", admin.AdminApiGroup)
		}
		//menu管理
		menuGroup := adminGroup.Group("/menu")
		{
			menuGroup.GET("/list", admin.AdminMenuList)
			menuGroup.POST("/create", admin.AdminMenuCreate)
			menuGroup.PUT("/update", admin.AdminMenuUpdate)
			menuGroup.DELETE("/delete", admin.AdminMenuDelete)
			menuGroup.PUT("/sort", admin.AdminMenuSort)
		}
		//config管理
		configGroup := adminGroup.Group("/config")
		{
			configGroup.GET("/list", admin.AdminConfigList)
			configGroup.PUT("/update", admin.AdminConfigUpdate)
		}
		//ssl证书
		adminGroup.GET("/ssl", admin.Ssl)
		//密码管理
		passwordGroup := adminGroup.Group("/password")
		{
			passwordGroup.GET("/list", admin.PasswordList)
			passwordGroup.POST("/create", admin.PasswordCreate)
			passwordGroup.PUT("/update", admin.PasswordUpdate)
			passwordGroup.DELETE("/delete", admin.PasswordDelete)
		}
		//youtube管理
		youtubeGroup := adminGroup.Group("/youtube")
		{
			youtubeGroup.POST("/create", admin.YoutubeCreate)
			youtubeGroup.GET("/list", admin.YoutubeList)
			youtubeGroup.DELETE("/delete", admin.YoutubeDelete)
			youtubeGroup.PUT("/retry", admin.YoutubeRetry)
		}
	}
	return r
}
