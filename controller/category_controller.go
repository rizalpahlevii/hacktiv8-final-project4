package controller

import (
	"github.com/gin-gonic/gin"
	"hacktiv8-final-project4/helper"
	"hacktiv8-final-project4/request"
	"hacktiv8-final-project4/service"
	"net/http"
	"strconv"
)

type CategoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return CategoryController{categoryService: categoryService}
}

func (controller CategoryController) GetCategories(context *gin.Context) {
	categories := controller.categoryService.All()
	helper.JsonResponse(context, http.StatusOK, categories)
}

func (controller CategoryController) CreateCategory(context *gin.Context) {
	var categoryRequest request.CategoryRequest
	helper.ReadFromRequestBody(context, &categoryRequest)

	category := controller.categoryService.Create(categoryRequest)
	helper.JsonResponse(context, http.StatusCreated, category)
}

func (controller CategoryController) UpdateCategory(context *gin.Context) {
	var categoryRequest request.CategoryRequest
	helper.ReadFromRequestBody(context, &categoryRequest)

	categoryId, _ := strconv.Atoi(context.Param("id"))
	category := controller.categoryService.Update(categoryId, categoryRequest)
	helper.JsonResponse(context, http.StatusOK, category)
}

func (controller CategoryController) DeleteCategory(context *gin.Context) {
	categoryId, _ := strconv.Atoi(context.Param("id"))
	err := controller.categoryService.Delete(categoryId)
	helper.PanicIfError(err)

	helper.JsonResponse(context, http.StatusOK, struct {
		Message string `json:"message"`
	}{
		"Category has been successfully deleted",
	})
}
