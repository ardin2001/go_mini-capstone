package controllers

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/ardin2001/go_mini-capstone/helpers"
	"github.com/ardin2001/go_mini-capstone/middlewares"
	"github.com/ardin2001/go_mini-capstone/models"
	"github.com/ardin2001/go_mini-capstone/services"
	"github.com/labstack/echo/v4"
)

type ProductInterfaceC interface {
	GetProductsController(c echo.Context) error
	GetProductController(c echo.Context) error
	CreateProductController(c echo.Context) error
	UpdateProductController(c echo.Context) error
	DeleteProductController(c echo.Context) error
}

type ProductStructC struct {
	productS services.ProductInterfaceS
}

func NewProductControllers(productS services.ProductInterfaceS) ProductInterfaceC {
	return &ProductStructC{
		productS,
	}
}

func (p *ProductStructC) GetProductsController(c echo.Context) error {
	products, check := p.productS.GetProductsService()
	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: "err()",
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    products,
		Message: "Successfull get users account",
		Status:  true,
	})
}

func (p *ProductStructC) GetProductController(c echo.Context) error {
	id := c.Param("id")
	product, check := p.productS.GetProductService(id)

	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: "err()",
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    product,
		Message: "Successfull get user account",
		Status:  true,
	})
}

func (p *ProductStructC) CreateProductController(c echo.Context) error {
	data, err := middlewares.AdminVerification(c)
	if !err {
		return data
	}

	product := models.Product{}
	c.Bind(&product)

	image, _ := c.FormFile("gambar")
	filename, err := UploadImage(image)
	if !err {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: filename,
			Status:  false,
		})
	}
	product.Gambar = filename

	_, check := p.productS.CreateProductService(&product)

	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: "err()",
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    product,
		Message: "Successfull create users account",
		Status:  true,
	})
}

func (p *ProductStructC) UpdateProductController(c echo.Context) error {
	data, err := middlewares.AdminVerification(c)
	if !err {
		return data
	}

	id := c.Param("id")
	product := models.Product{}
	c.Bind(&product)

	dataProduct, check := p.productS.UpdateProductService(&product, id)

	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: "err()",
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    dataProduct,
		Message: "Successfull update user account",
		Status:  true,
	})
}

func (p *ProductStructC) DeleteProductController(c echo.Context) error {
	data, err := middlewares.AdminVerification(c)
	if !err {
		return data
	}

	id := c.Param("id")
	check := p.productS.DeleteProductService(id)

	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: "err()",
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    id,
		Message: "Successfull delete user account",
		Status:  true,
	})
}

func UploadImage(img *multipart.FileHeader) (string, bool) {
	src, err := img.Open()
	if err != nil {
		return "error upload image", false
	}
	defer src.Close()

	// Destination
	ext := filepath.Ext(img.Filename)
	currentTime := time.Now().UnixNano()
	newFileName := strconv.Itoa(int(currentTime)) + ext
	dst, err := os.Create("images/" + newFileName)
	if err != nil {
		return "error directory image", false
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return "error save image", false
	}

	return newFileName, true
}
