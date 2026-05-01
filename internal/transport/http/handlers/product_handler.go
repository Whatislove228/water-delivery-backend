package handlers

import (
	"net/http"

	"water-delivery/internal/service"
)

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

type productResponse struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	PriceCents  int     `json:"price_cents"`
	VolumeML    int     `json:"volume_ml"`
}

func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	products, err := h.productService.ListActive(r.Context())
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "failed to list products")
		return
	}

	response := make([]productResponse, 0, len(products))

	for _, product := range products {
		response = append(response, productResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			PriceCents:  product.PriceCents,
			VolumeML:    product.VolumeML,
		})
	}

	writeJSON(w, http.StatusOK, response)
}
