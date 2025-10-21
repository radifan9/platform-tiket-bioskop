package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/radifan9/platform-tiket-bioskop/models"
	"github.com/radifan9/platform-tiket-bioskop/pkg"
	"github.com/radifan9/platform-tiket-bioskop/repositories"
	"github.com/radifan9/platform-tiket-bioskop/utils"
	"github.com/redis/go-redis/v9"
)

type UserHandler struct {
	ur *repositories.UserRepository
	ac *utils.AuthCacheManager
}

func NewUserHandler(ur *repositories.UserRepository, rdb *redis.Client) *UserHandler {
	return &UserHandler{
		ur: ur,
		ac: utils.NewAuthCacheManager(rdb),
	}
}

func (u *UserHandler) Register(ctx *gin.Context) {
	var user models.RegisterUser
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Hash password (contoh)
	// "password": "ceganssangar123(DF&&"
	// format : email + sangar123(DF&&
	hashCfg := pkg.NewHashConfig()
	hashCfg.UseRecommended()
	hashedPassword, err := hashCfg.GenHash(user.Password)
	if err != nil {
		utils.HandleError(ctx, http.StatusInternalServerError, "failed to hash password", err.Error())
		return
	}

	newUser, err := u.ur.CreateUser(ctx, user.Email, hashedPassword)
	if err != nil {
		log.Println("error : ", err)
		utils.HandleError(ctx, http.StatusConflict, "failed to register", err.Error())
		return
	}

	utils.HandleResponse(ctx, http.StatusOK, newUser)
}

func (u *UserHandler) Login(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBind(&user); err != nil {
		utils.HandleError(ctx, http.StatusBadRequest, "bad request", err.Error())
		return
	}

	// GetID from Database
	infoUser, err := u.ur.GetIDFromEmail(ctx, user.Email)
	if err != nil {
		utils.HandleError(ctx, http.StatusBadRequest, "bad request", err.Error())
		return
	}

	// Get password & role from where ID is match
	userCred, err := u.ur.GetPasswordFromID(ctx, infoUser.Id)
	if err != nil {
		log.Println("error getting password & role")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Bandingkan password
	hashCfg := pkg.NewHashConfig()
	isMatched, err := hashCfg.CompareHashAndPassword(user.Password, userCred.Password)
	if err != nil {
		log.Println("Internal Server Error.\nCause: ", err.Error())
		re := regexp.MustCompile("hash|crypto|argon2id|format")
		if re.Match([]byte(err.Error())) {
			log.Println("Error during Hashing")
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "internal server error",
		})
		return
	}

	if !isMatched {
		err := fmt.Errorf("login failed: nama atau password salah")
		utils.HandleError(ctx, http.StatusBadRequest, "nama atau password salah", err.Error())
		return
	}

	// Jika match, maka buatkan jwt dan kirim via response
	claims := pkg.NewJWTClaims(infoUser.Id, userCred.Role)
	jwtToken, err := claims.GenToken()
	if err != nil {
		utils.HandleError(ctx, http.StatusInternalServerError, "internal server error", err.Error())
		return
	}

	loginData := models.SuccessLoginData{
		Role:  claims.Role,
		Token: jwtToken,
	}

	response := models.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Data:    loginData,
	}

	ctx.JSON(http.StatusOK, response)
}
