package config

import "github.com/tal-tech/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Postgresql struct {
		DataSource string
	}
}
