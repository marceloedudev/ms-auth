package grpc_exception

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func InternalServerExceptionGRPC(name string, err error) error {
	return status.Errorf(ParseGRPCErrStatusCode(err), fmt.Sprintf("%s: %v", name, err))
}

func ParseGRPCErrStatusCode(err error) codes.Code {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return codes.NotFound
	case errors.Is(err, context.DeadlineExceeded):
		return codes.DeadlineExceeded
	case errors.Is(err, context.Canceled):
		return codes.Canceled
	}
	return codes.Internal
}
