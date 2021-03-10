package http

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/alexander-melentyev/bizzen/internal/domain"
	"github.com/alexander-melentyev/bizzen/internal/pkg/query"
	"github.com/alexander-melentyev/bizzen/internal/pkg/uri"
	"github.com/gin-gonic/gin"
)

// Handler -.
type Handler struct {
	useCase domain.OrgUseCase
}

// NewHandler -.
func NewHandler(r *gin.RouterGroup, s domain.OrgUseCase) {
	h := &Handler{
		useCase: s,
	}

	org := r.Group("org")
	{
		org.POST("", h.Create)
		org.GET("", h.ReadAll)
		org.GET(":id", h.ReadByID)
		org.GET(":id/history", h.ReadHistoryByID)
		org.PUT(":id", h.UpdateByID)
		org.DELETE(":id", h.SoftDeleteByID)
	}
}

// Create -.
func (h *Handler) Create(c *gin.Context) {
	var org domain.Org

	if err := c.ShouldBindJSON(&org); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	if err := h.useCase.Create(c.Request.Context(), org); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	c.AbortWithStatus(http.StatusCreated)
}

// ReadAll -.
func (h *Handler) ReadAll(c *gin.Context) {
	var p query.Pagination

	if err := c.ShouldBind(&p); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	res, err := h.useCase.ReadAll(c.Request.Context(), p.Limit, p.Offset)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

// ReadByID -.
func (h *Handler) ReadByID(c *gin.Context) {
	var id uri.ID

	if err := c.ShouldBindUri(&id); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	res, err := h.useCase.ReadByID(c.Request.Context(), id.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

// ReadHistoryByID -.
func (h *Handler) ReadHistoryByID(c *gin.Context) {
	var id uri.ID

	if err := c.ShouldBindUri(&id); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	var p query.Pagination

	if err := c.ShouldBind(&p); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	res, err := h.useCase.ReadHistoryByID(c.Request.Context(), id.ID, p.Limit, p.Offset)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

// UpdateByID -.
func (h *Handler) UpdateByID(c *gin.Context) {
	var id uri.ID

	if err := c.ShouldBindUri(&id); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	var org domain.Org

	if err := c.ShouldBindJSON(&org); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	res, err := h.useCase.UpdateByID(c.Request.Context(), id.ID, org)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

// SoftDeleteByID -.
func (h *Handler) SoftDeleteByID(c *gin.Context) {
	var id uri.ID

	if err := c.ShouldBindUri(&id); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	if err := h.useCase.SoftDeleteByID(c.Request.Context(), id.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		return
	}

	c.AbortWithStatus(http.StatusOK)
}
