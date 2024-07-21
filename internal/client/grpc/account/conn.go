package account

import (
	grpcutil "github.com/vnworkday/gateway/internal/common/grpc"
	"google.golang.org/grpc"
)

func NewConnection(params grpcutil.ConnectionParams) *grpc.ClientConn {
	return grpcutil.NewConnection(params, params.Config.GRPCAccountServiceURI)
}
