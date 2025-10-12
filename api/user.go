package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/go-api-payline/db/sqlc"
)

// 🧠 NOTE: Kenapa kita punya 2 struct (createUserRequest & db.CreateUserParams)?
// ============================================================================
// 1️⃣ createUserRequest
//    → Ini khusus untuk menerima data dari client (HTTP JSON request).
//    → Dideklarasikan di layer API agar bisa:
//        - Validasi dengan tag `binding:"required"`
//        - Mapping field sesuai format JSON (`json:"..."`)
//        - Menghindari tipe database seperti sql.NullString yang tidak bisa di-bind dari JSON.
//    → Singkatnya: Struct ini hanya untuk API request body.
//
// 2️⃣ db.CreateUserParams
//    → Ini struct yang otomatis dibuat oleh sqlc (untuk query ke database).
//    → Formatnya mengikuti struktur tabel di PostgreSQL.
//    → Digunakan untuk mengirim data ke database (bukan untuk menerima dari user).
//
// Kenapa tidak langsung pakai db.CreateUserParams di API?
// -------------------------------------------------------
// Karena:
//    - sqlc menggunakan tipe seperti sql.NullString, sql.NullInt64, dll.
//      yang tidak bisa otomatis diisi dari JSON oleh Gin.
//    - Struktur DB bisa berbeda dari struktur API (misalnya password_hash, created_at, dll).
//    - Layer API dan layer DB sebaiknya terpisah (separation of concern).
//
// Maka flow-nya seperti ini:
//    [HTTP Request JSON] -> createUserRequest -> mapping -> db.CreateUserParams -> SQL Query
//
// Dengan cara ini, API tetap fleksibel & aman meskipun struktur database berubah.
// Jadi, kita butuh kedua struct ini untuk tujuan yang berbeda.

type createUserRequest struct {
	RoleID      int    `json:"role_id" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Name        string `json:"name" binding:"required"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		RoleID:      int64(req.RoleID),
		Email:       sql.NullString{String: req.Email, Valid: req.Email != ""},
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
	}

	user, err := server.system.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}