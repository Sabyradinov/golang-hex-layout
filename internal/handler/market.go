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

type Market struct {
	log       logger.AppLogger
	marketSrv service.IMarket
}

func NewMarket(opt *Options) *Market {
	return &Market{log: opt.Logger, marketSrv: opt.Services.Market}
}

// Buy method to buy product
// @Summary method to buy product
// @Description method to buy product
// @Tags Market
// @Accept  json
// @Produce  json
// @Param comparisons body model.ProductRequest{} false "request body"
// @Success 200 {object} common.BaseResponse{} "response body"
// @Failure 400,404 {object} common.BaseResponse{} "error body"
// @Router /market/pay [post]
func (h Market) Buy(ctx *gin.Context) {
	var request model.ProductRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		h.log.ErrorWithCode(ctx, "Buy", -911, err.Error(), nil)
		return
	}

	err := h.marketSrv.Buy(ctx, request)
	if err != nil {
		h.log.ErrorWithCode(ctx, "Buy", -911, err.Error(), nil)
		return
	}

	ctx.JSON(http.StatusOK, common.BaseResponse{
		Code:    0,
		Message: "success",
	})
}

// GetById method to get product by id
// @Summary method to get product by id
// @Description method to get product by id
// @Tags Market
// @Accept  json
// @Produce  json
// @Param id path string true "product id"
// @Success 200 {object} common.BaseResponse{} "response body"
// @Failure 400,404 {object} common.BaseResponse{} "error body"
// @Router /market/:id [get]
func (h Market) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		h.log.ErrorWithCode(ctx, "GetById", -911, err.Error(), nil)
		return
	}

	res, err := h.marketSrv.GetProductById(ctx, id)
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
