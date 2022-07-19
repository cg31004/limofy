package mysqlcli

import (
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/dig"

	"simon/limofy/service/internal/config"
	"simon/limofy/service/internal/thirdparty/logger"
)

func NewDBClient(in digIn) IMySQLClient {
	return initWithConfig(in)
}

type digIn struct {
	dig.In

	AppConf   config.IAppConfig
	ServiceConf   config.IServiceConfig
	SysLogger logger.ILogger `name:"sysLogger"`
}
