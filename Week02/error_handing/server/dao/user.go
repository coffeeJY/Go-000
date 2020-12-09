package dao

import (
	myerr "error_handing/server/err"
	"github.com/pkg/errors"
)

type User struct {
	ID   string
	Name string
	Age  int
}

func QueryUserByID(ID string) (*User, error) {
	row := MysqlDB.QueryRow("select id from t_user where id = ?", ID)
	resUser := User{}
	if err := row.Scan(&resUser.ID); err != nil {
		return nil, errors.Wrap(myerr.ErrNoRowsSql, err.Error())
	}
	return &resUser, nil
}
