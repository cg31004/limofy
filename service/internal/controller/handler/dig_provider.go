package handler

import (
	"go.uber.org/dig"

	"simon/limofy/service/internal/config"
	"simon/limofy/service/internal/thirdparty/logger"
)

func NewRequestParse(in digIn) IRequestParse {
	return &requestParseHandler{
		in: in,
	}
}

type digIn struct {
	dig.In

	AppConf   config.IAppConfig
	OpsConf   config.ServiceConfig
	SysLogger logger.ILogger `name:"sysLogger"`
}
