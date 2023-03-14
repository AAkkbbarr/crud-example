package dto

type CreateProductDTO struct {
	Name  string  `json:"name"`
	Text  string  `json:"text"`
	Price float64 `json:"price"`
	Image string  `json:"image"`
}

type UpdateProductDTO struct {
	Name  string  `json:"name"`
	Text  string  `json:"text"`
	Price float64 `json:"price"`
	Image string  `json:"image"`
}
