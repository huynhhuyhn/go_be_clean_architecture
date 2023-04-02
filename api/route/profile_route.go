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

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}
