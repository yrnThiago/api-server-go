package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/yrnThiago/gdlp-go/internal/cmd/pub"
	"github.com/yrnThiago/gdlp-go/internal/usecase"
)

type OrderHandlers struct {
	CreateOrderUseCase *usecase.CreateOrderUseCase
	ListOrdersUseCase  *usecase.ListOrdersUseCase
}

func NewOrderHandlers(
	createOrderUseCase *usecase.CreateOrderUseCase,
	listOrdersUseCase *usecase.ListOrdersUseCase,
) *OrderHandlers {
	return &OrderHandlers{
		CreateOrderUseCase: createOrderUseCase,
		ListOrdersUseCase:  listOrdersUseCase,
	}
}

func (p *OrderHandlers) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateOrderInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := p.CreateOrderUseCase.Execute(input)
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
	output, err := p.ListOrdersUseCase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
