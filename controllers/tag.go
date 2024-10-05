package controllers

import (
	"net/http"

	"github.com/DaniloFaraum/studere-backend/models"
	"github.com/DaniloFaraum/studere-backend/requests"
	"github.com/DaniloFaraum/studere-backend/utils"
	"github.com/gin-gonic/gin"
)

func CreateTagController(ctx *gin.Context) {
    request := requests.CreateTagRequest{}

    ctx.BindJSON(&request)

    if err := request.Validate(); err != nil {
        utils.HandleControllerError(ctx, http.StatusBadRequest, "tag validation error", err)
        return
    }

    tag := models.Tag{
        Name: request.Name,
    }

    if err := db.Create(&tag).Error; err != nil {
        utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not create tag on db", err)
        return
    }

    utils.SendSuccess(ctx, "create-tag", request)
}

func ListTagsController(ctx *gin.Context) {
	var tags []models.Tag

	if err := db.Find(&tags).Error; err!=nil{
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find tags", err)
		return
	}
	utils.SendSuccess(ctx, "list-tags", tags)
}