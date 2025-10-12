package api

// type createUserRequest struct {
// 	RoleID int `json:"role_id" binding:"required,oneof=super_admin customer owner"`
// 	Email string `json:"email" binding:"required,email"`
// 	PhoneNumber string `json:"phone_number" binding:"required"`
// 	Name string `json:"name" binding:"required"`
// }

// func (server *Server) createUser(ctx *gin.Context) {
// 	var req createUserRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	role, err := server.system.createUser(ctx, req.RoleName)

// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, role)
// }
