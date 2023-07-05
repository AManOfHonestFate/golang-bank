package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/techschool/simplebank/db/sqlc"
)

// Server serves http requests
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server
// if router is nil, will set up default router
func NewServer(store *db.Store, router *gin.Engine) *Server {
	if router == nil {
		router = gin.Default()
	}

	server := &Server{
		store:  store,
		router: router,
	}

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)

	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
