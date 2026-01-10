package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HomeHandler struct {
	logger *zap.Logger
}

func NewHomeHandler(logger *zap.Logger) *HomeHandler {
	return &HomeHandler{
		logger: logger,
	}
}

func (h *HomeHandler) HandlePage(c *gin.Context) {
	h.logger.Info("Serving homepage")
	c.HTML(http.StatusOK, "index.html", gin.H{"message": "page successfully loaded"})
}

