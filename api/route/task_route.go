package route

import (
	"time"

	"go_be_clean_architecture/api/controller"
	"go_be_clean_architecture/bootstrap"
	"go_be_clean_architecture/domain"
	"go_be_clean_architecture/mongo"
	"go_be_clean_architecture/repository"
	"go_be_clean_architecture/usecase"

	"github.com/gin-gonic/gin"
)

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}
