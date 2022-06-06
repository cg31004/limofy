package cron

import (
	"sync"
)

type Job interface {
	Run()
}

type FuncJob func()

func (f FuncJob) Run() {
	f()
}

func NewCustomJobFunc(f func(), profile *Profile) *CustomJob {
	return NewCustomJob(FuncJob(f), profile)
}

func NewCustomJob(job Job, profile *Profile) *CustomJob {
	return &CustomJob{
		job:     job,
		profile: profile,
	}
}

type CustomJob struct {
	job     Job
	profile *Profile
	mx      sync.Mutex
}

func (cj *CustomJob) Run() {
	if !cj.canRun() {
		return
	}

	defer cj.finish()

	cj.job.Run()
}

func (cj *CustomJob) canRun() bool {
	cj.mx.Lock()
	defer cj.mx.Unlock()

	if !cj.profile.isRunning() {
		return false
	}

	// 不可重覆執行，且執行中
	if !cj.profile.isOverlapping() && cj.profile.isProcessing() {
		return false
	}

	cj.profile.processingAdd()

	return true
}

func (cj *CustomJob) finish() {
	cj.mx.Lock()
	cj.profile.processingDone()
	cj.mx.Unlock()
}
