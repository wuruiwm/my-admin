package initialize

import (
	"app/api/middleware"
	"app/api/response"
	"app/global"
	"app/router"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Server 初始化gin的http服务
func Server() {
	//是否开启debug
	if global.Config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	//实例化gin
	r := gin.New()
	//Logger中间件
	r.Use(middleware.Logger)
	//Recovery中间件
	r.Use(middleware.Recover)
	//上传文件目录
	r.Static("/upload", "./resource/upload")
	//处理不存在的路由
	r.NoMethod(ErrRoute)
	r.NoRoute(ErrRoute)
	//路由设置
	r = router.Router(r)
	//实例化server
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", global.Config.ServerPort),
		Handler: r,
	}
	//启动服务
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("http server exit error:%s\n", err.Error())
		}
	}()
	//优雅退出
	GracefulExit(server)
}

func ErrRoute(c *gin.Context) {
	response.Error(c, fmt.Sprintf("api不存在 path:%s method:%s", c.Request.URL, c.Request.Method))
}

func GracefulExit(server *http.Server) {
	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-exit
	//生产环境下 防止程序关闭太快 k8s service 还没有将pod从负载中拿掉 请求过来没有处理 导致服务间断
	if !global.Config.Debug {
		time.Sleep(time.Second * 5)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("http server shutdown exit error:%s\n", err.Error())
	}
	fmt.Println("http server exit")
}
