package main

import (
	E "errors"
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	err0 := t1()
	err := errors.Wrap(err0, "附加信息")
	if err != nil {
		//打印错误需要%+v才能详细输出
		fmt.Printf("err :%+v\n", err)
	}

	fmt.Println("Hello world")
}

func t1() error {
	return E.New("错误")
}
