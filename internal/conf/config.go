package conf

import (
	"time"

	"github.com/vnworkday/config"
)

type Conf struct {
	ServiceName string `config:"service_name"`
	Profile     string `config:"profile"`

	HTTPPathPrefix           string        `config:"http_path_prefix"`
	HTTPRequestReadTimeout   time.Duration `config:"http_request_read_timeout"`
	HTTPResponseWriteTimeout time.Duration `config:"http_response_write_timeout"`

	GRPCMaxMessageSizeMB int           `config:"grpc_max_message_size_mb"`
	GRPCKeepaliveTime    int           `config:"grpc_keepalive_time"`
	GRPCKeepaliveTimeout int           `config:"grpc_keepalive_timeout"`
	GRPCCallTimeout      time.Duration `config:"grpc_call_timeout"`

	GRPCAccountServiceURI string `config:"grpc_account_service_uri"`
}

func New() (*Conf, error) {
	return config.LoadConfig[Conf](new(Conf))
}
