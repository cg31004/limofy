package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"simon/limofy/service/internal/constant"
	"simon/limofy/service/internal/utils/timelogger"
)

const (
	RespFormat_Standard = "RespFormat_Standard"

	Resp_Format  = "Resp_Format"
	Resp_Data    = "Resp_Data"
	Resp_Status  = "Resp_Status"
	Resp_Success = "success"
)

type IResponseMiddleware interface {
	Handler(ctx *gin.Context)
}

type responseMiddleware struct {
	in digIn
}

func (m *responseMiddleware) Handle(ctx *gin.Context) {
	if m.in.SysLogger.Level() == "debug" {
		ctx.Set(timelogger.ContextKey, timelogger.NewTimeLogger())
	}

	ctx.Set(constant.App_ChainID, uuid.New().String())

	ctx.Next()

	switch ctx.GetString(Resp_Format) {
	case RespFormat_Standard:
		m.standardResponse(ctx)
	default:
	}
}

func (m *responseMiddleware) standardResponse(ctx *gin.Context) {
	resp := m.gen
}
