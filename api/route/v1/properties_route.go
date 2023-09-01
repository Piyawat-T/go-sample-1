package route

import (
	"time"

	"github.com/Piyawat-T/go-sample-1/api/controller"
	"github.com/Piyawat-T/go-sample-1/bootstrap"
	"github.com/Piyawat-T/go-sample-1/client"
	"github.com/Piyawat-T/go-sample-1/usecase"
	"github.com/gin-gonic/gin"
)

func PropertiesRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	pc := client.NewPropertiesClient()
	cont := controller.PropertiesController{
		PropertiesUsecase: usecase.NewPropertiesUsecase(pc, timeout),
		Env:               env,
	}
	group.GET("/:application/:profile", cont.GetConfiguration)
}
