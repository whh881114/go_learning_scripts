package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=192.168.255.252 port=5432 user=wanghaohao password=MyPassword007. dbname=test sslmode=disable")
	checkErr(err)

	// dataSourceName支持格式如下：
	// user@unix(/path/to/socket)/dbname?charset=utf8
	// user:password@tcp(localhost:5555)/dbname?charset=utf8
	// user:password@/dbname
	// user:password@tcp([de:ad:be:ef::ca:fe]:80/dbname)

	// 插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) VALUES($1, $2, $3) RETURNING uid")
	checkErr(err)
	res, err := stmt.Exec("张三", "研发部门", "2017-09-09")
	checkErr(err)

	// pg不支持这个函数，因为它没有类似MySQL的自增ID
	// id, err := res.LastInsertId()
	// checkErr(err)
	// fmt.Println(id)

	var lastInsertId int
	err = db.QueryRow("INSERT INTO userinfo(username, departname, created) VALUES($1, $2, $3) RETURNING uid", "张三", "研发部门", "2017-09-09").Scan(&lastInsertId)
	checkErr(err)
	fmt.Println("最后插入id = ", lastInsertId)

	// 更新数据
	stmt, err = db.Prepare("UPDATE userinfo SET username=$1 WHERE uid=$2")
	checkErr(err)
	res, err = stmt.Exec("zuolanupdate", 1)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	// 查询数据
	rows, err := db.Query("SELECT * FROM userinfo")

	for rows.Next() {
		var uid int
		var username string
		var departname string
		var created string
		err = rows.Scan(&uid, &username, &departname, &created)
		checkErr(err)
		fmt.Printf("%-3d%-20s%-15s%-10s\n", uid, username, departname, created)
	}

	// 删除数据
	stmt, err = db.Prepare("DELETE FROM userinfo WHERE uid=$1")
	checkErr(err)
	res, err = stmt.Exec(1)
	checkErr(err)
	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
