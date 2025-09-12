package template

// WebController is a template for a web controller
const WebController = `package controller

import (
	"net/http"
	"strconv"

	"{{ .module }}/internal/transport/{{ .layer }}/{{ .folder }}/request"
	"{{ .module }}/internal/transport/{{ .layer }}/{{ .folder }}/response"
	"{{ .module }}/internal/transport/{{ .layer }}/{{ .folder }}/view"
	"{{ .module }}/pkg/logger"

	"github.com/gin-gonic/gin"
)

// {{ .entity }}Controller ...
type {{ .entity }}Controller struct {
	log *logger.Logger
}

// New{{ .entity }}Controller ...
func New{{ .entity }}Controller(r *gin.RouterGroup, log *logger.Logger) *{{ .entity }}Controller {
	ctrl := &{{ .entity }}Controller{
		log: log,
	}

	r.GET(":id", ctrl.GetByID)
	r.GET("", ctrl.GetAll)
	r.POST("", ctrl.Create)

	return ctrl
}

// GetByID ...
func (c *{{ .entity }}Controller) GetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.log.Error("Invalid ID parameter", logger.Error(err))
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid ID parameter"))
		return
	}

	ctx.JSON(http.StatusOK, view.{{ .entity }}Response{
		ID: id,
	})
}

// GetAll ...
func (c *{{ .entity }}Controller) GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, view.{{ .entity }}ListResponse{
		Items: []view.{{ .entity }}Response{},
	})
}

// Create ...
func (c *{{ .entity }}Controller) Create(ctx *gin.Context) {
	var req request.Create{{ .entity }}Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.log.Error("Invalid request body", logger.Error(err))
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request body"))
		return
	}

	ctx.JSON(http.StatusCreated, view.{{ .entity }}Response{
		ID: 1, // Simulated created ID
	})
}
`
