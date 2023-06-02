package actor

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterActorStruct struct {
	actorRequestHandler RequestHandlerActorStruct
}

func NewRouter(
	dbCrud *gorm.DB,
) RouterActorStruct {
	return RouterActorStruct{
		actorRequestHandler: RequestHandler(
			dbCrud,
		),
	}
}

func (r RouterActorStruct) Handle(router *gin.Engine) {
	basepath := "v1/actor"
	actorRouter := router.Group(basepath)

	actorRouter.POST("/register",
		r.actorRequestHandler.CreateActor,
	)

	actorRouter.GET("/:id",
		r.actorRequestHandler.GetActorById,
	)
	actorRouter.GET("",
		r.actorRequestHandler.GetAllActor,
	)

	actorRouter.PUT("/:id",
		r.actorRequestHandler.UpdateActorById,
	)
	actorRouter.DELETE("/:id",
		r.actorRequestHandler.DeleteActorById,
	)
}