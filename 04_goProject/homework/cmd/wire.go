package main

import (
	"github.com/google/wire"
	"learn/04_goProject/homework/internal/config"
	"learn/04_goProject/homework/internal/db"
)

func InitApp() (*App, error) {
	panic(wire.Build(config.Provider, db.Provider, NewApp)) // 调用wire.Build方法传入所有的依赖对象以及构建最终对象的函数得到目标对象
}
