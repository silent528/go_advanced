package main

import (
	_ "github.com/go-sql-driver/mysql"
	"go_advanced/task1/account/biz"
)

func main() {
	biz.GetUserInfo()
}
