package redis

import (
	"context"
	rds "github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	log "mykit/pkg/log"
)

type Config struct {
	Addrs    []string
	Addr     string
	Password string
	Db       int
}

func NewRedis(c *Config) (db *rds.Client) {
	if c.Addr != "" {
		db = rds.NewClient(&rds.Options{
			Addr:     c.Addr,
			Password: c.Password,
			DB:       c.Db,
		})
		if err := db.Ping(context.Background()).Err(); err != nil {
			log.GetLogger().Error("[NewRedis] Ping", zap.Any("Addr", c.Addr), zap.Error(err))
			return
		}
		log.GetLogger().Info("[NewRedis] success", zap.Any("Addr", c.Addr))
	}
	return
}

func NewClusterRedis(c *Config) (db *rds.ClusterClient) {
	if len(c.Addrs) > 0 {
		db = rds.NewClusterClient(&rds.ClusterOptions{
			Addrs:    c.Addrs,
			Password: c.Password,
		})
		if err := db.Ping(context.Background()).Err(); err != nil {
			log.GetLogger().Error("[NewClusterRedis] Ping", zap.Any("Addr", c.Addr), zap.Error(err))
			return
		}
		log.GetLogger().Info("[NewClusterRedis] success", zap.Any("Addr", c.Addr))
	}
	return
}
