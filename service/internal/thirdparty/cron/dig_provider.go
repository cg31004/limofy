package cron

import (
	"context"
	"time"

	"go.uber.org/dig"

	"gitlab.paradise-soft.com.tw/glob/cron"
	"simon/limofy/service/internal/thirdparty/logger"
)

type ICronJob interface {
	Start()
	AddFunc(spec string, cmd func(), opt ...cron.FuncCronOpt) (*cron.Profile, error)
	AddScheduleFunc(duration time.Duration, cmd func(), opt ...cron.FuncCronOpt) (*cron.Profile, error)
}

type cronJob struct {
	cron *cron.Cron
}

type digIn struct {
	dig.In

	SysLogger logger.ILogger `name:"sysLogger"`
}

func NewCronJob(in digIn) ICronJob {
	cron := cron.New()
	ctx := context.Background()
	in.SysLogger.Info(ctx, "serve start [job]")
	return &cronJob{cron: cron}
}

func (job *cronJob) Start() {
	job.cron.Start()
}

func (job *cronJob) AddScheduleFunc(duration time.Duration, cmd func(), opt ...cron.FuncCronOpt) (*cron.Profile, error) {
	return job.cron.AddScheduleFunc(cron.Every(duration), cmd, opt...)
}

func (job *cronJob) AddFunc(spec string, cmd func(), opt ...cron.FuncCronOpt) (*cron.Profile, error) {
	return job.cron.AddFunc(spec, cmd, opt...)
}
