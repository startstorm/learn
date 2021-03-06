/*
*	因为错误在DAO层，不在业务层，而且该错误只是查询数据没有找到，不算是个错误
*	该错误应该降级处理，返回nil
*/
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"log"
)
var (
	db *sql.DB
	err error
)

// 测试函数
func test() error {
	var username string
	// 查询不存在的一行,id=4不存在
	err := db.QueryRow("select username from sys_users where id = ?", 4).Scan(&username)
	if err != nil {
		if err == sql.ErrNoRows{
			// 处理错误
			log.Fatal("sql.ErrNoRows", err)
			return nil
		}else{
			log.Fatal(err)
			return errors.Wrap(err, "no sql.ErrNoRows err")
		}
	}
	return nil
}

// 初始化
func init()  {
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/ginvueadmin")
	err = db.Ping()
	// 检查数据库是否连接成功
	if err != nil {
		fmt.Println("数据库连接失败")
		log.Fatal(err)
	}
	fmt.Println("数据库连接成功")
}

func main() {
	err = test()
	if err != nil{
		// 处理错误
		fmt.Println(err)
	}
	defer db.Close()
}