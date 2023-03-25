package api

import (
	"encoding/gob"

	"github.com/ShadrackAdwera/go-auth/api/callback"
	"github.com/ShadrackAdwera/go-auth/api/login"
	"github.com/ShadrackAdwera/go-auth/api/logout"
	"github.com/ShadrackAdwera/go-auth/authenticator"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	auth   *authenticator.Authenticator
}

func NewServer(auth *authenticator.Authenticator) *Server {
	server := Server{}

	router := gin.Default()
	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.GET("/", server.home)
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/protected", IsAuthenticated, server.protected)
	router.GET("/logout", logout.Handler)

	server.router = router
	server.auth = auth

	return &server
}

func (s *Server) StartServer(address string) error {
	return s.router.Run(address)
}
