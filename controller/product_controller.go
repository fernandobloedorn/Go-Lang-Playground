package controller

import (
	"api/model"
	"api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {

	products := []model.Product{
		{
			ID:    1,
			Name:  "iPhone 15 Plus",
			Price: 1064.60,
		},
		{
			ID:    2,
			Name:  "Galaxy S23",
			Price: 1099.00,
		},
	}

	ctx.JSON(http.StatusOK, products)

}
