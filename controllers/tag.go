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

func ShowTagController(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == ""{
		utils.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsrequired("id","queryParam").Error())
		return
	}

	tag := models.Tag{}

	if err := db.First(&tag, id).Error; err!=nil{
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find tag", err)
		return
	}

	utils.SendSuccess(ctx, "show-tag", tag)
}

func UpdateTagController(ctx *gin.Context) {
	request := requests.UpdateTagRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
        utils.HandleControllerError(ctx, http.StatusBadRequest, "tag validation error", err)
		return
	}

	id := ctx.Query("id")
	if id == ""{
		utils.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsrequired("id","queryParam").Error())
		return
	}

	tag := models.Tag{}

	if err := db.First(&tag, id).Error; err!=nil{
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find tag", err)
		return
	}

	if request.Name != "" {
        tag.Name = request.Name
    }

	if err := db.Save(&tag).Error; err!=nil{
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not update tag", err)
	}
	utils.SendSuccess(ctx, "update-tag", tag)
}