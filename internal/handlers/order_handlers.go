package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/yrnThiago/gdlp-go/internal/cmd/pub"
	"github.com/yrnThiago/gdlp-go/internal/usecase"
)

type OrderHandlers struct {
	OrderUseCase *usecase.OrderUseCase
}

func NewOrderHandlers(
	createOrderUseCase *usecase.OrderUseCase,
) *OrderHandlers {
	return &OrderHandlers{
		OrderUseCase: createOrderUseCase,
	}
}

func (p *OrderHandlers) Add(w http.ResponseWriter, r *http.Request) {
	var input usecase.OrderInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := p.OrderUseCase.Add(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)

	go pub.SendMessage(output)
}

func (p *OrderHandlers) GetMany(w http.ResponseWriter, r *http.Request) {
	output, err := p.OrderUseCase.GetMany()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (p *OrderHandlers) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	output, err := p.OrderUseCase.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (p *OrderHandlers) UpdateById(w http.ResponseWriter, r *http.Request) {
	var input usecase.OrderInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := chi.URLParam(r, "id")
	_, err = p.OrderUseCase.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	new, err := p.OrderUseCase.UpdateById(id, &input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(new)
}
func (p *OrderHandlers) DeleteById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := p.OrderUseCase.DeleteById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Order removed")
}
