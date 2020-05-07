package global

import "errors"

var ServiceNotFoundError = errors.New("service not found")
var MethodFoundError = errors.New("method not found")

// 新定义返回结果函数
func NewResp(code int, msg string) func(...interface{}) *JsonResponse {
	return func(i ...interface{}) *JsonResponse {
		length := len(i)
		jsonResponse := &JsonResponse{
			Status:  code,
			Message: msg,
		}
		if length > 0 {
			jsonResponse.Data = i[0]
		}
		return jsonResponse
	}
}

var (
	RespSuccess     = NewResp(0, "请求成功")
	RespNotAuth     = NewResp(1, "未授权")
	RespFailed      = NewResp(2, "请求失败")
	RespServerError = NewResp(3, "服务异常")
	RespNotFound    = NewResp(4, "资源未找到")
	RespOptFailed   = NewResp(5, "操作失败")
	RespNotLogin    = NewResp(6, "未登录")
	RespIllegal     = NewResp(7, "名词非法")
	// 无推荐的习题或视频
	RespEmptyGood      = NewResp(-1, "您本次作文表现出色，暂未发现问题")
	RespEmptyBad       = NewResp(-1, "很遗憾，您本次作文不合格，没有有效数据")
	RespEmpty          = NewResp(-1, "暂无数据")
	RespEmptyRecommend = NewResp(-2, "您本次考试没有薄弱点，暂无推荐哦")
)
