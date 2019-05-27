package user

import (
	"cloud-disk/logger"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	db     *sql.DB
	logger logger.Logger
}

func NewHandler(db *sql.DB, l logger.Logger) *Handler {
	return &Handler{db: db, logger: l}
}

func (h *Handler) Get(c *gin.Context) {
	c.Set("sql", "select * from users")
	c.Set("sql1", "select * from users")
	c.Set("sql2", "select * from users")
	h.logger.Write("test log", " adb  as")
	c.JSON(http.StatusOK, "get")
}

func (h *Handler) Page(c *gin.Context) {
	c.JSON(http.StatusOK, "page")
}

func (h *Handler) Create(c *gin.Context) {
	c.JSON(http.StatusOK, "create")
}

func (h *Handler) Update(c *gin.Context) {
	u := &User{
		Name: c.Param("name"),
		Type: 0,
	}

	if err := updateValidator(u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "update")
		return
	}

	c.JSON(http.StatusOK, "update")
}

func (h *Handler) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, "delete")
}
