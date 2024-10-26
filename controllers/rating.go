package controllers

import (
    "net/http"

    "github.com/DaniloFaraum/studere-backend/models"
    "github.com/DaniloFaraum/studere-backend/requests"
    "github.com/DaniloFaraum/studere-backend/utils"
    "github.com/gin-gonic/gin"
)

func CreateRatingController(ctx *gin.Context) {
    request := requests.CreateRatingRequest{}

    ctx.BindJSON(&request)

    if err := request.Validate(); err != nil {
        utils.HandleControllerError(ctx, http.StatusBadRequest, "rating validation error", err)
        return
    }

    rating := models.Rating{
		IDUser:    request.IDUser,
		IDCourse:  request.IDCourse,
		Opinion:   request.Opinion,
		Commentary: request.Commentary,
    }

    if err := db.Create(&rating).Error; err != nil {
        utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not create rating on db", err)
        return
    }

    utils.SendSuccess(ctx, "create-rating", request)
}

func ListRatingsController(ctx *gin.Context) {
    var ratings []models.Rating
    var ratingResponses []models.RatingResponse

    if err := db.Find(&ratings).Error; err != nil {
        utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find ratings", err)
        return
    }

    for _, rating := range ratings {
        ratingResponses = append(ratingResponses, models.RatingResponse{
			ID: 	   rating.ID,
			IDUser:    rating.IDUser,
			IDCourse:  rating.IDCourse,
			Opinion:   rating.Opinion,
			Commentary: rating.Commentary,
        })
    }

    utils.SendSuccess(ctx, "list-ratings", ratingResponses)
}

func ShowRatingController(ctx *gin.Context) {
    id := ctx.Query("id")
    if id == "" {
        utils.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsrequired("id", "queryParam").Error())
        return
    }

    rating := models.Rating{}

    if err := db.First(&rating, id).Error; err != nil {
        utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find rating", err)
        return
    }

    utils.SendSuccess(ctx, "show-rating", rating)
}

func UpdateRatingController(ctx *gin.Context) {
    request := requests.UpdateRatingRequest{}

    ctx.BindJSON(&request)

    if err := request.Validate(); err != nil {
        utils.HandleControllerError(ctx, http.StatusBadRequest, "rating validation error", err)
        return
    }

    id := ctx.Query("id")
    if id == "" {
        utils.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsrequired("id", "queryParam").Error())
        return
    }

    rating := models.Rating{}

    if err := db.First(&rating, id).Error; err != nil {
        utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find rating", err)
        return
    }

    // Update the fields based on the request
    if request.Commentary != "" {
        rating.Commentary = request.Commentary
    }

    if err := db.Save(&rating).Error; err != nil {
        utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not update rating", err)
    }
    utils.SendSuccess(ctx, "update-rating", rating)
}