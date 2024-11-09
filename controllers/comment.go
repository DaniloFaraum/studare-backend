package controllers

import (
    "net/http"

    "github.com/DaniloFaraum/studere-backend/models"
    "github.com/DaniloFaraum/studere-backend/requests"
    "github.com/DaniloFaraum/studere-backend/utils"
    "github.com/gin-gonic/gin"
)

func CreateCommentController(ctx *gin.Context) {
    request := requests.CreateCommentRequest{}

    ctx.BindJSON(&request)

    if err := request.Validate(); err != nil {
        utils.HandleControllerError(ctx, http.StatusBadRequest, "comment validation error", err)
        return
    }

    comment := models.Comment{
		IDUser: request.IDUser,
		IDCourse: request.IDCourse,
		Content: request.Content,
    }

    if err := db.Create(&comment).Error; err != nil {
        utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not create comment on db", err)
        return
    }

    utils.SendSuccess(ctx, "create-comment", request)
}

func ListCommentsController(ctx *gin.Context) {
    var comments []models.Comment
    var commentResponses []models.CommentResponse

    if err := db.Find(&comments).Error; err != nil {
        utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find comments", err)
        return
    }

    for _, comment := range comments {
        commentResponses = append(commentResponses, models.CommentResponse{
            ID: comment.ID,
            IDUser: comment.IDUser,
            IDCourse: comment.IDCourse,
            Content: comment.Content,
            Likes: comment.Likes,
            Dislikes: comment.Dislikes,
        })
    }

    utils.SendSuccess(ctx, "list-comments", commentResponses)
}

func ShowCommentController(ctx *gin.Context) {
    id := ctx.Param("id")
    if id == "" {
        utils.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsrequired("id", "queryParam").Error())
        return
    }

    comment := models.Comment{}

    if err := db.First(&comment, id).Error; err != nil {
        utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find comment", err)
        return
    }

    utils.SendSuccess(ctx, "show-comment", comment)
}

func UpdateCommentController(ctx *gin.Context) {
    request := requests.UpdateCommentRequest{}

    ctx.BindJSON(&request)

    if err := request.Validate(); err != nil {
        utils.HandleControllerError(ctx, http.StatusBadRequest, "comment validation error", err)
        return
    }

    id := ctx.Param("id")
    if id == "" {
        utils.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsrequired("id", "queryParam").Error())
        return
    }

    comment := models.Comment{}

    if err := db.First(&comment, id).Error; err != nil {
        utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find comment", err)
        return
    }

    // Update the fields based on the request
    if request.Content != "" {
        comment.Content = request.Content
    }

    if err := db.Save(&comment).Error; err != nil {
        utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not update comment", err)
    }
    utils.SendSuccess(ctx, "update-comment", comment)
}

func DeleteCommentController(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == ""{
		utils.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsrequired("id","queryParam").Error())
		return
	}

	comment := models.Comment{}

	if err := db.First(&comment, id).Error; err!=nil{
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find comment", err)
		return
	}

	if err := db.Delete(&comment).Error; err!=nil{
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not delete comment", err)
		return
	}

	utils.SendSuccess(ctx, "delete-comment", comment)
}