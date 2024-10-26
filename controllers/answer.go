package controllers

import (
	"net/http"

	"github.com/DaniloFaraum/studere-backend/models"
	"github.com/DaniloFaraum/studere-backend/requests"
	"github.com/DaniloFaraum/studere-backend/utils"
	"github.com/gin-gonic/gin"
)

func CreateAnswerController(ctx *gin.Context) {
	request := requests.CreateAnswerRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		utils.HandleControllerError(ctx, http.StatusBadRequest, "answer validation error", err)
		return
	}

	answer := models.Answer{
		IDQuestion: request.IDQuestion,
		Text:       request.Text,
	}

	if err := db.Create(&answer).Error; err != nil {
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not create answer on db", err)
		return
	}

	utils.SendSuccess(ctx, "create-answer", request)
}

func ListAnswersController(ctx *gin.Context) {
	var answers []models.Answer
	var answerResponses []models.AnswerResponse

	if err := db.Find(&answers).Error; err != nil {
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find answers", err)
		return
	}

	for _, answer := range answers {
		answerResponses = append(answerResponses, models.AnswerResponse{
			ID:         answer.ID,
			IDQuestion: answer.IDQuestion,
			Text:       answer.Text,
		})
	}

	utils.SendSuccess(ctx, "list-answers", answerResponses)
}

func ShowAnswerController(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		utils.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsrequired("id", "queryParam").Error())
		return
	}

	answer := models.Answer{}

	if err := db.First(&answer, id).Error; err != nil {
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find answer", err)
		return
	}

	utils.SendSuccess(ctx, "show-answer", answer)
}

func UpdateAnswerController(ctx *gin.Context) {
	request := requests.UpdateAnswerRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		utils.HandleControllerError(ctx, http.StatusBadRequest, "answer validation error", err)
		return
	}

	id := ctx.Query("id")
	if id == "" {
		utils.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsrequired("id", "queryParam").Error())
		return
	}

	answer := models.Answer{}

	if err := db.First(&answer, id).Error; err != nil {
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find answer", err)
		return
	}

	// Update the fields based on the request
	if request.Text != "" {
		answer.Text = request.Text
	}

	if err := db.Save(&answer).Error; err != nil {
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not update answer", err)
	}
	utils.SendSuccess(ctx, "update-answer", answer)
}
