package xerr

import (
	"github.com/gzorm/commons/core/xerr/types/err"
)

/**
常用通用固定错误
*/

type CodeError struct {
	err.Error
	// errCode uint32
	// errMsg  string
}

// 返回给前端的错误码
func (e *CodeError) GetErrCode() int32 {
	return e.ErrCode
}

// 返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	return e.ErrMsg
}

// func (e *CodeError) Error() string {
// 	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
// }

func NewErrCodeMsg(errCode int32, errMsg string) *CodeError {
	err := &CodeError{}
	err.ErrCode = errCode
	err.ErrMsg = errMsg
	return err
}
func NewErrCode(errCode int32) *CodeError {
	err := &CodeError{}
	err.ErrCode = errCode
	err.ErrMsg = MapErrMsg(ErrCode(errCode))
	return err
}

func NewErrMsg(errMsg string) *CodeError {
	err := &CodeError{}
	err.ErrCode = int32(ServerCommonError)
	err.ErrMsg = errMsg
	return err
}
