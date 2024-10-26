package controllers

import (
    "net/http"

    "github.com/DaniloFaraum/studere-backend/models"
    "github.com/DaniloFaraum/studere-backend/requests"
    "github.com/DaniloFaraum/studere-backend/utils"
    "github.com/gin-gonic/gin"
)

func CreateCourseController(ctx *gin.Context) {
    request := requests.CreateCourseRequest{}

    ctx.BindJSON(&request)

    if err := request.Validate(); err != nil {
        utils.HandleControllerError(ctx, http.StatusBadRequest, "course validation error", err)
        return
    }

    course := models.Course{
		Name: request.Name,
		Description: request.Description,
		Link: request.Link,
		Rating: request.Rating,
		Duration: request.Duration,
		Author: request.Author,
		Institution: request.Institution,
		IDImage: request.IDImage,
    }

    if err := db.Create(&course).Error; err != nil {
        utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not create course on db", err)
        return
    }

    utils.SendSuccess(ctx, "create-course", request)
}

func ListCoursesController(ctx *gin.Context) {
    var courses []models.Course
    var courseResponses []models.CourseResponse

    if err := db.Find(&courses).Error; err != nil {
        utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find courses", err)
        return
    }

    for _, course := range courses {
        courseResponses = append(courseResponses, models.CourseResponse{
			ID: course.ID,
			Name: course.Name,
			Description: course.Description,
			Link: course.Link,
			Rating: course.Rating,
			Duration: course.Duration,
			Author: course.Author,
			Institution: course.Institution,
			IDImage: course.IDImage,
        })
    }

    utils.SendSuccess(ctx, "list-courses", courseResponses)
}

func ShowCourseController(ctx *gin.Context) {
    id := ctx.Query("id")
    if id == "" {
        utils.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsrequired("id", "queryParam").Error())
        return
    }

    course := models.Course{}

    if err := db.First(&course, id).Error; err != nil {
        utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find course", err)
        return
    }

    utils.SendSuccess(ctx, "show-course", course)
}

func UpdateCourseController(ctx *gin.Context) {
    request := requests.UpdateCourseRequest{}

    ctx.BindJSON(&request)

    if err := request.Validate(); err != nil {
        utils.HandleControllerError(ctx, http.StatusBadRequest, "course validation error", err)
        return
    }

    id := ctx.Query("id")
    if id == "" {
        utils.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsrequired("id", "queryParam").Error())
        return
    }

    course := models.Course{}

    if err := db.First(&course, id).Error; err != nil {
        utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find course", err)
        return
    }

    // Update the fields based on the request
    switch {
    case request.Name != "":
        course.Name = request.Name
    case request.Description != "":
        course.Description = request.Description
    case request.Link != "":
        course.Link = request.Link
    case request.Rating != 0:
        course.Rating = request.Rating
    case request.Duration != "":
        course.Duration = request.Duration
    case request.Author != "":
        course.Author = request.Author
    case request.Institution != "":
        course.Institution = request.Institution
    case request.IDImage != 0:
        course.IDImage = request.IDImage
    }

    if err := db.Save(&course).Error; err != nil {
        utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not update course", err)
    }
    utils.SendSuccess(ctx, "update-course", course)
}