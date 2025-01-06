package xerr

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TransError(ctx context.Context, err error) error {
	if isGrpcError(err) {
		if status.Code(err) >= codes.Canceled && status.Code(err) <= codes.Unauthenticated {
			return status.Error(status.Code(err), "common.grpc_err_"+fmt.Sprint(uint32(status.Code(err))))
		}
	}
	return err
}

func isGrpcError(err error) bool {
	if err == nil {
		return false
	}

	_, ok := err.(interface {
		GRPCStatus() *status.Status
	})

	return ok
}
