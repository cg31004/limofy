package web

import (
	"context"
	"net/http"
	"runtime"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"simon/limofy/service/internal/config"
	"simon/limofy/service/internal/thirdparty/logger"
)

func NewWebRestService(in restServiceIn) IService {
	self := &restService{
		in: in,
	}

	return self
}

type restService struct {
	in restServiceIn
}

type restServiceIn struct {
	dig.In

	AppConf     config.IAppConfig
	ServiceConf config.IServiceConfig
	SysLogger   logger.ILogger `name:"sysLogger"`
	AppLogger   logger.ILogger `name:"appLogger"`
}

type IService interface {
	Run()
}

func (s *restService) Run(ctx context.Context) {

	engine := s.newEngine()
	s.setRoutes(engine)

}

func (s *restService) newEngine() *gin.Engine {
	return gin.New()

}

func (s *restService) setRoutes(engine *gin.Engine) {
	engine.SetTrustedProxies([]string{})

	// setting middlewares
	engine.Use(
		gin.Logger(),
		gin.Recovery(),

	// cors middleware
	// response middleware
	)

	s.setPublicRoutes(engine)
	s.setPrivateRoutes(engine)
}

func (s *restService) setPublicRoutes(engine *gin.Engine) {
	s.setWebRoutes(engine) // Gateway 自己的功能
}

func (s *restService) setWebRoutes(engine *gin.Engine) {
	privateRouteGroup := engine.Group("")
	
	// 設定路由
	s.setApiRouters(privateRouteGroup)
}

func (s *restService) setPrivateRoutes(engine *gin.Engine) {
	privateRouteGroup := engine.Group("/_")

	// health check
	privateRouteGroup.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	// list error codes
	privateRouteGroup.GET("error-codes", getErrorCodes)

	// pprof
	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(1)
	pprof.Register(engine, "/_/debug")

	// prometheus
	// prometheusHandler := promhttp.Handler()
	// privateRouteGroup.GET("/metrics", func(c *gin.Context) {
	// 	prometheusHandler.ServeHTTP(c.Writer, c.Request)
	// })
}

func getErrorCodes(ctx *gin.Context) {
	type code struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}

	codes := errortool.Codes.List()

	resp := make([]code, len(codes))
	for i, v := range codes {
		resp[i] = code{
			Code:    v.GetCode(),
			Message: v.GetMessage(),
		}
	}

	ctx.JSON(http.StatusOK, resp)
}
