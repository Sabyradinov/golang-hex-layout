package market

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Sabyradinov/golang-hex-layout/common"
	"github.com/Sabyradinov/golang-hex-layout/internal/domain/port/httpClient"
	"github.com/Sabyradinov/golang-hex-layout/internal/dto"
	"strconv"

	"github.com/Sabyradinov/golang-hex-layout/config"
	"github.com/Sabyradinov/golang-hex-layout/internal/domain/port/repository"
	"github.com/Sabyradinov/golang-hex-layout/internal/domain/port/service"
	"github.com/Sabyradinov/golang-hex-layout/internal/model"
)

type product struct {
	cfg          *config.MarketUrls
	httpClient   httpClient.IClient
	productRepo  repository.IProduct
	shipmentRepo repository.IShipment
}

func New(cfg *config.Configs, repo *repository.Manager, httpClient httpClient.IClient) service.IMarket {
	return &product{cfg: &cfg.MarketUrls, productRepo: repo.Product, shipmentRepo: repo.Shipment, httpClient: httpClient}
}

func (s *product) Buy(ctx context.Context, req model.ProductRequest) (err error) {
	err = s.productRepo.Create(ctx, dto.ToProductEntity(req))

	return
}

func (s *product) SendNotification(ctx context.Context, req model.ProductRequest) (err error) {
	url := fmt.Sprintf("%v%v", s.cfg.BaseAddress.Url, s.cfg.Pay)
	requestByte, err := json.Marshal(req)
	if err != nil {
		return
	}

	args := common.HttpArgs{
		Context:       ctx,
		Url:           url,
		TimeoutSecond: s.cfg.BaseAddress.HttpTimeout,
		Data:          requestByte,
	}

	httpCode, _, err := s.httpClient.PostRequest(args)
	if err != nil {
		return
	}
	if httpCode != 200 {
		err = errors.New("http error code: " + strconv.Itoa(httpCode))
		return
	}

	return
}

func (s *product) GetProductById(ctx context.Context, Id int64) (res interface{}, err error) {
	res, err = s.productRepo.GetById(ctx, Id)
	return
}

func (s *product) GetShipmentById(ctx context.Context, Id int64) (res interface{}, err error) {
	res, err = s.shipmentRepo.GetById(ctx, Id)
	return
}

func (s *product) DeliveryToClient(ctx context.Context, req model.ShipmentRequest) (err error) {
	err = s.shipmentRepo.Create(ctx, dto.ToShipmentEntity(req))
	return
}
