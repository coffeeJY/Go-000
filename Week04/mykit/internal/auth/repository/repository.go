package repository

type Repository struct {
	// db *xorm.Engine
	// rd *redis.Client
}

type Server interface {
	Login(d string, psw string) (token string, err error)
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
