package controllers

import (
	"net/http"
	"os"

	"time"

	"github.com/DaniloFaraum/studere-backend/models"
	"github.com/DaniloFaraum/studere-backend/requests"
	"github.com/DaniloFaraum/studere-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserController(ctx *gin.Context) {
	request := requests.CreateUserRequest{}

	ctx.BindJSON(&request)

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not encrypt password", err)
		return
	}

	user := models.User{
		Email: request.Email,
		Name:  request.Name,
		//ProfilePicture: request.ProfilePicture,
		Password: hash,
		RoleID:   request.RoleID,
	}

	if err := db.Create(&user).Error; err != nil {
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not create user on db", err)
		return
	}

	utils.SendSuccess(ctx, "create-user", request)
}

func ListUsersController(ctx *gin.Context) {
	var users []models.User
	var userResponses []models.UserResponse

	if err := db.Find(&users).Error; err != nil {
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find users", err)
		return
	}

	for _, user := range users {
		userResponses = append(userResponses, models.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	utils.SendSuccess(ctx, "list-users", userResponses)
}

func ShowUserController(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		utils.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsrequired("id", "queryParam").Error())
		return
	}

	user := models.User{}
	userResponse := models.UserResponse{}

	if err := db.First(&user, id).Error; err != nil {
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find user", err)
		return
	}

	userResponse = models.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	utils.SendSuccess(ctx, "show-user", userResponse)
}

func UpdateUserController(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		utils.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsrequired("id", "queryParam").Error())
		return
	}

	request := requests.UpdateUserRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		utils.HandleControllerError(ctx, http.StatusBadRequest, "user validation error", err)
		return
	}

	user := models.User{}
	if err := db.First(&user, id).Error; err != nil {
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find user", err)
		return
	}

	user.Name = request.Name
	user.Email = request.Email

	if err := db.Save(&user).Error; err != nil {
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not update user on db", err)
		return
	}

	utils.SendSuccess(ctx, "update-user", request)
}

func Login(ctx *gin.Context) {
	request := requests.LoginRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		utils.HandleControllerError(ctx, http.StatusBadRequest, "could not login", err)
		return
	}

	user := models.User{}
	if err := db.First(&user, "email = %s", request.Email).Error; err != nil {
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "invalid email or password", err)
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(request.Password)); err != nil {
		utils.HandleControllerError(ctx, http.StatusBadRequest, "invalid password or password", err)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString(os.Getenv("JWT_STRING")) //pode ser q de problema com não ter carregado o env

	if err != nil {
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not create jwt token", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})

	utils.SendSuccess(ctx, "login sucessfull", tokenString)
}
