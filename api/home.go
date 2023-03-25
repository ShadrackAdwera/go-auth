package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) home(ctx *gin.Context) {

	ctx.JSON(http.StatusAccepted, gin.H{"message": "unprotected home path"})
}
