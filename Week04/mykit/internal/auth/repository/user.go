package repository

import (
	model "mykit/internal/auth/model"
)

func (r *Repository) Login(id string, psw string) (string, error) {
	user := new(model.User)

	// bussiness to Db
	_ = user
	token := "create token"

	return token, nil
}
