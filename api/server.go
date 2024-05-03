package api

import (
	"net/http"
    
	db "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/db/sqlc"
	"github.com/gin-gonic/gin"
)



type Server struct {
	Nps *db.Nps
	router *gin.Engine
}



func NewServer(nps *db.Nps) *Server {
	server := &Server{Nps: nps}
	router := gin.New()


	router.POST("/Recipients", server.createRecipient)

	server.router = router
	return server
}

func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}	

func errorResponse(c *gin.Context, err error) {
    c.JSON(http.StatusInternalServerError, gin.H{
        "error": err.Error(),
    })
	c.Abort()
	return
}