package handler

import (
	"net/http"
	"strconv"

	"github.com/bahtey101/go-rest-api-service/model"
	"github.com/bahtey101/go-rest-api-service/package/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (handler Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/cars", handler.Create)
	router.GET("/cars", handler.GetAll)
	router.GET("/cars/:id", handler.Get)
	router.PUT("/cars/:id", handler.Replace)
	router.PATCH("/cars/:id", handler.Update)
	router.DELETE("/cars/:id", handler.Delete)

	return router
}

func (handler *Handler) Create(context *gin.Context) {
	var requset struct {
		Brand        string `json:"brand" binding:"required"`
		Model        string `json:"model" binding:"required"`
		Mileage      int64  `json:"mileage" binding:"required"`
		OwnersNumber int    `json:"owners_number" binding:"required"`
	}

	if err := context.BindJSON(&requset); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "create request have invalid params"})
		return
	}

	car, err := handler.service.Create(model.Car{
		Brand:        requset.Brand,
		Model:        requset.Model,
		Mileage:      requset.Mileage,
		OwnersNumber: requset.OwnersNumber,
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, car)
}

func (handler *Handler) GetAll(context *gin.Context) {
	cars, err := handler.service.GetAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, cars)
}

func (handler *Handler) Get(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid car id param"})
		return
	}

	car, err := handler.service.Get(model.Car{ID: id})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "invalid car id param"})
		return
	}

	context.JSON(http.StatusOK, car)
}

func (handler *Handler) Replace(context *gin.Context) {
	var requset struct {
		Brand        string `json:"brand" binding:"required"`
		Model        string `json:"model" binding:"required"`
		Mileage      int64  `json:"mileage" binding:"required"`
		OwnersNumber int    `json:"owners_number" binding:"required"`
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid car id param"})
		return
	}

	car, err := handler.service.Replace(model.Car{
		ID:           id,
		Brand:        requset.Brand,
		Model:        requset.Model,
		Mileage:      requset.Mileage,
		OwnersNumber: requset.OwnersNumber,
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, car)
}

func (handler *Handler) Update(context *gin.Context) {
	var requset struct {
		Brand        string `json:"brand"`
		Model        string `json:"model"`
		Mileage      int64  `json:"mileage"`
		OwnersNumber int    `json:"owners_number"`
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid car id param"})
		return
	}

	err = handler.service.Update(model.Car{
		ID:           id,
		Brand:        requset.Brand,
		Model:        requset.Model,
		Mileage:      requset.Mileage,
		OwnersNumber: requset.OwnersNumber,
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, "")
}

func (handler *Handler) Delete(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid car id param"})
		return
	}

	err = handler.service.Delete(model.Car{ID: id})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, "")
}
