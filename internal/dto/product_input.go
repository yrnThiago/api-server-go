package dto

type ProductInputDto struct {
	Name  string  `validate:"required,gt=5"`
	Price float64 `validate:"required,min=1"`
	Stock int     `validate:"required,min=1"`
}
