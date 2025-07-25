package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keshu12345/overlap/internal/model"
	"github.com/keshu12345/overlap/service"
)

type OverlapHandler struct {
	svc *service.OverlapService
}

func NewOverlapHandler(svc *service.OverlapService) *OverlapHandler {
	return &OverlapHandler{svc: svc}
}

func (h *OverlapHandler) CheckOverlap(c *gin.Context) {
	var req model.OverlapRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := h.svc.CheckOverlap(req.Range1, req.Range2)
	c.JSON(http.StatusOK, model.OverlapResponse{Overlap: result})
}
