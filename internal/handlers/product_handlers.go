package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/yrnThiago/gdlp-go/internal/usecase"
)

type ProductHandlers struct {
	ProductUseCase *usecase.ProductUseCase
}

func NewProductHandlers(
	createProductUseCase *usecase.ProductUseCase,
) *ProductHandlers {
	return &ProductHandlers{
		ProductUseCase: createProductUseCase,
	}
}

func (p *ProductHandlers) ProductHandler(w http.ResponseWriter, r *http.Request) {
	var input usecase.ProductInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := p.ProductUseCase.Create(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *ProductHandlers) ListProductsHandler(w http.ResponseWriter, r *http.Request) {
	output, err := p.ProductUseCase.GetMany()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (p *ProductHandlers) FindByProductIdHandler(w http.ResponseWriter, r *http.Request) {
	// output, err := p.ProductUseCase.Create(input)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	//
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(output)
}
