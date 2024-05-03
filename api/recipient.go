package api

import (
	"net/http"
	"time"

	db "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createRecipientRequest struct {
	OcdMasterId string    `json:"ocdMasterId" binding:"required"`
	Username    string    `json:"username" binding:"required"`
	Role        string    `json:"role" binding:"required"`
	CreatedAt   time.Time `json:"created_at" binding:"required"`
}

func (server *Server) createRecipient(ctx *gin.Context) {
	var req createRecipientRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.NpsRecipientCreationTxParams{
		CreateRecipientParams: db.CreateRecipientParams{
			OcdMasterId: req.OcdMasterId,
			Username:    req.Username,
			Role:        req.Role,
			CreatedAt:   req.CreatedAt,
		},
	}

	res, err := server.Nps.NpsRecipientCreationTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": res.Recipient})
	return
}
