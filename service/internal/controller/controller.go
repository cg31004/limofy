package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"simon/limofy/service/internal/controller/handler"
	"simon/limofy/service/internal/controller/middleware"
	"simon/limofy/service/internal/core/usecase/example"
	"simon/limofy/service/internal/thirdparty/logger"
)

func NewWebRestCtl(in digIn) digOut {
	self := &packet{
		in: in,
		digOut: digOut{
			ExampleCtrl: newExample(in),
		},
	}

	return self.digOut
}

type packet struct {
	in digIn

	digOut
}

type digIn struct {
	dig.In

	SysLogger   logger.ILogger `name:"sysLogger"`
	Request     handler.IRequestParse
	SetResponse response `optional:"true"`

	ExampleUseCase exampleUseCaseIn
}

type digOut struct {
	dig.Out

	ExampleCtrl IExampleCtrl
}

type exampleUseCaseIn struct {
	dig.In

	Get example.IGetExampleUseCase
}

type response struct{}

func (response) StandardResp(ctx *gin.Context, statusCode int, data interface{}) {
	middleware.SetResp(ctx, middleware.RespFormat_Standard, statusCode, "0", data)
}
