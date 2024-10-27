package controllers

import (
    "net/http"

    "github.com/DaniloFaraum/studere-backend/models"
    "github.com/DaniloFaraum/studere-backend/requests"
    "github.com/DaniloFaraum/studere-backend/utils"
    "github.com/gin-gonic/gin"
)

func CreateQuestionnaireController(ctx *gin.Context) {
    request := requests.CreateQuestionnaireRequest{}

    ctx.BindJSON(&request)

    if err := request.Validate(); err != nil {
        utils.HandleControllerError(ctx, http.StatusBadRequest, "questionnaire validation error", err)
        return
    }

    questionnaire := models.Questionnaire{
        IDUser: request.IDUser,
        Name:   request.Name,
        Ready:  request.Ready,
    }

    if err := db.Create(&questionnaire).Error; err != nil {
        utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not create questionnaire on db", err)
        return
    }

    utils.SendSuccess(ctx, "create-questionnaire", request)
}

func ListQuestionnairesController(ctx *gin.Context) {
    var questionnaires []models.Questionnaire
    var questionnaireResponses []models.QuestionnaireResponse

    if err := db.Find(&questionnaires).Error; err != nil {
        utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find questionnaires", err)
        return
    }

    for _, questionnaire := range questionnaires {
        questionnaireResponses = append(questionnaireResponses, models.QuestionnaireResponse{
            ID:     questionnaire.ID,
            IDUser: questionnaire.IDUser,
            Name:   questionnaire.Name,
            Ready:  questionnaire.Ready,
        })
    }

    utils.SendSuccess(ctx, "list-questionnaires", questionnaireResponses)
}

func ShowQuestionnaireController(ctx *gin.Context) {
    id := ctx.Query("id")
    if id == "" {
        utils.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsrequired("id", "queryParam").Error())
        return
    }

    questionnaire := models.Questionnaire{}
    questionnaireResponse := models.QuestionnaireResponse{}

    if err := db.First(&questionnaire, id).Error; err != nil {
        utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find questionnaire", err)
        return
    }

    questionnaireResponse = models.QuestionnaireResponse{
        ID:     questionnaire.ID,
        IDUser: questionnaire.IDUser,
        Name:   questionnaire.Name,
        Ready:  questionnaire.Ready,
    }

    utils.SendSuccess(ctx, "show-questionnaire", questionnaireResponse)
}

func UpdateQuestionnaireController(ctx *gin.Context) {
    id := ctx.Query("id")
    if id == "" {
        utils.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsrequired("id", "queryParam").Error())
        return
    }

    request := requests.UpdateQuestionnaireRequest{}
    ctx.BindJSON(&request)

    if err := request.Validate(); err != nil {
        utils.HandleControllerError(ctx, http.StatusBadRequest, "questionnaire validation error", err)
        return
    }

    questionnaire := models.Questionnaire{}
    if err := db.First(&questionnaire, id).Error; err != nil {
        utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find questionnaire", err)
        return
    }

    questionnaire.Name = request.Name
    questionnaire.Ready = request.Ready

    if err := db.Save(&questionnaire).Error; err != nil {
        utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not update questionnaire on db", err)
        return
    }

    utils.SendSuccess(ctx, "update-questionnaire", request)
}