package xerr

import (
	gerr "github.com/gzorm/commons/core/xerr/types/err"
	"google.golang.org/grpc/status"
)

func NewInsufficientBalanceErr() *CodeError {
	err := &CodeError{}
	err.ErrCode = int32(ActiveBetCoinNoEnough)
	err.ErrMsg = "Insufficient valid bet amount!"
	return err
}

func NewDBUpdateErr() *CodeError {
	err := &CodeError{}
	err.ErrCode = int32(UpdateException)
	err.ErrMsg = "Database save exception"
	return err
}

func NewDBQueryErr() *CodeError {
	err := &CodeError{}
	err.ErrCode = int32(QueryException)
	err.ErrMsg = "Database query exception"
	return err
}

func NewDBCreateErr() *CodeError {
	err := &CodeError{}
	err.ErrCode = int32(CreateException)
	err.ErrMsg = "Database create exception"
	return err
}

func NewDBDeleteErr() *CodeError {
	err := &CodeError{}
	err.ErrCode = int32(DeleteException)
	err.ErrMsg = "Database delete exception"
	return err
}

func NewBetRepetitionErr() *CodeError {
	err := &CodeError{}
	err.ErrCode = int32(BetslipsRepetitionError)
	err.ErrMsg = "betslips save repetition error"
	return err
}
func NewDBUpdateAffectedZeroError() *CodeError {
	err := &CodeError{}
	err.ErrCode = int32(DbUpdateAffectedZeroError)
	err.ErrMsg = MapErrMsg(DbUpdateAffectedZeroError)
	return err
}

func NewRedisError() *CodeError {
	err := &CodeError{}
	err.ErrCode = int32(RedisException)
	err.ErrMsg = "redis err"
	return err
}

func ParseErrCodeFromErr(err error) ErrCode {
	if rpcErr, ok := status.FromError(err); ok {
		for _, v := range rpcErr.Details() {
			if detail, ok := v.(*gerr.Error); ok {
				return ErrCode(detail.ErrCode)
			}
		}
	}
	return ServerCommonError
}
