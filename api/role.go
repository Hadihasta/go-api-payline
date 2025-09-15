package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type createRoleRequest struct {
	RoleName string `json:"role_name" binding:"required,oneof=super_admin customer owner"`
}

func (server *Server) createRoles(ctx *gin.Context) {
	var req createRoleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	role, err := server.system.CreateRoles(ctx, req.RoleName)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, role)
}
