package main

/*
为了避免以后更换"database/sql"时，业务层调用不用做修改，这里自定义自己错误：DaoErrNoRow
dao 返回错误时，用"github.com/pkg/errors" wrap 保存调用栈信息。方便上层%+v 查看调用栈信息。
上层只需要判断errors.Is(err,DaoErrNoRow) 就知道底层错误是 DaoErrNoRow. 代表 NoRow.
*/
import (
	"database/sql"
	"errors"
	"fmt"

	//_ "github.com/go-sql-driver/mysql"
	wraperrors "github.com/pkg/errors"
)

type DaoErrNoRow struct {
	msg string
	//err error
}

func (e *DaoErrNoRow) Error() string {
	return e.msg
}

/*
func (e *DaoErrNoRow) Unwrap() error {
	return e.err
}
*/

func main() {
	err := Query()
	if errors.Is(err, DaoNoRowErr) { //业务层引用dao库的DaoNoRowErr, 不需要sql?
		fmt.Println(err)
	}
	fmt.Println("done")
}

//dao
var DaoNoRowErr = &DaoErrNoRow{"no row result"}

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
				return wraperrors.Wrapf(DaoNoRowErr, "sql query str:%s, err:%s", queryStr, err)
			}
			return err
		}
	}
	return nil
}
