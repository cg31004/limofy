package errs

import "simon/limofy/service/internal/thirdparty/errortool"

var (
	ConvertDB = errortool.ConvertDB
	Parse     = errortool.Parse
	Equal     = errortool.Equal
)

var (
	ErrDB = errortool.ErrDB
)

var (
	commGroup                = errortool.Codes.Group()
	CommonUnknownError       = commGroup.Error("未知错误")
	CommonNoData             = commGroup.Error("查无资料")
	CommonRawSQLNotFound     = commGroup.Error("找不到执行档")
	CommonServiceUnavailable = commGroup.Error("系统维护中")
	CommonConfigureInvalid   = commGroup.Error("设置参数错误")
	CommonParseError         = commGroup.Error("解析失败")
)

var (
	requestGroup                  = errortool.Codes.Group()
	RequestParamInvalid           = requestGroup.Error("请求参数错误")
	RequestParamParseFailed       = requestGroup.Error("请求参数解析失败")
	RequestPageError              = requestGroup.Error("请求的页数错误")
	RequestParseError             = requestGroup.Error("解析失败")
	RequestParseTimeZoneError     = requestGroup.Error("時區解析錯誤")
	RequestFrequentOperationError = requestGroup.Error("频繁操作，请稍后再尝试")
)

var (
	fileServerGroup = errortool.Codes.Group()

	FileServerUploadFailed  = fileServerGroup.Error("图片上传失败")
	FileServerResponseNotOK = fileServerGroup.Error("图片库异常")
)

var (
	dataConvertGroup = errortool.Codes.Group()
	DataConvertError = dataConvertGroup.Error("格式轉換錯誤")
)
