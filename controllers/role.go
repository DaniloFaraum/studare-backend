package controllers

import (
	"net/http"

	"github.com/DaniloFaraum/studere-backend/models"
	"github.com/DaniloFaraum/studere-backend/requests"
	"github.com/DaniloFaraum/studere-backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRoleController(ctx *gin.Context) {
    request := requests.CreateRoleRequest{}

    ctx.BindJSON(&request)

    if err := request.Validate(); err != nil {
        utils.HandleControllerError(ctx, http.StatusBadRequest, "role validation error", err)
        return
    }

    role := models.Role{
        Name: request.Name,
    }

    if err := db.Create(&role).Error; err != nil {
        utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not create role on db", err)
        return
    }

    utils.SendSuccess(ctx, "create-role", request)
}

func ListRolesController(ctx *gin.Context) {
	var roles []models.Role

	if err := db.Find(&roles).Error; err!=nil{
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find roles", err)
		return
	}
	utils.SendSuccess(ctx, "list-roles", roles)
}

func ShowRoleController(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == ""{
		utils.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsrequired("id","queryParam").Error())
		return
	}

	role := models.Role{}

	if err := db.First(&role, id).Error; err!=nil{
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find role", err)
		return
	}

	utils.SendSuccess(ctx, "show-role", role)
}

func UpdateRoleController(ctx *gin.Context) {
	request := requests.UpdateRoleRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
        utils.HandleControllerError(ctx, http.StatusBadRequest, "role validation error", err)
		return
	}

	id := ctx.Param("id")
	if id == ""{
		utils.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsrequired("id","queryParam").Error())
		return
	}

	role := models.Role{}

	if err := db.First(&role, id).Error; err!=nil{
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find role", err)
		return
	}

	if request.Name != "" {
        role.Name = request.Name
    }

	if err := db.Save(&role).Error; err!=nil{
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not update role", err)
	}
	utils.SendSuccess(ctx, "update-role", role)
}

func DeleteRoleController(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == ""{
		utils.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsrequired("id","queryParam").Error())
		return
	}

	role := models.Role{}

	if err := db.First(&role, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
			utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find role", err)
        }
        return
    }

	if err := db.Delete(&role).Error; err!=nil{
		utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not delete role", err)
		return
	}

	utils.SendSuccess(ctx, "delete-role", role)
}