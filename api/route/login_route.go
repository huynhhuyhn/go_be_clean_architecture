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

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}
