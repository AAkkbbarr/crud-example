package entities

type Product struct {
	ID    int
	Name  string
	Text  string
	Price float64
	Image string
}

type Products []Product
