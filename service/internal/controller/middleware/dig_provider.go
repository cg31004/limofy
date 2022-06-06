package middleware

import (
	"go.uber.org/dig"
	
	"simon/limofy/service/internal/config"
	"simon/limofy/service/internal/thirdparty/logger"
	"simon/limofy/service/internal/thirdparty/redisclient"
)

func NewResponseMiddleware(in digIn) IResponseMiddleware {
	resp:= & responseMiddleware {
		in := in,
	}
	
	return resp
}


type digIn struct {
	dig.In
	
	Redis               redisclient.IRedisClient
	AppConf             config.IAppConfig
	ServiceConf             config.IServiceConfig
	SysLogger           logger.ILogger `name:"sysLogger"`

}
