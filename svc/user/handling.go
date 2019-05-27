package user

import (
	"cloud-disk/cfg"
	"cloud-disk/logger"
	"database/sql"
	"encoding/json"
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

//example
func (h *Handler) Update(c *gin.Context) {

	c.Set("sql", "select * from users")
	c.Set("sql1", "select * from users")
	c.Set("sql2", "select * from users")
	h.logger.Write("test log", " adb  as")

	parseErr := c.Request.ParseForm()

	if parseErr != nil {
		c.JSON(http.StatusBadRequest, cfg.NewErrResponse(cfg.InvalidRequest))
		return
	}

	u := &User{}
	deErr := json.NewDecoder(c.Request.Body).Decode(u)
	if deErr != nil {
		c.JSON(http.StatusBadRequest, cfg.NewErrResponse(cfg.InvalidRequest))
		return
	}

	if err := updateValidator(u, h.db); err != nil {
		c.JSON(http.StatusUnprocessableEntity, cfg.NewErrResponse(err.Code()))
		return
	}

	c.JSON(http.StatusOK, cfg.NewResponse(cfg.UpdateUserSuccess, "update success"))
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

func (h *Handler) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, "delete")
}
