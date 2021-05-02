package config

import "errors"

type ErrorCode int64

const (
	ErrCodeNone                     ErrorCode = 0
	ErrCodeErrREQParamInvalid       ErrorCode = 4000 //非法参数
	ErrCodeErrSystemInnerException  ErrorCode = 5001 //系统错误
	ErrCodeErrBusinessException     ErrorCode = 7001 //业务错误
	ErrQueryFail                    ErrorCode = 7002 //查询失败
	ErrExportFail                   ErrorCode = 7002 //导出失败
	ErrCodeErrPermissionException   ErrorCode = 401  //权限错误
	ErrCodeErrUserMenuInfoException ErrorCode = 402  //用户信息错误
	ErrCodeErrUserNotLogin          ErrorCode = 403  //用户未登录
	ErrCodeDuplicateBpm             ErrorCode = 9001 //存在BPM重复工单
)

const (
	ErrMsgREQParamInvalid     = "参数错误"
	ErrMsgBusinessException   = "系统错误"
	ErrMsgUserNotLogin        = "尚未登录"
	ErrMsgUserWrongUser       = "账号或密码错误"
	ErrMsgCdKeyNotLegal       = "激活码不合法或已失效"
	ErrMsgNotExistProducts    = "不存在该产品"
	ErrMsgLackOfBalance       = "余额不足"
	ErrMsgAccountAlreadyExist = "账号已存在"
	ErrMsgHaveBought          = "已购买当前产品"
)

var (
	ErrREQParamInvalid = errors.New("参数错误")
)

const (
	ErrQueryFailMsg  = "查询失败，请稍后"
	ErrSystemMsg     = "服务器繁忙，请稍后重试"
	ErrExportFailMsg = "导出失败"
)

const (
	UserInfo = "UserInfo"
)
