package usecase

import (
	"context"
	"encoding/json"

	"go.uber.org/zap"

	"github.com/yrnThiago/api-server-go/config"
	"github.com/yrnThiago/api-server-go/internal/dto"
	"github.com/yrnThiago/api-server-go/internal/entity"
	infra "github.com/yrnThiago/api-server-go/internal/infra/redis"
	"github.com/yrnThiago/api-server-go/internal/utils"
)

const RedisProductsKey = "all-products"

type ProductUseCase struct {
	ProductRepository IProductRepository
}

func NewProductUseCase(productRepository IProductRepository) *ProductUseCase {
	return &ProductUseCase{
		ProductRepository: productRepository,
	}
}

func NewProduct(name string, price float64, stock int) *entity.Product {
	return &entity.Product{
		Name:  name,
		Price: price,
		Stock: stock,
	}
}

func (u *ProductUseCase) Add(
	input dto.ProductInputDto,
) (*dto.ProductOutputDto, error) {
	err := utils.ValidateStruct(input)
	if err != nil {
		return nil, err
	}

	product := entity.NewProduct(input.Name, input.Price, input.Stock)
	err = u.ProductRepository.Add(product)
	if err != nil {
		return nil, err
	}

	config.Logger.Info(
		"new product added",
		zap.String("name", input.Name),
	)

	infra.Redis.Del(context.Background(), RedisProductsKey)

	return dto.NewProductOutputDto(product), nil
}

func (u *ProductUseCase) GetMany() ([]*dto.ProductOutputDto, error) {
	var productsDto []*dto.ProductOutputDto
	ctx := context.Background()

	productsRedis, _ := infra.Redis.Get(ctx, RedisProductsKey)

	if !utils.IsEmpty(productsRedis) {
		json.Unmarshal([]byte(productsRedis), &productsDto)
		return productsDto, nil
	}

	products, err := u.ProductRepository.GetMany()
	if err != nil {
		return nil, err
	}

	for _, product := range products {
		productsDto = append(productsDto, dto.NewProductOutputDto(product))
	}

	productsJson, err := json.Marshal(productsDto)
	if err != nil {
		return nil, err
	}

	infra.Redis.Set(ctx, RedisProductsKey, string(productsJson), config.Env.RateLimitWindow)
	return productsDto, nil
}

func (u *ProductUseCase) GetById(userId, productId string) (*dto.ProductOutputDto, error) {
	var product *entity.Product

	productWithOffer, err := infra.Redis.Get(context.Background(), "offer-"+userId+"-"+productId)
	if err := json.Unmarshal([]byte(productWithOffer), &product); err == nil {
		return dto.NewProductOutputDto(product), nil
	}

	product, err = u.ProductRepository.GetById(productId)
	if err != nil {
		return nil, err
	}

	return dto.NewProductOutputDto(product), nil
}

func (u *ProductUseCase) UpdateById(
	id string,
	input dto.ProductInputDto,
) (*dto.ProductOutputDto, error) {

	err := utils.ValidateStruct(input)
	if err != nil {
		return nil, err
	}

	product, err := u.ProductRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	product.Name = input.Name
	product.Price = input.Price
	product.Stock = input.Stock

	updatedProduct, err := u.ProductRepository.UpdateById(product)
	if err != nil {
		return nil, err
	}

	infra.Redis.Del(context.Background(), RedisProductsKey)

	return dto.NewProductOutputDto(updatedProduct), nil
}

func (u *ProductUseCase) DeleteById(
	id string,
) (*dto.ProductOutputDto, error) {

	product, err := u.ProductRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	err = u.ProductRepository.DeleteById(id)
	if err != nil {
		return nil, err
	}

	infra.Redis.Del(context.Background(), RedisProductsKey)
	return dto.NewProductOutputDto(product), nil
}
