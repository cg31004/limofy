package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"simon/limofy/service/internal/errs"
	"simon/limofy/service/internal/model/bo"
	"simon/limofy/service/internal/model/dto"
	"simon/limofy/service/internal/utils/timelogger"
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

	req := &dto.ExampleGet{}
	if err := ctrl.in.Request.Bind(ctx, &req); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, errs.RequestParamParseFailed)
		return
	}

	cond := &bo.ExampleGet{}
	if err := copier.Copy(cond, req); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusInternalServerError, errs.DataConvertError)
		return
	}

	resp, err := ctrl.in.ExampleUseCase.Get.Handle(ctx, cond)
	if err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusInternalServerError, err)
		return
	}

	dtoResp := make([]*dto.ExampleGet, 0)
	if err := copier.Copy(&dtoResp, resp); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusInternalServerError, errs.DataConvertError)
		return
	}

	ctrl.in.SetResponse.StandardResp(ctx, http.StatusOK, &dto.ListResp{
		List: dtoResp,
	})
}
