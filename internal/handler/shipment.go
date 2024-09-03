package handler

import (
	"github.com/Sabyradinov/golang-hex-layout/common"
	"github.com/Sabyradinov/golang-hex-layout/internal/domain/port/logger"
	"github.com/Sabyradinov/golang-hex-layout/internal/domain/port/service"
	"github.com/Sabyradinov/golang-hex-layout/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Shipment struct {
	log       logger.AppLogger
	marketSrv service.IMarket
}

func NewShipment(opt *Options) *Shipment {
	return &Shipment{log: opt.Logger, marketSrv: opt.Services.Market}
}

// DeliveryToClient method to start delivery process
// @Summary method to start delivery process
// @Description method to start delivery process
// @Tags Shipment
// @Accept  json
// @Produce  json
// @Param comparisons body model.ShipmentRequest{} false "request body"
// @Success 200 {object} common.BaseResponse{} "response body"
// @Failure 400,404 {object} common.BaseResponse{} "error body"
// @Router /shipment/delivery [post]
func (h Shipment) DeliveryToClient(ctx *gin.Context) {
	var request model.ShipmentRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		h.log.ErrorWithCode(ctx, "Buy", -911, err.Error(), nil)
		return
	}

	err := h.marketSrv.DeliveryToClient(ctx, request)
	if err != nil {
		h.log.ErrorWithCode(ctx, "Buy", -911, err.Error(), nil)
		return
	}

	ctx.JSON(http.StatusOK, common.BaseResponse{
		Code:    0,
		Message: "success",
	})
}

// GetMyDelivery method to get shipment by id
// @Summary method to get shipment by id
// @Description method to get shipment by id
// @Tags Shipment
// @Accept  json
// @Produce  json
// @Param id path string true "shipment id"
// @Success 200 {object} common.BaseResponse{} "response body"
// @Failure 400,404 {object} common.BaseResponse{} "error body"
// @Router /shipment/:id [get]
func (h Shipment) GetMyDelivery(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		h.log.ErrorWithCode(ctx, "GetById", -911, err.Error(), nil)
		return
	}

	res, err := h.marketSrv.GetShipmentById(ctx, id)
	if err != nil {
		h.log.ErrorWithCode(ctx, "GetById", -911, err.Error(), nil)
		return
	}

	ctx.JSON(http.StatusOK, common.BaseResponse{
		Code:    0,
		Message: "success",
		Data:    res,
	})
}
