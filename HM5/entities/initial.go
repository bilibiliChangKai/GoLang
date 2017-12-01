package entities

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var mydb *xorm.Engine

func init() {
	//https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	db, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:2048)/test?charset=utf8&parseTime=true")
	//db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:2048)/test?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}

	// 同步注册表
	err = db.Sync(new(UserInfo))
	if err != nil {
		panic(err)
	}

	mydb = db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
