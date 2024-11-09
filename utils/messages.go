package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ErrParamIsrequired(name string, typ string) error {
	return fmt.Errorf("param: %s (%s) is required", name, typ)
}

func ErrNoValidFields() error {
	return fmt.Errorf("at least one valid field must be provided")
}

func SendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})

}

func SendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s operation sucessful", op),
		"data":    data,
	})
}

func HandleControllerError(ctx *gin.Context, statusCode int, userMessage string, err error) {
	if err == gorm.ErrRecordNotFound {
		SendError(ctx, http.StatusNotFound, userMessage)
	} else {
		logger.Errorf("%s: %v", userMessage, err.Error())
		SendError(ctx, statusCode, userMessage)
	}
}
