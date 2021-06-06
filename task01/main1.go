package main

import (
	"database/sql"
	"errors"
	"fmt"

	//_ "github.com/go-sql-driver/mysql"
	wraperrors "github.com/pkg/errors"
)

func main() {
	err := Query()
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println(err)
	}
}

//dao
func Query() error {
	var db *sql.DB
	//Todo: init db...

	//查询数据
	queryStr := "SELECT * FROM userinfo"
	rows, err := db.Query(queryStr)
	if err != nil {
		return err
	}
	colsNum := 10
	refs := make([]interface{}, colsNum)
	//todo: init refs

	for {
		if err := rows.Scan(refs...); err != nil {
			if err == sql.ErrNoRows {
				//wrap 原始错误,同时记录sql语句和调用栈
				return wraperrors.Wrapf(err, "sql query str:%s", queryStr)
			}
			return err
		}
	}
	return nil
}
