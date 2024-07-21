package port

import (
	"context"
	"strings"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/gofiber/fiber/v2"
	"github.com/vnworkday/gateway/internal/common/converter"
	"google.golang.org/grpc"
)

type Router interface {
	Register(router fiber.Router)
	Path() string
}

func NewGRPCClient[Req, IReq, IResp, Resp any](
	conn *grpc.ClientConn,
	service, method string,
	encodeRequest converter.ConvertFunc[Req, IReq],
	decodeResponse converter.ConvertFunc[IResp, Resp],
	opts ...grpctransport.ClientOption,
) endpoint.Endpoint {
	var resp Resp

	return grpctransport.NewClient(
		conn,
		service,
		toMethodName(method),
		func(ctx context.Context, in any) (any, error) {
			return converter.Convert(ctx, in, encodeRequest)
		},
		func(ctx context.Context, out any) (any, error) {
			return converter.Convert(ctx, out, decodeResponse)
		},
		&resp,
		opts...,
	).Endpoint()
}

func toMethodName(fullMethodName string) string {
	return fullMethodName[strings.LastIndex(fullMethodName, "/")+1:]
}
