package controller

import (
	"github.com/gin-gonic/gin"
	"hacktiv8-final-project4/helper"
	"hacktiv8-final-project4/request"
	"hacktiv8-final-project4/service"
	"net/http"
	"strconv"
)

type ProductController struct {
	productService *service.ProductService
}

func NewProductController(productService *service.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (controller *ProductController) GetProducts(context *gin.Context) {
	categories := controller.productService.All()
	helper.JsonResponse(context, http.StatusOK, categories)
}

func (controller *ProductController) CreateProduct(context *gin.Context) {
	var productRequest request.ProductRequest
	helper.ReadFromRequestBody(context, &productRequest)

	product := controller.productService.Create(productRequest)
	helper.JsonResponse(context, http.StatusCreated, product)
}

func (controller *ProductController) UpdateProduct(context *gin.Context) {
	var productRequest request.ProductRequest
	helper.ReadFromRequestBody(context, &productRequest)

	productId, _ := strconv.Atoi(context.Param("id"))
	product := controller.productService.Update(uint(productId), productRequest)
	helper.JsonResponse(context, http.StatusOK, product)
}

func (controller *ProductController) DeleteProduct(context *gin.Context) {
	productId, _ := strconv.Atoi(context.Param("id"))
	err := controller.productService.Delete(uint(productId))
	helper.PanicIfError(err)

	helper.JsonResponse(context, http.StatusOK, struct {
		Message string `json:"message"`
	}{
		"Product has been successfully deleted",
	})
}
