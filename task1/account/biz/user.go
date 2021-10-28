package biz

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"go_advanced/task1/account/dao"
	"log"
)

// 1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
func GetUserInfo() {
	name, err := dao.GetUserNickName(22)
	if errors.Cause(err) == sql.ErrNoRows {
		// todo user not found
		return
	}
	if err != nil {
		// todo error handle
		log.Fatal(err)
	}
	fmt.Print(name)
}
