package route

import (
	"time"

	"github.com/Piyawat-T/go-sample-1/api/controller"
	"github.com/Piyawat-T/go-sample-1/bootstrap"
	"github.com/gin-gonic/gin"
)

func HomeRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	sc := controller.HomeController{
		Env: env,
	}
	group.GET("/ping", sc.Home)
}
