package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

func GetUserNickName(uid int) (string, error) {
	db, err := sql.Open("mysql",
		"root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		return "", err
	}
	defer db.Close()
	var nickName string
	if err := db.QueryRow("select nick_name from sys_user where user_id = ?", uid).Scan(&nickName); err != nil {
		return "", errors.Wrapf(err, "user is not found")
	}
	return nickName, nil
}
