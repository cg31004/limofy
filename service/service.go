package service

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"simon/limofy/service/internal/app/web"
	"simon/limofy/service/internal/binder"
	"simon/limofy/service/internal/config"
	"simon/limofy/service/internal/thirdparty/logger"
)

func Run() {
	binder := binder.New()
	if err := binder.Invoke(initService); err != nil {
		panic(err)
	}

	select {}
}

type digIn struct {
	dig.In

	AppConf     config.IAppConfig
	ServiceConf config.IServiceConfig
	SysLogger   logger.ILogger `name:"sysLogger"`

	WebRestService web.IService
}

func initService(in digIn) {
	ctx := context.Background()

	serverInterrupt(ctx, in)
	ginMode(in)
	in.SysLogger.Info(ctx, fmt.Sprintf("[Build Info] %s", getBuildInfo()))

	go in.WebRestService.Run(ctx)
}

func serverInterrupt(ctx context.Context, in digIn) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)

	go func() {
		select {
		case c := <-interrupt:
			in.SysLogger.Warn(ctx, fmt.Sprintf("Server Shutdown, osSignal: %v", c))
			os.Exit(0)
		}
	}()

}

func ginMode(in digIn) {
	gin.DisableConsoleColor()
	if in.AppConf.GetGinConfig().DebugMode == false {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}
