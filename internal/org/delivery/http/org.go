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
	org domain.OrgUseCase
}

// NewHandler -.
func NewHandler(r *gin.RouterGroup, o domain.OrgUseCase) {
	h := &Handler{
		org: o,
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

// Create - handler for creating a new organization.
// @Tags org
// @Summary Creating a new organization
// @ID create-org
// @Accept  json
// @Produce  json
// @Param input body domain.OrgDTO true "organization data"
// @Success 201 {object} object "created"
// @Failure 400 {object} object "bad request"
// @Failure 500 {object} object "internal server error"
// @Router /org [POST]
func (h *Handler) Create(c *gin.Context) {
	var org domain.OrgDTO

	if err := c.ShouldBindJSON(&org); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	if err := h.org.Create(c.Request.Context(), org); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	c.AbortWithStatus(http.StatusCreated)
}

// ReadAll - handler for getting a organizations list.
// @Tags org
// @Summary Getting a organizations list
// @ID read-all-org
// @Accept  json
// @Produce  json
// @Param limit query int true "limit"
// @Param offset query int true "offset"
// @Success 200 {object} object "ok"
// @Failure 400 {object} object "bad request"
// @Failure 500 {object} object "internal server error"
// @Router /org [GET]
func (h *Handler) ReadAll(c *gin.Context) {
	var p query.Pagination

	if err := c.ShouldBindQuery(&p); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	res, err := h.org.ReadAll(c.Request.Context(), p.Limit, p.Offset)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

// ReadByID - handler for getting organization by ID.
// @Tags org
// @Summary Getting organization by ID
// @ID read-org-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Organization ID"
// @Success 200 {object} object "ok"
// @Failure 400 {object} object "bad request"
// @Failure 500 {object} object "internal server error"
// @Router /org/{id} [GET]
func (h *Handler) ReadByID(c *gin.Context) {
	var id uri.ID

	if err := c.ShouldBindUri(&id); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	res, err := h.org.ReadByID(c.Request.Context(), id.ID)
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

// ReadHistoryByID - handler for getting organization row changes in table.
// @Tags org
// @Summary Getting organization row changes in table
// @ID read-org-history-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Organization ID"
// @Success 200 {object} object "ok"
// @Failure 400 {object} object "bad request"
// @Failure 500 {object} object "internal server error"
// @Router /org/{id}/history [GET]
func (h *Handler) ReadHistoryByID(c *gin.Context) {
	var id uri.ID

	if err := c.ShouldBindUri(&id); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	var p query.Pagination

	if err := c.ShouldBindQuery(&p); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	res, err := h.org.ReadHistoryByID(c.Request.Context(), id.ID, p.Limit, p.Offset)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

// UpdateByID - updating organization data.
// @Tags org
// @Summary Updating organization data
// @ID update-org-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Organization ID"
// @Param input body domain.OrgDTO true "organization data"
// @Success 200 {object} object "ok"
// @Failure 400 {object} object "bad request"
// @Failure 404 {object} object "not found"
// @Failure 500 {object} object "internal server error"
// @Router /org/{id} [PUT]
func (h *Handler) UpdateByID(c *gin.Context) {
	var id uri.ID

	if err := c.ShouldBindUri(&id); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	var org domain.OrgDTO

	if err := c.ShouldBindJSON(&org); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	res, err := h.org.UpdateByID(c.Request.Context(), id.ID, org)
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

// SoftDeleteByID - filling deletion data.
// @Tags org
// @Summary Filling deletion data
// @ID soft-delete-org-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Organization ID"
// @Success 200 {object} object "ok"
// @Failure 400 {object} object "bad request"
// @Failure 404 {object} object "not found"
// @Failure 500 {object} object "internal server error"
// @Router /org/{id} [DELETE]
func (h *Handler) SoftDeleteByID(c *gin.Context) {
	var id uri.ID

	if err := c.ShouldBindUri(&id); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	if err := h.org.SoftDeleteByID(c.Request.Context(), id.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		return
	}

	c.AbortWithStatus(http.StatusOK)
}
