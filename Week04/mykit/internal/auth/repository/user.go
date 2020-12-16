package repository

import (
	// "encoding/json"
	// "fmt"
	// "github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	model "mykit/internal/auth/model"
	"mykit/pkg/log"
)

func (r *Repository) GetUserById(id int64) (user *model.User, err error) {
	user = new(model.User)
	// user, err = r.GetCacheUserById(id)
	// if err != nil {
	// 	if err == redis.Nil { //没有查到
	// 		user, err = r.GetDbUserById(id)
	// 		return
	// 	} else { //redis 错误 直接把错误抛出
	// 		log.GetLogger().Error("[GetCacheUserById] Get", zap.Any("id", id), zap.Error(err))
	// 		return
	// 	}
	// }
	log.GetLogger().Debug("[GetCacheUserById]", zap.Any("data", user))
	return user, nil
}
