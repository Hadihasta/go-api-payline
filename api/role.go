package api

import (
	"database/sql"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/go-api-payline/db/sqlc"
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

type getRoleRequest struct {
	// samakan tipe data dengan yang di generate oleh sqlc di role.sql.go
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetRoles(ctx *gin.Context) {
	var req getRoleRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		fmt.Println("error console: ", err)
		return
	}

	role, err := server.system.GetRoles(ctx, req.ID)
	if err != nil {
		fmt.Println("error console: ", role)
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse((err)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse((err)))
		return
	}

	ctx.JSON(http.StatusOK, role)
}

type listRoleRequest struct {
	// samakan tipe data dengan yang di generate oleh sqlc di role.sql.go
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) ListRoles(ctx *gin.Context) {
	var req listRoleRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		fmt.Println("error console: ", err)
		return
	}

	arg := db.ListRolesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	role, err := server.system.ListRoles(ctx, arg)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse((err)))
		return
	}

	ctx.JSON(http.StatusOK, role)
}
