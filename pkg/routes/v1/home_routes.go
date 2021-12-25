package v1

import (
	"github.com/Xwudao/junet/app"
	"github.com/gin-gonic/gin"
)

type HomeRoutes struct {
	Engine *app.App
}

func (r *HomeRoutes) SetUpRoutes() {
	v1 := r.Engine.Group("/")
	{
		v1.GET("", r.index())
	}
}

func (r *HomeRoutes) index() gin.HandlerFunc {
	return func(c *gin.Context) {
		app.SendJson(c, app.OkRtn("ok"))
	}
}
