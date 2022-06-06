package web

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gitlab.paradise-soft.com.tw/pdsbe/disposable/sport-quiz/service/internal/model/dto"
	"gitlab.paradise-soft.com.tw/pdsbe/disposable/sport-quiz/service/isnternal/model/bo"
	"net/http"

	"gitlab.paradise-soft.com.tw/pdsbe/disposable/sport-quiz/service/internal/errs"
	"gitlab.paradise-soft.com.tw/pdsbe/disposable/sport-quiz/service/internal/util/timelogger"
)

type IExampleCtrl interface {
	Get(ctx *gin.Context)
}

func newExample(in digIn) IExampleCtrl {
	return &exampleCtrl{
		in: in,
	}
}

type exampleCtrl struct {
	in digIn
}

func (ctrl *exampleCtrl) Get(ctx *gin.Context) {
	defer timelogger.LogTime(ctx)()

	req := &dto.ExampleCondByGet{}
	if err := ctrl.in.Request.Bind(ctx, &req); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, errs.RequestParamParseFailed)
		return
	}

	cond := &bo.ExampleCondByGet{}
	if err := copier.Copy(cond, req); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusInternalServerError, errs.DataConvertError)
		return
	}

	resp, pager, err := ctrl.in.ExampleUseCase.Get.Handle(ctx, cond)
	if err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusInternalServerError, err)
		return
	}

	dtoResp := make([]*dto.ExampleRespByGet, 0)
	if err := copier.Copy(&dtoResp, resp); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusInternalServerError, errs.DataConvertError)
		return
	}

	dtoPagerResp := &dto.PagerResp{}
	if err := copier.Copy(dtoPagerResp, pager); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusInternalServerError, errs.DataConvertError)
		return
	}

	ctrl.in.SetResponse.StandardResp(ctx, http.StatusOK, &dto.ListResp{
		List:  dtoResp,
		Pager: dtoPagerResp,
	})
}
