package conf

import (
	t "github.com/BurntSushi/toml"
	"mykit/pkg/database/mysql"
	"mykit/pkg/database/redis"
	"mykit/pkg/log"
)

var (
	Conf = new(Config)
)

type Config struct {
	Development     bool
	SnowFlakeId     int64
	UseRemotelyConf bool
	Log             *log.Options
	Mysql           *mysql.Config
	Redis           *redis.Config
	Grpc            *grpcConfig
	Http            *httpConf
	Etcd            *etcdConf
}

type httpConf struct {
	Port int
}

type etcdConf struct {
	Addrs []string
}
type grpcConfig struct {
	SrvName          string
	Addrs            []string
	RegisterTTL      int
	RegisterInterval int
}

func Init(confPath string) error {
	return local(confPath)
}

func local(confPath string) (err error) {
	if _, err = t.DecodeFile(confPath, &Conf); err != nil {
		panic(err)
	}
	Conf.Mysql.Debug = Conf.Development
	return
}
