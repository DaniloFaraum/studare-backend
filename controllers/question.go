package controllers

import (
	"net/http"

	"github.com/DaniloFaraum/studere-backend/models"
	"github.com/DaniloFaraum/studere-backend/requests"
	"github.com/DaniloFaraum/studere-backend/utils"
	"github.com/gin-gonic/gin"
)

func CreateQuestionController(ctx *gin.Context) {
    request := requests.CreateQuestionRequest{}

    ctx.BindJSON(&request)

    if err := request.Validate(); err != nil {
        utils.HandleControllerError(ctx, http.StatusBadRequest, "request validation error", err)
        return
    }

    question := models.Question{
        //completar
    }

    if err := db.Create(&question).Error; err != nil {
        utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not create question on db", err)
        return
    }

    utils.SendSuccess(ctx, "create-question", request)
}

func ListQuestionsController(ctx *gin.Context) {
	var questions []models.Question

	if err := db.Find(&questions).Error; err!=nil{
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find questions", err)
		return
	}
	utils.SendSuccess(ctx, "list-questions", questions)
}

func ShowQuestionController(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == ""{
		utils.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsrequired("id","queryParam").Error())
		return
	}

	question := models.Question{}

	if err := db.First(&question, id).Error; err!=nil{
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find question", err)
		return
	}

	utils.SendSuccess(ctx, "show-question", question)
}

func UpdateQuestionController(ctx *gin.Context) {
	request := requests.UpdateQuestionRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
        utils.HandleControllerError(ctx, http.StatusBadRequest, "question validation error", err)
		return
	}

	id := ctx.Query("id")
	if id == ""{
		utils.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsrequired("id","queryParam").Error())
		return
	}

	question := models.Question{}

	if err := db.First(&question, id).Error; err!=nil{
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find question", err)
		return
	}

	if request.Name != "" {
        //uuuu
    }

	if err := db.Save(&question).Error; err!=nil{
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not update question", err)
	}
	utils.SendSuccess(ctx, "update-question", question)
}