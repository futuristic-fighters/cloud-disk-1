package user

import (
	"cloud-disk/logger"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func MakeSvc(db *sql.DB, r *gin.Engine, l logger.Logger) {
	rg := r.Group("/user")
	H := NewHandler(db, l)

	rg.GET("", H.Page)
	rg.GET("/:uid", H.Get)
	rg.POST("/:uid", H.Create)
	rg.PUT("/:uid", H.Update)
	rg.DELETE("/:uid", H.Delete)
}
