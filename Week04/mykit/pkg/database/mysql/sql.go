package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"go.uber.org/zap"
	"mykit/pkg/log"
	"time"
	"xorm.io/core"
)

type Config struct {
	Addr        string
	User        string
	Password    string
	DbName      string
	Parameters  string
	TablePrefix string

	MaxConn      int
	IdleConn     int
	Debug        bool
	IdleTimeout  int
	QueryTimeout int //查询时间
	ExecTimeout  int //执行时间
}

//user:password@(addr)/dbname?charset=utf8&parseTime=True&loc=Local
func NewMysql(c *Config) (db *xorm.Engine) {
	var err error
	URL := fmt.Sprintf("%s:%s@(%s)/%s?%s", c.User, c.Password, c.Addr, c.DbName, c.Parameters)
	db, err = xorm.NewEngine("mysql", URL)
	if err != nil {
		log.GetLogger().Error("[NewMysql] Open", zap.Any("conf", c), zap.Error(err))
		panic(err)
	}
	db.SetMaxIdleConns(c.IdleConn)
	db.SetMaxOpenConns(c.MaxConn)
	db.SetConnMaxLifetime(time.Duration(c.IdleTimeout) * time.Millisecond)
	db.ShowSQL(c.Debug)
	db.ShowExecTime(c.Debug)
	db.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, c.TablePrefix))
	db.SetColumnMapper(core.SnakeMapper{})
	if err = db.DB().Ping(); err != nil {
		log.GetLogger().Error("[NewMysql] Ping", zap.Any("conf", c), zap.Error(err))
		panic(err)
	}
	log.GetLogger().Info("[NewMysql] success", zap.Any("URL", URL))
	return
}
