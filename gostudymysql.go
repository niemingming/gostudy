//golang实现mysql连接
package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"time"
)

//错误校验
func checkErr(err error) {
	if err != nil {
		panic(err) //抛出异常
	}
}

func main() {
	//dsn数据源名称
	//user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname?charset=utf8mb4,utf8
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/monitor_dump?charset=utf8")
	checkErr(err)
	/*建标*/
	//加密处理
	pwd := []byte("test")
	p := md5.Sum(pwd)
	fmt.Println(hex.EncodeToString(p[:]))
	ctime := time.Now().Format("2006-01-01 15:00:00")
	fmt.Println(ctime)
	// stmst = db.Prepare(query)
	//查询数据
	rows, err := db.Query("select * from monitor_quota_defined")
	checkErr(err)
	for rows.Next() {
		var (
			id          int
			expressioin string
			name        string
			content     string
			active      string
		)
		err = rows.Scan(&id, &expressioin, &name, &content, &active)
		//打印)
		fmt.Println("id:", id, ";expression:", expressioin, ";name:", name, ";content:", content, ";active", active)
	}
}
