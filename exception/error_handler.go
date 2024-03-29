package exception

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"hacktiv8-final-project4/helper"
	"net/http"
	"strings"
)

func ErrorHandler(context *gin.Context, err interface{}) {
	defer func() {
		gin.DefaultErrorWriter.Write([]byte(fmt.Sprintf("%v", err)))
		if notFoundError(context, err) {
			return
		}

		if loginError(context, err) {
			return
		}

		if validationErrors(context, err) {
			return
		}

		if badRequestError(context, err) {
			return
		}

		internalServerError(context, err)
	}()
}

func validationErrors(context *gin.Context, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		errorMap := make(map[string]string)
		for _, e := range exception {
			fieldName := strings.ToLower(e.Field())
			tag := e.Tag()
			param := e.Param()
			customMessage := helper.CustomMessages[tag]
			if customMessage != "" {
				message := strings.Replace(customMessage, "%s", fieldName, -1)
				if strings.Contains(message, "%param%") {
					message = strings.Replace(message, "%param%", param, -1)
				}
				errorMap[fieldName] = message
			} else {
				errorMap[fieldName] = fmt.Sprintf("%s is invalid.", fieldName)
			}
		}

		helper.JsonResponse(context, http.StatusBadRequest, gin.H{
			"message": "Validation error(s) occurred.",
			"errors":  errorMap,
		})
		return true
	} else {
		return false
	}
}

func loginError(context *gin.Context, err interface{}) bool {
	if exception, ok := err.(LoginError); ok {
		helper.JsonResponse(context, http.StatusUnauthorized, gin.H{
			"message": exception.Error.Error(),
		})
		return true
	} else {
		return false
	}
}

func badRequestError(context *gin.Context, err interface{}) bool {
	exception, ok := err.(BadRequestError)
	if ok {
		helper.JsonResponse(context, http.StatusBadRequest, gin.H{
			"message": exception.Error.Error(),
		})
		return true
	} else {
		return false
	}
}

func notFoundError(context *gin.Context, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		helper.JsonResponse(context, http.StatusNotFound, gin.H{
			"message": exception.Error.Error(),
		})
		return true
	} else {
		return false
	}
}

func internalServerError(context *gin.Context, err interface{}) {
	helper.JsonResponse(context, http.StatusInternalServerError, gin.H{
		"message": err,
	})
}
