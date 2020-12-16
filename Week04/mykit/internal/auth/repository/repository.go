package repository

import (
	// "github.com/go-redis/redis/v8"
	// "github.com/go-xorm/xorm"
	"mykit/internal/auth/model"
)

type Repository struct {
	// db *xorm.Engine
	// rd *redis.Client
}

type Server interface {
	GetUserById(id int64) (user *model.User, err error)
	Close()
}

// Server 可定义那些接口向外暴露
func NewRepository() (r Server) {
	r = &Repository{}
	return r
}

func (r *Repository) Close() {
	// _ = r.db.Close()
	// _ = r.rd.Close()
}
