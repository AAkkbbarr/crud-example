package controller

import (
	"crud-example/initial/entities"
	"crud-example/initial/services"
	"html/template"
	"net/http"
)

type IProductController interface {
	Index(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Store(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
	Edit(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type productController struct {
	service services.IProductService
}

func NewProductController(
	service services.IProductService,
) IProductController {
	return &productController{
		service: service,
	}
}

func (product *productController) Index(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles(
		"public/product/index.html", "public/includes/header.html", "public/includes/footer.html",
	)

	type data struct {
		Products entities.Products
		Count    int
	}

	products, _ := product.service.GetAllProducts(r)
	count := len(products)

	view.ExecuteTemplate(w, "PRODUCT_INDEX", data{Products: products, Count: count})
}

func (product *productController) Create(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles(
		"public/product/create.html", "public/includes/header.html", "public/includes/footer.html",
	)

	view.ExecuteTemplate(w, "PRODUCT_CREAT", nil)
}

func (product *productController) Show(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles(
		"public/product/show.html", "public/includes/header.html", "public/includes/footer.html",
	)

	p, _ := product.service.FirstProduct(r)

	view.ExecuteTemplate(w, "PRODUCT_SHOW", p)
}

func (product *productController) Edit(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles(
		"public/product/edit.html", "public/includes/header.html", "public/includes/footer.html",
	)

	p, _ := product.service.FirstProduct(r)

	view.ExecuteTemplate(w, "PRODUCT_EDIT", p)
}

func (product *productController) Store(w http.ResponseWriter, r *http.Request) {
	product.service.CreateProduct(r)
	http.Redirect(w, r, "/products/create", 302)
}

func (product *productController) Update(w http.ResponseWriter, r *http.Request) {

	product.service.UpdateProduct(r)

	http.Redirect(w, r, "/products", 302)
}

func (product *productController) Delete(w http.ResponseWriter, r *http.Request) {

	product.service.DeleteProductById(r)

	http.Redirect(w, r, "/products", 302)
}
