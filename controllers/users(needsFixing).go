package controllers

import (
    "net/http"

    "github.com/DaniloFaraum/studere-backend/models"
    "github.com/DaniloFaraum/studere-backend/requests"
    "github.com/DaniloFaraum/studere-backend/utils"
    "github.com/gin-gonic/gin"
)

func CreateUserController(ctx *gin.Context) {
    request := requests.CreateUserRequest{}

    ctx.BindJSON(&request)

    if err := request.Validate(); err != nil {
        utils.HandleControllerError(ctx, http.StatusBadRequest, "user validation error", err)
        return
    }

    user := models.User{
		Email: request.Email,
		Name:  request.Name,
		ProfilePicture: request.ProfilePicture,
		Password: request.Password,
		RoleID: request.RoleID,
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
    id := ctx.Query("id")
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
    id := ctx.Query("id")
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