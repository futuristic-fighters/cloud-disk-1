package user

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) Get(c *gin.Context) {
	c.JSON(http.StatusOK, "get")
}

func (h *Handler) Page(c *gin.Context) {
	c.JSON(http.StatusOK, "page")
}

func (h *Handler) Create(c *gin.Context) {
	c.JSON(http.StatusOK, "create")
}

func (h *Handler) Update(c *gin.Context) {
	c.JSON(http.StatusOK, "update")
}

func (h *Handler) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, "delete")
}
