package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/yrnThiago/api-server-go/internal/usecase"
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

func (p *ProductHandlers) Add(w http.ResponseWriter, r *http.Request) {
	var input usecase.ProductInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := p.ProductUseCase.Add(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *ProductHandlers) GetMany(w http.ResponseWriter, r *http.Request) {
	output, err := p.ProductUseCase.GetMany()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (p *ProductHandlers) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	output, err := p.ProductUseCase.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (p *ProductHandlers) UpdateById(w http.ResponseWriter, r *http.Request) {
	var input usecase.ProductInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := chi.URLParam(r, "id")
	_, err = p.ProductUseCase.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	new, err := p.ProductUseCase.UpdateById(id, &input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(new)
}
func (p *ProductHandlers) DeleteById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := p.ProductUseCase.DeleteById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Product removed")
}
