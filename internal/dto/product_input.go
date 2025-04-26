package dto

type ProductInputDto struct {
	Name  string  `validate:"required"`
	Price float64 `validate:"required,gt=0"`
	Stock int     `validate:"required"`
}
