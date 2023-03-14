package services

import (
	"crud-example/initial/entities"
	"crud-example/initial/repositories"
	"net/http"
	"strconv"
)

type IProductService interface {
	CreateProduct(req *http.Request)
	UpdateProduct(req *http.Request)
	DeleteProductById(req *http.Request)
	GetAllProducts(req *http.Request) (entities.Products, error)
	FirstProduct(req *http.Request) (entities.Product, error)
}

type productService struct {
	repository repositories.IProductRepository
}

func NewProductService(
	respository repositories.IProductRepository,
) IProductService {
	return &productService{
		repository: respository,
	}
}

func (product *productService) CreateProduct(req *http.Request) {

	price, err := strconv.ParseFloat(req.FormValue("price"), 64)

	if err != nil {
		return
	}

	var p entities.Product = entities.Product{
		Name:  req.FormValue("name"),
		Text:  req.FormValue("text"),
		Price: price,
		Image: "",
	}

	product.repository.Create(&p)

}

func (product *productService) UpdateProduct(req *http.Request) {
	productID, _ := strconv.Atoi(req.URL.Query().Get("id"))

	p, err := product.repository.First(productID)

	if err != nil {
		return
	}

	name := req.FormValue("name")
	text := req.FormValue("text")
	price, _ := strconv.Atoi(req.FormValue("price"))

	data := &entities.Product{
		ID:    p.ID,
		Name:  name,
		Text:  text,
		Price: float64(price),
	}

	product.repository.Update(data)
}

func (product *productService) DeleteProductById(req *http.Request) {
	
	productID, err := strconv.Atoi(req.URL.Query().Get("id"))

	if err != nil {
		return
	}

	product.repository.DeleteById(productID)
}

func (product *productService) GetAllProducts(req *http.Request) (entities.Products, error) {

	products, err := product.repository.GetAll()

	return products, err
}

func (product *productService) FirstProduct(req *http.Request) (entities.Product, error) {

	id, err := strconv.Atoi(req.URL.Query().Get("id"))

	if err != nil {
		return entities.Product{}, err
	}

	p, err := product.repository.First(id)

	if err != nil {
		return entities.Product{}, err
	}

	return p, nil
}
