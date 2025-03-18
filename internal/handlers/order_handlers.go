package handlers

import (
	"encoding/json"
	"net/http"

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

func (p *OrderHandlers) OrderHandler(w http.ResponseWriter, r *http.Request) {
	var input usecase.OrderInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := p.OrderUseCase.Create(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)

	go pub.SendMessage(output)
}

func (p *OrderHandlers) ListOrderHandler(w http.ResponseWriter, r *http.Request) {
	output, err := p.OrderUseCase.GetMany()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
