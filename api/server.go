package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/go-api-payline/db/sqlc"
)

// Server HTTP request for our service
type Server struct {
	system *db.Queries
	router *gin.Engine
}

// create a new server and setup routing
func NewServer(system *db.Queries) *Server {
	server := &Server{system: system}
	router := gin.Default()


	// roles
	router.POST("/role",server.createRoles)
	router.GET("/role/:id",server.GetRoles)
	router.GET("/role",server.ListRoles)
	// users
	// router.GET("/users",server.createUser)
	// router.GET("/users",server.ListUsers)


	// add routes to router
	server.router = router
	return server
}


func (server *Server) Start(address string) error {
	return server.router.Run(address)
}


func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
