package main

import (
	"database/sql"
	"learn/04_goProject/homework/api"
	"log"
)

type App struct { // 最终需要的对象
	db *sql.DB
}

func NewApp(db *sql.DB) *App {
	return &App{db: db}
}

func main() {
	app, err := InitApp() // 使用wire生成的injector方法获取app对象
	if err != nil {
		log.Fatal(err)
	}
	var version string
	row := app.db.QueryRow("SELECT VERSION()")
	if err := row.Scan(&version); err != nil {
		log.Fatal(err)
	}
	log.Println(version)

	api.Router.Run()
}

func init() {


}
