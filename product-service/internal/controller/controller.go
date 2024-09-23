package controller

import (
	model "harsh/internal/models"
	"harsh/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductController struct {
	ProductService *service.ProductService
}

func NewProductController(productservice *service.ProductService) *ProductController {
	return &ProductController{
		ProductService: productservice,
	}
}

func (p *ProductController) GetAllProduct(c *gin.Context) {
	products, err := p.ProductService.GetAllProduct(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot get product"})
	}
	c.JSON(http.StatusAccepted, products)

}

func (p *ProductController) GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := p.ProductService.GetProductById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot get product"})
	}
	c.JSON(http.StatusOK, product)
}

func (p *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	err := p.ProductService.DeleteProduct(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot delete product"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func (p *ProductController) CreateProduct(c *gin.Context) {
	var product model.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "invalid data"})
	}
	err = p.ProductService.CreateProduct(c.Request.Context(), &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create product"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Product created!"})
}

func (p *ProductController) ModifyProduct(c *gin.Context) {
	//get the product id to be modified
	id := c.Param("id")

	// covert id to objectId
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	// get the new data
	var product *model.Product
	err = c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "invalid data"})
	}
	// assign the oid[extracted from url] to the product
	product.Id = oid
	// modify product
	err = p.ProductService.ModifyProduct(c.Request.Context(), product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to modify product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product modified!"})
}
