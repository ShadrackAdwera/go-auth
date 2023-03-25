package api

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (s *Server) protected(ctx *gin.Context) {
	session := sessions.Default(ctx)
	profile := session.Get("profile")

	fmt.Println(profile)

	ctx.JSON(http.StatusAccepted, gin.H{"message": "ping protected route"})
}
