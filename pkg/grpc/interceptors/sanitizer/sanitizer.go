package sanitizer

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ISanitize interface {
	Sanitize() error
}

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if err := sanitize(ctx, req); err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

func sanitize(ctx context.Context, reqOrRes interface{}) (err error) {
	switch v := reqOrRes.(type) {
	case ISanitize:
		err = v.Sanitize()
	}

	if err == nil {
		return nil
	}

	return status.Error(codes.InvalidArgument, err.Error())
}
