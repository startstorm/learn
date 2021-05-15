package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	v11 "learn/04_goProject/homework/internal/biz/api/v1"
	v22 "learn/04_goProject/homework/internal/biz/api/v2"
)

var Router *gin.Engine

//var Provider = wire.NewSet(CreatRouter)

func main(){

}

func CreatRouter(r *gin.Engine) (*gin.RouterGroup, *gin.RouterGroup){
	// 简单的路由组: v1
	v1 := r.Group("/v1")
	{
		v1.POST("/login", v11.LoginEndpoint)
		v1.POST("/submit", v11.SubmitEndpoint)
		v1.POST("/read", v11.ReadEndpoint)
	}

	// 简单的路由组: v2
	v2 := r.Group("/v2")
	{
		v2.POST("/login", v22.LoginEndpoint)
		v2.POST("/submit", v22.SubmitEndpoint)
		v2.POST("/read", v22.ReadEndpoint)
	}
	return v1,v2
}

func CreateDefaultGin() (r *gin.Engine, err error){
	r = gin.Default()
	if r == nil{
		return nil, errors.New("gin default fault")
	}
	return r,nil
}

func init(){

	Router, err := CreateDefaultGin()
	if err != nil{
		panic("gin defalt fault")
	}
	CreatRouter(Router)

}


